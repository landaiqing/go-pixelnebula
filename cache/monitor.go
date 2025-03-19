package cache

import (
	"sync"
	"time"
)

// MonitorOptions 缓存监控选项
type MonitorOptions struct {
	Enabled          bool          // 是否启用监控
	SampleInterval   time.Duration // 采样间隔时间
	AdjustInterval   time.Duration // 调整间隔时间
	MinSize          int           // 最小缓存大小
	MaxSize          int           // 最大缓存大小
	TargetHitRate    float64       // 目标命中率
	SizeGrowthFactor float64       // 缓存大小增长因子
	SizeShrinkFactor float64       // 缓存大小收缩因子
	ExpirationFactor float64       // 过期时间调整因子
}

// DefaultMonitorOptions 默认监控选项
var DefaultMonitorOptions = MonitorOptions{
	Enabled:          true,
	SampleInterval:   time.Minute,      // 每分钟采样一次
	AdjustInterval:   time.Minute * 10, // 每10分钟调整一次
	MinSize:          50,               // 最小缓存大小
	MaxSize:          1000,             // 最大缓存大小
	TargetHitRate:    0.8,              // 目标命中率80%
	SizeGrowthFactor: 1.2,              // 增长20%
	SizeShrinkFactor: 0.8,              // 收缩20%
	ExpirationFactor: 1.5,              // 过期时间调整因子
}

// CacheStats 缓存统计信息
type CacheStats struct {
	Size          int       // 当前缓存大小
	Hits          int       // 命中次数
	Misses        int       // 未命中次数
	HitRate       float64   // 命中率
	MemoryUsage   int64     // 内存使用量（字节）
	LastAdjusted  time.Time // 最后调整时间
	SamplesCount  int       // 样本数量
	AvgAccessTime float64   // 平均访问时间（纳秒）
}

// Monitor 缓存监控器
type Monitor struct {
	options       MonitorOptions
	cache         *PNCache
	stats         CacheStats
	sampleHistory []CacheStats
	mutex         sync.RWMutex
	stopChan      chan struct{}
	isRunning     bool
}

// NewMonitor 创建一个新的缓存监控器
func NewMonitor(cache *PNCache, options MonitorOptions) *Monitor {
	return &Monitor{
		options:       options,
		cache:         cache,
		sampleHistory: make([]CacheStats, 0, 100), // 预分配100个样本的容量
		stopChan:      make(chan struct{}),
		isRunning:     false,
	}
}

// Start 启动监控器
func (m *Monitor) Start() {
	if !m.options.Enabled || m.isRunning {
		return
	}

	m.mutex.Lock()
	m.isRunning = true
	m.mutex.Unlock()

	go m.monitorRoutine()
}

// Stop 停止监控器
func (m *Monitor) Stop() {
	if !m.isRunning {
		return
	}

	m.mutex.Lock()
	m.isRunning = false
	m.mutex.Unlock()

	m.stopChan <- struct{}{}
}

// GetStats 获取当前缓存统计信息
func (m *Monitor) GetStats() CacheStats {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	return m.stats
}

// monitorRoutine 监控例程
func (m *Monitor) monitorRoutine() {
	sampleTicker := time.NewTicker(m.options.SampleInterval)
	adjustTicker := time.NewTicker(m.options.AdjustInterval)

	defer sampleTicker.Stop()
	defer adjustTicker.Stop()

	for {
		select {
		case <-m.stopChan:
			return
		case <-sampleTicker.C:
			m.collectSample()
		case <-adjustTicker.C:
			m.adjustCache()
		}
	}
}

// collectSample 收集缓存样本
func (m *Monitor) collectSample() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	// 获取缓存统计信息
	hits, misses, hitRate := m.cache.Stats()
	size := m.cache.Size()

	// 估算内存使用量（简化计算，实际应用中可能需要更精确的方法）
	memoryUsage := int64(size * 1024) // 假设每个缓存项平均占用1KB

	// 创建新的统计样本
	newStat := CacheStats{
		Size:         size,
		Hits:         hits,
		Misses:       misses,
		HitRate:      hitRate,
		MemoryUsage:  memoryUsage,
		LastAdjusted: time.Now(),
	}

	// 添加到历史样本
	m.sampleHistory = append(m.sampleHistory, newStat)

	// 限制历史样本数量，保留最近的100个样本
	if len(m.sampleHistory) > 100 {
		m.sampleHistory = m.sampleHistory[len(m.sampleHistory)-100:]
	}

	// 更新当前统计信息
	m.stats = newStat
	m.stats.SamplesCount = len(m.sampleHistory)
}

// adjustCache 根据统计信息调整缓存
func (m *Monitor) adjustCache() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	// 如果样本数量不足，不进行调整
	if len(m.sampleHistory) < 5 {
		return
	}

	// 计算平均命中率
	totalHitRate := 0.0
	for _, stat := range m.sampleHistory {
		totalHitRate += stat.HitRate
	}
	avgHitRate := totalHitRate / float64(len(m.sampleHistory))

	// 获取当前缓存选项
	cacheOptions := m.cache.GetOptions()

	// 根据命中率调整缓存大小
	newSize := cacheOptions.Size
	if avgHitRate < m.options.TargetHitRate {
		// 命中率低于目标，增加缓存大小
		newSize = int(float64(newSize) * m.options.SizeGrowthFactor)
		// 确保不超过最大大小
		if newSize > m.options.MaxSize {
			newSize = m.options.MaxSize
		}
	} else if avgHitRate > m.options.TargetHitRate+0.1 && m.stats.Size > m.options.MinSize {
		// 命中率远高于目标且缓存大小大于最小值，可以适当减小缓存
		newSize = int(float64(newSize) * m.options.SizeShrinkFactor)
		// 确保不小于最小大小
		if newSize < m.options.MinSize {
			newSize = m.options.MinSize
		}
	}

	// 根据访问模式调整过期时间
	newExpiration := cacheOptions.Expiration
	if avgHitRate < m.options.TargetHitRate {
		// 命中率低，增加过期时间
		newExpiration = time.Duration(float64(newExpiration) * m.options.ExpirationFactor)
	} else if avgHitRate > m.options.TargetHitRate+0.1 {
		// 命中率高，可以适当减少过期时间
		newExpiration = time.Duration(float64(newExpiration) / m.options.ExpirationFactor)
	}

	// 应用新的缓存选项
	if newSize != cacheOptions.Size || newExpiration != cacheOptions.Expiration {
		cacheOptions.Size = newSize
		cacheOptions.Expiration = newExpiration
		m.cache.UpdateOptions(cacheOptions)

		// 更新最后调整时间
		m.stats.LastAdjusted = time.Now()
	}
}
