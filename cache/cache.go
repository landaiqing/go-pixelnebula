package cache

import (
	"container/list"
	"sync"
	"time"
)

// CacheOptions 缓存配置选项
type CacheOptions struct {
	Enabled      bool            // 是否启用缓存
	Size         int             // 缓存大小，0表示无限制
	Expiration   time.Duration   // 缓存项过期时间，0表示永不过期
	EvictionType string          // 缓存淘汰策略，支持"lru"(最近最少使用)和"fifo"(先进先出)
	Compression  CompressOptions // 压缩选项
	Monitoring   MonitorOptions  // 监控选项
}

// DefaultCacheOptions 默认缓存配置
var DefaultCacheOptions = CacheOptions{
	Enabled:      true,
	Size:         100,                    // 默认缓存100个SVG
	Expiration:   time.Hour,              // 默认缓存项过期时间为1小时
	EvictionType: "lru",                  // 默认使用LRU淘汰策略
	Compression:  DefaultCompressOptions, // 默认压缩选项
	Monitoring:   DefaultMonitorOptions,  // 默认监控选项
}

// CacheKey 缓存键结构
type CacheKey struct {
	Id      string
	SansEnv bool
	Theme   int
	Part    int
}

// CacheItem 缓存项结构
type CacheItem struct {
	SVG          string    // SVG内容
	Compressed   []byte    // 压缩后的SVG数据
	IsCompressed bool      // 是否已压缩
	CreatedAt    time.Time // 创建时间
	LastUsed     time.Time // 最后使用时间
}

// PNCache SVG缓存结构
type PNCache struct {
	Options      CacheOptions
	Items        map[CacheKey]*list.Element // 存储缓存项的映射
	EvictionList *list.List                 // 用于实现LRU/FIFO的双向链表
	Mutex        sync.RWMutex
	Hits         int      // 缓存命中次数
	Misses       int      // 缓存未命中次数
	Monitor      *Monitor // 缓存监控器
}

// NewCache 创建一个新的缓存实例
func NewCache(options CacheOptions) *PNCache {
	cache := &PNCache{
		Options:      options,
		Items:        make(map[CacheKey]*list.Element),
		EvictionList: list.New(),
		Hits:         0,
		Misses:       0,
	}

	// 如果启用了监控，创建并启动监控器
	if options.Monitoring.Enabled {
		cache.Monitor = NewMonitor(cache, options.Monitoring)
		cache.Monitor.Start()
	}

	return cache
}

// NewDefaultCache 使用默认配置创建一个新的缓存实例
func NewDefaultCache() *PNCache {
	return NewCache(DefaultCacheOptions)
}

// Get 从缓存中获取SVG
func (c *PNCache) Get(key CacheKey) (string, bool) {
	if !c.Options.Enabled {
		c.Misses++
		return "", false
	}

	c.Mutex.Lock() // 使用写锁以便更新LRU信息
	defer c.Mutex.Unlock()

	element, found := c.Items[key]
	if !found {
		c.Misses++
		return "", false
	}

	// 获取缓存项
	cacheItem := element.Value.(*CacheItem)

	// 检查是否过期
	if c.Options.Expiration > 0 {
		if time.Since(cacheItem.CreatedAt) > c.Options.Expiration {
			// 删除过期项
			c.EvictionList.Remove(element)
			delete(c.Items, key)
			c.Misses++
			return "", false
		}
	}

	// 更新LRU信息
	if c.Options.EvictionType == "lru" {
		cacheItem.LastUsed = time.Now()
		c.EvictionList.MoveToFront(element)
	}

	c.Hits++

	// 如果数据已压缩，需要解压缩
	if cacheItem.IsCompressed {
		svg, err := DecompressSVG(cacheItem.Compressed, true)
		if err != nil {
			// 解压失败，返回未压缩的原始数据
			return cacheItem.SVG, true
		}
		return svg, true
	}

	return cacheItem.SVG, true
}

