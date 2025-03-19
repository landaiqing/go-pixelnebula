# <div align="center">🚀 PixelNebula Benchmarks</div>


 [中文](README.md) | English


## 📊 Introduction

This directory contains comprehensive benchmark tests for the PixelNebula library, designed to measure and analyze performance across various operational scenarios. These tests help identify potential performance bottlenecks and guide subsequent optimization efforts.

<hr/>

## 🧪 Test Content

<div>
  <table>
    <tr>
      <td align="center" width="20%">
        <img src="../assets/example_avatar.svg" width="80" height="80" alt="Basic Generation" /><br/>
        <strong>Basic Avatar Generation</strong>
      </td>
      <td align="center" width="20%">
        <img src="../assets/style.svg" width="80" height="80" alt="Styles & Themes" /><br/>
        <strong>Styles & Themes</strong>
      </td>
      <td align="center" width="20%">
        <img src="../assets/animation.svg" width="80" height="80" alt="Animations" /><br/>
        <strong>Animation Effects</strong>
      </td>
      <td align="center" width="20%">
        <img src="../assets/cache.svg" width="80" height="80" alt="Cache System" /><br/>
        <strong>Cache System</strong>
      </td>
      <td align="center" width="20%">
        <img src="../assets/performance.svg" width="80" height="80" alt="Concurrency" /><br/>
        <strong>Concurrency & Memory</strong>
      </td>
    </tr>
  </table>
</div>

The benchmarks cover the following core functionalities of PixelNebula:

### 1. 🖼️ Basic Avatar Generation (`basic_benchmark_test.go`)
- Regular avatar generation
- No-environment avatar generation
- Avatar generation with different sizes
- Multiple generations with the same ID

### 2. 🎨 Styles and Themes (`style_theme_benchmark_test.go`)
- Performance comparison between different styles
- Performance comparison between different themes
- Custom themes
- Style and theme combinations

### 3. ✨ Animation Effects (`animation_benchmark_test.go`)
- Rotation animation
- Gradient animation
- Fade-in/out animation
- Transform animation
- Color animation
- Multiple animation combinations

### 4. 💾 Cache System (`cache_benchmark_test.go`)
- No cache vs. default cache
- Different cache sizes
- Cache compression effects
- Different expiry time configurations

### 5. ⚡ Concurrency and Memory Usage (`concurrency_memory_benchmark_test.go`)
- Performance at different concurrency levels
- Concurrent performance with shared instances
- Memory usage analysis for various operations

<hr/>

## 🚀 Running Benchmarks

### Run All Tests

```bash
cd benchmark
go test -bench=. -benchmem
```

### Run Specific Test Groups

<div>
  <table>
    <tr>
      <td width="33%" align="center">
        <div style="padding: 15px; border-radius: 10px;">
          <h4>🏃 Basic Tests</h4>
          <pre>go test -bench=BenchmarkBasic -benchmem</pre>
        </div>
      </td>
      <td width="33%" align="center">
        <div style="padding: 15px; border-radius: 10px;">
          <h4>💾 Cache Tests</h4>
          <pre>go test -bench=BenchmarkCache -benchmem</pre>
        </div>
      </td>
      <td width="33%" align="center">
        <div style="padding: 15px; border-radius: 10px;">
          <h4>✨ Animation Tests</h4>
          <pre>go test -bench=BenchmarkAnimation -benchmem</pre>
        </div>
      </td>
    </tr>
  </table>
</div>

### Advanced Configuration

<div align="center">
  <table>
    <tr>
      <td width="50%" align="center">
        <div style="padding: 15px; border-radius: 10px;">
          <h4>⚙️ Set CPU Count</h4>
          <pre>go test -bench=. -benchmem -cpu=1,2,4,8</pre>
        </div>
      </td>
      <td width="50%" align="center">
        <div style="padding: 15px; border-radius: 10px;">
          <h4>⏱️ Set Iteration Count and Duration</h4>
          <pre>go test -bench=. -benchmem -count=5 -benchtime=5s</pre>
        </div>
      </td>
    </tr>
  </table>
</div>

<hr/>

## 📈 Test Result Analysis

After running the tests, results will be displayed in the following format:

```
BenchmarkBasicAvatarGeneration-8     5000     234567 ns/op    12345 B/op    123 allocs/op
```

<div>
  <table>
    <tr>
      <th align="center">Component</th>
      <th align="center">Description</th>
    </tr>
    <tr>
      <td align="center"><code>BenchmarkBasicAvatarGeneration-8</code></td>
      <td>Test name, 8 indicates using 8 CPUs</td>
    </tr>
    <tr>
      <td align="center"><code>5000</code></td>
      <td>Number of iterations the test ran</td>
    </tr>
    <tr>
      <td align="center"><code>234567 ns/op</code></td>
      <td>Average time per operation (nanoseconds)</td>
    </tr>
    <tr>
      <td align="center"><code>12345 B/op</code></td>
      <td>Average memory allocation per operation (bytes)</td>
    </tr>
    <tr>
      <td align="center"><code>123 allocs/op</code></td>
      <td>Average number of memory allocations per operation</td>
    </tr>
  </table>
</div>

<hr/>

## 📝 Adding Benchmark Tests

When adding new features, it's recommended to add corresponding benchmark tests:

1. **Create a Test Function**: Create a new test function for the specific feature, named `BenchmarkXXX`
2. **Reset Timer**: Use `b.ResetTimer()` to reset the timer after setup work is complete
3. **Create Sub-tests**: Use `b.Run()` to create sub-tests for comparing different variables
4. **Use Iteration Count**: Use `b.N` as the iteration count to ensure statistical accuracy

<div>
  <p><strong>Example:</strong></p>
  
```go
func BenchmarkMyFeature(b *testing.B) {
    // Setup code
    pn := pixelnebula.NewPixelNebula()
    
    // Reset timer before the actual benchmark
    b.ResetTimer()
    
    // Run the benchmark
    for i := 0; i < b.N; i++ {
        // Code to benchmark
        pn.MyFeature()
    }
}
```
</div>

<hr/>

<div align="center">
  <p><em>For more information, check out the PixelNebula documentation and examples.</em></p>
  <p>© 2024 landaiqing</p>
</div> 