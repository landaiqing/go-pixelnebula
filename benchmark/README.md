# <div align="center">🚀 PixelNebula 基准测试</div>


 [英文](README_EN.md) | 中文


## 📊 简介

本目录包含了 PixelNebula 库的全面基准测试，用于测量和分析库在各种操作场景下的性能表现。这些测试可以帮助我们识别潜在的性能瓶颈，并指导后续的优化工作。

<hr/>

## 🧪 测试内容

<div>
  <table>
    <tr>
      <td align="center" width="20%">
        <img src="../assets/example_avatar.svg" width="80" height="80" alt="基本生成" /><br/>
        <strong>基本头像生成</strong>
      </td>
      <td align="center" width="20%">
        <img src="../assets/style.svg" width="80" height="80" alt="样式和主题" /><br/>
        <strong>样式和主题</strong>
      </td>
      <td align="center" width="20%">
        <img src="../assets/animation.svg" width="80" height="80" alt="动画效果" /><br/>
        <strong>动画效果</strong>
      </td>
      <td align="center" width="20%">
        <img src="../assets/cache.svg" width="80" height="80" alt="缓存系统" /><br/>
        <strong>缓存系统</strong>
      </td>
      <td align="center" width="20%">
        <img src="../assets/performance.svg" width="80" height="80" alt="并发性能" /><br/>
        <strong>并发与内存</strong>
      </td>
    </tr>
  </table>
</div>

基准测试涵盖了 PixelNebula 的以下核心功能：

### 1. 🖼️ 基本头像生成 (`basic_benchmark_test.go`)
- 普通头像生成
- 无环境头像生成
- 不同尺寸头像生成
- 相同ID多次生成

### 2. 🎨 样式和主题 (`style_theme_benchmark_test.go`)
- 不同样式的性能对比
- 不同主题的性能对比
- 自定义主题
- 样式与主题组合

### 3. ✨ 动画效果 (`animation_benchmark_test.go`)
- 旋转动画
- 渐变动画
- 淡入淡出动画
- 变换动画
- 颜色动画
- 多个动画组合

### 4. 💾 缓存系统 (`cache_benchmark_test.go`)
- 无缓存 vs. 默认缓存
- 不同缓存大小
- 缓存压缩效果
- 不同过期时间配置

### 5. ⚡ 并发与内存使用 (`concurrency_memory_benchmark_test.go`)
- 不同并发级别的性能
- 共享实例的并发性能
- 各种操作的内存占用分析

<hr/>

## 🚀 运行基准测试

### 运行所有测试

```bash
cd benchmark
go test -bench=. -benchmem
```

### 运行特定测试组

<div align="center">
  <table>
    <tr>
      <td width="33%" align="center">
        <div style="padding: 15px; border-radius: 10px;">
          <h4>🏃 基本测试</h4>
          <pre>go test -bench=BenchmarkBasic -benchmem</pre>
        </div>
      </td>
      <td width="33%" align="center">
        <div style="padding: 15px; border-radius: 10px;">
          <h4>💾 缓存测试</h4>
          <pre>go test -bench=BenchmarkCache -benchmem</pre>
        </div>
      </td>
      <td width="33%" align="center">
        <div style="padding: 15px; border-radius: 10px;">
          <h4>✨ 动画测试</h4>
          <pre>go test -bench=BenchmarkAnimation -benchmem</pre>
        </div>
      </td>
    </tr>
  </table>
</div>

### 高级配置

<div align="center">
  <table>
    <tr>
      <td width="50%" align="center">
        <div style="padding: 15px; border-radius: 10px;">
          <h4>⚙️ 设置CPU计数</h4>
          <pre>go test -bench=. -benchmem -cpu=1,2,4,8</pre>
        </div>
      </td>
      <td width="50%" align="center">
        <div style="padding: 15px; border-radius: 10px;">
          <h4>⏱️ 设置迭代次数和时间</h4>
          <pre>go test -bench=. -benchmem -count=5 -benchtime=5s</pre>
        </div>
      </td>
    </tr>
  </table>
</div>

<hr/>

## 📈 测试结果分析

运行测试后，结果会以如下格式显示：

```
BenchmarkBasicAvatarGeneration-8     5000     234567 ns/op    12345 B/op    123 allocs/op
```

<div>
  <table>
    <tr>
      <th align="center">组成部分</th>
      <th align="center">描述</th>
    </tr>
    <tr>
      <td align="center"><code>BenchmarkBasicAvatarGeneration-8</code></td>
      <td>测试名称，8表示使用8个CPU</td>
    </tr>
    <tr>
      <td align="center"><code>5000</code></td>
      <td>测试运行的迭代次数</td>
    </tr>
    <tr>
      <td align="center"><code>234567 ns/op</code></td>
      <td>每次操作的平均耗时（纳秒）</td>
    </tr>
    <tr>
      <td align="center"><code>12345 B/op</code></td>
      <td>每次操作的平均内存分配（字节）</td>
    </tr>
    <tr>
      <td align="center"><code>123 allocs/op</code></td>
      <td>每次操作的平均内存分配次数</td>
    </tr>
  </table>
</div>

<hr/>

## 📝 添加基准测试

在添加新功能时，建议同时添加相应的基准测试：

1. **创建测试函数**：为特定功能创建新的测试函数，命名为 `BenchmarkXXX`
2. **重置计时器**：使用 `b.ResetTimer()` 在准备工作完成后重置计时器
3. **创建子测试**：使用 `b.Run()` 创建子测试以对比不同变量
4. **使用迭代计数**：使用 `b.N` 作为迭代次数以确保统计准确性

<div>
  <p><strong>示例：</strong></p>
  
```go
func BenchmarkMyFeature(b *testing.B) {
    // 准备代码
    pn := pixelnebula.NewPixelNebula()
    
    // 在实际基准测试前重置计时器
    b.ResetTimer()
    
    // 运行基准测试
    for i := 0; i < b.N; i++ {
        // 要测试的代码
        pn.MyFeature()
    }
}
```
</div>

<hr/>

<div align="center">
  <p><em>更多信息，请查看 PixelNebula 文档和示例。</em></p>
  <p>© 2024 landaiqing</p>
</div> 