// Set 将SVG存入缓存
func (c *PNCache) Set(key CacheKey, svg string) {
	if !c.Options.Enabled {
		return
	}

	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	// 尝试压缩SVG数据
	var compressed []byte
	var isCompressed bool

	// 如果启用了压缩，尝试压缩SVG
	if c.Options.Compression.Enabled {
		// 首先优化SVG
		optimizedSVG := OptimizeSVG(svg)

		// 然后压缩
		compressed, isCompressed = CompressSVG(optimizedSVG, c.Options.Compression)

		// 如果压缩成功，使用优化后的SVG
		if isCompressed {
			svg = optimizedSVG
		}
	}

	// 检查是否已存在
	if element, exists := c.Items[key]; exists {
		// 更新现有项
		cacheItem := element.Value.(*CacheItem)
		cacheItem.SVG = svg
		cacheItem.Compressed = compressed
		cacheItem.IsCompressed = isCompressed
		cacheItem.LastUsed = time.Now()
		cacheItem.CreatedAt = time.Now()

		// 如果使用LRU策略，将项移到链表前端
		if c.Options.EvictionType == "lru" {
			c.EvictionList.MoveToFront(element)
		}
		return
	}

	// 如果达到大小限制，需要淘汰一个项
	if c.Options.Size > 0 && len(c.Items) >= c.Options.Size {
		c.evictItem()
	}

	// 创建新的缓存项
	now := time.Now()
	cacheItem := &CacheItem{
		SVG:          svg,
		Compressed:   compressed,
		IsCompressed: isCompressed,
		CreatedAt:    now,
		LastUsed:     now,
	}

	// 添加到链表和映射
	element := c.EvictionList.PushFront(cacheItem)
	c.Items[key] = element
}

// evictItem 根据淘汰策略移除一个缓存项
func (c *PNCache) evictItem() {
	if c.EvictionList.Len() == 0 {
		return
	}

	// 获取要淘汰的元素
	var element *list.Element
	switch c.Options.EvictionType {
	case "lru":
		// LRU策略：移除链表尾部元素（最近最少使用）
		element = c.EvictionList.Back()
	default:
		// 默认使用FIFO策略：移除链表尾部元素（最先添加）
		element = c.EvictionList.Back()
	}

	if element != nil {
		// 从链表中移除
		c.EvictionList.Remove(element)

		// 从映射中找到并删除对应的键
		for k, v := range c.Items {
			if v == element {
				delete(c.Items, k)
				break
			}
		}
	}
}

// Clear 清空缓存
func (c *PNCache) Clear() {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	c.Items = make(map[CacheKey]*list.Element)
	c.EvictionList = list.New()
	c.Hits = 0
	c.Misses = 0
}

// Size 返回当前缓存项数量
func (c *PNCache) Size() int {
	c.Mutex.RLock()
	defer c.Mutex.RUnlock()

	return len(c.Items)
}

// Stats 返回缓存统计信息
func (c *PNCache) Stats() (Hits, Misses int, hitRate float64) {
	c.Mutex.RLock()
	defer c.Mutex.RUnlock()

	Hits = c.Hits
	Misses = c.Misses
	total := Hits + Misses
	if total > 0 {
		hitRate = float64(Hits) / float64(total)
	}
	return
}

// RemoveExpired 移除所有过期的缓存项
func (c *PNCache) RemoveExpired() int {
	if c.Options.Expiration <= 0 {
		return 0
	}

	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	count := 0
	now := time.Now()

	// 遍历所有缓存项，检查是否过期
	for key, element := range c.Items {
		cacheItem := element.Value.(*CacheItem)
		if now.Sub(cacheItem.CreatedAt) > c.Options.Expiration {
			// 从链表中移除
			c.EvictionList.Remove(element)
			// 从映射中删除
			delete(c.Items, key)
			count++
		}
	}

	return count
}

// GetOptions 获取当前缓存选项
func (c *PNCache) GetOptions() CacheOptions {
	c.Mutex.RLock()
	defer c.Mutex.RUnlock()

	return c.Options
}

// UpdateOptions 更新缓存选项
func (c *PNCache) UpdateOptions(options CacheOptions) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	// 更新选项
	c.Options = options

	// 如果新的缓存大小小于当前项数，需要淘汰一些项
	if c.Options.Size > 0 && c.Options.Size < len(c.Items) {
		// 计算需要淘汰的项数
		toEvict := len(c.Items) - c.Options.Size

		// 淘汰多余的项
		for i := 0; i < toEvict; i++ {
			c.evictItem()
		}
	}
}
