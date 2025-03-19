<div style="background-image: linear-gradient(rgba(26, 26, 26, 0.85), rgba(45, 45, 45, 0.85)), url('https://img.picui.cn/free/2025/03/19/67da3b51a5dea.png'); background-size: cover; background-position: center; background-repeat: no-repeat; padding: 40px 20px; border-radius: 15px; margin-bottom: 30px; box-shadow: 0 4px 15px rgba(0,0,0,0.2);">

# <div align="center" style="display: flex; justify-content: center; align-items: center; margin-bottom: 20px;"><img src="./assets/pixel_planet.gif" height="40"/><span style="font-family: 'Press Start 2P', monospace; background: linear-gradient(45deg, #FF6B6B, #4ECDC4); -webkit-background-clip: text; -webkit-text-fill-color: transparent; text-shadow: 2px 2px 4px rgba(0,0,0,0.2);"> PixelNebula </span></div>

<div align="center" style="margin-bottom: 20px;">
  <p style="color: #ffffff;"><a href="README_ZH.md" style="color: #4ECDC4;">中文</a> | <strong>English</strong></p>
</div>

<p align="center" style="margin-bottom: 20px;">
  <img src="./assets/golang_logo.gif" alt="PixelNebula Logo" height="150" style="border-radius: 10px; box-shadow: 0 4px 8px rgba(0,0,0,0.3);" />
</p>

<div align="center" style="margin-bottom: 20px;">
  <p style="color: #ffffff; font-size: 1.2em;">
    <strong>🚀 A powerful, efficient, and customizable SVG avatar generation library for Go</strong>
  </p>
</div>

<div align="center" style="margin-bottom: 20px; padding: 15px; border-radius: 10px;">
  <a href="https://pkg.go.dev/github.com/landaiqing/go-pixelnebula"><img src="https://img.shields.io/badge/go-reference-blue?style=flat-square&logo=go" alt="Go Reference"></a>
  <a href="https://goreportcard.com/"><img src="https://img.shields.io/badge/go%20report-A+-brightgreen?style=flat-square&logo=go" alt="Go Report Card"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/license-MIT-green?style=flat-square&logo=github" alt="License"></a>
  <a href="https://github.com/landaiqing/go-pixelnebula/releases"><img src="https://img.shields.io/badge/version-0.1.0-blue?style=flat-square&logo=github" alt="Version"></a>
</div>

<p align="center" style="color: #ffffff;background: rgba(255, 255, 255, 0.1);padding: 15px; border-radius: 10px; ">
  <a href="#-features" style="color: #4ECDC4;">✨ Features</a> •
  <a href="#-installation" style="color: #4ECDC4;">📦 Installation</a> •
  <a href="#-quick-start" style="color: #4ECDC4;">🚀 Quick Start</a> •
  <a href="#-advanced-usage" style="color: #4ECDC4;">🔧 Advanced Usage</a> •
  <a href="#-animation-effects" style="color: #4ECDC4;">💫 Animation Effects</a> •
  <a href="#-cache-system" style="color: #4ECDC4;">⚡ Cache System</a> •
  <a href="#-benchmarks" style="color: #4ECDC4;">📊 Benchmarks</a> •
  <a href="#-examples-and-demos" style="color: #4ECDC4;">💡 Examples and Demos</a> •
  <a href="#-contribution-guide" style="color: #4ECDC4;">👥 Contribution Guide</a> •
  <a href="#-license" style="color: #4ECDC4;">📄 License</a>
</p>

</div>

<hr/>

## 📋 Introduction

<table>
  <tr>
    <td>
      <p><strong style="font-family: 'Press Start 2P', monospace; background: linear-gradient(45deg, #FF6B6B, #4ECDC4); -webkit-background-clip: text; -webkit-text-fill-color: transparent; text-shadow: 2px 2px 4px rgba(0,0,0,0.2);">PixelNebula</strong> is a high-performance SVG avatar generation library written in Go, focused on creating beautiful, customizable, and highly animated vector avatars. With PixelNebula, you can easily generate avatars in various styles, add animation effects, and apply different themes and style variations.</p>
      <p>Whether you're creating user avatars for applications, generating unique identifier icons, or making dynamic visual effects, PixelNebula can meet your needs.</p>
    </td>
    <td width="380">
      <p align="center">
        <img src="./assets/example_avatar.svg" height="100" width="auto" alt="PixelNebula Demo" />
        <br/>
        <a href="./assets/example_avatar.svg">Sample Avatar Display</a><br><code><a href="./pixelnebula_test.go">Sample Code</a></code>
      </p>
    </td>
  </tr>
</table>

<hr/>

## ✨ Features

<div>
  <table>
    <tr>
      <td align="center" width="100">
        <img src="./assets/diversify.svg" height="60" width="60" alt="Various Styles" /><br/>
        <strong>Various Styles</strong>
      </td>
      <td align="center" width="100">
        <img src="./assets/animation.svg" height="60" width="60" alt="Animation Effects" /><br/>
        <strong>Animation Effects</strong>
      </td>
      <td align="center" width="100">
        <img src="./assets/customize.svg" height="60" width="60" alt="Customizable" /><br/>
        <strong>Customizable</strong>
      </td>
      <td align="center" width="100">
        <img src="./assets/performance.svg" height="60" width="60" alt="High Performance" /><br/>
        <strong>High Performance</strong>
      </td>
      <td align="center" width="100">
        <img src="./assets/cache.svg" height="60" width="60" alt="Cache System" /><br/>
        <strong>Cache System</strong>
      </td>
    </tr>
  </table>
</div>

- 🎨 **Various Styles and Themes**: Built-in multiple styles and theme combinations to meet different design needs
- 🔄 **Rich Animation Effects**: Support for rotation, gradient, transformation, fade-in/out, and other animations
- 🛠️ **Fully Customizable**: Custom styles, themes, colors, and animation effects
- ⚡ **High-Performance Design**: Optimized code structure and caching mechanism for fast generation
- 💾 **Smart Cache System**: Built-in cache system to improve efficiency of repetitive generation
- 📊 **Cache Monitoring**: Support for cache usage monitoring and analysis
- 🔍 **Chainable API**: Clean and clear API design with support for fluent chaining
- 📱 **Responsive Design**: Support for custom sizes, adapting to various display environments

<hr/>

## 📦 Installation

Using Go toolchain to install the package:

```bash
go get github.com/landaiqing/go-pixelnebula
```

<hr/>

## 🚀 Quick Start

### Basic Usage

Below is a basic example of generating a simple SVG avatar:

<table>
<tr>
<td>

```go
package main

import (
    "fmt"
    "os"
    
    "github.com/landaiqing/go-pixelnebula"
    "github.com/landaiqing/go-pixelnebula/style"
)

func main() {
    // Create a new PixelNebula instance
    pn := pixelnebula.NewPixelNebula()
    
    // Set style and size
    pn.WithStyle(style.GirlStyle)
    pn.WithSize(231, 231)
    
    // Generate SVG and save to file
    svg, err := pn.Generate("unique-id-123", false).ToSVG()
    if err != nil {
        fmt.Printf("Failed to generate SVG: %v\n", err)
        return
    }
    
    // Save to file
    err = os.WriteFile("my_avatar.svg", []byte(svg), 0644)
    if err != nil {
        fmt.Printf("Failed to save file: %v\n", err)
        return
    }
    
    fmt.Println("Avatar successfully generated: my_avatar.svg")
}
```

</td>
<td width="250">
<div align="center">
  <img src="./assets/example_avatar_1.svg" alt="Avatar Example" width="auto" height="100"/>
  <p><a href="./assets/example_avatar_1.svg">Generated Avatar Example</a></p>
</div>
</td>
</tr>
</table>

<hr/>

## 🔧 Advanced Usage

### Custom Themes and Styles

<details open>
<summary><strong>Click to Expand/Collapse Code Example</strong></summary>

```go
// Custom style
customStyles := []style.StyleSet{
    {
        // First custom style
        style.TypeEnv:   `<circle cx="50%" cy="50%" r="48%" fill="#FILL0;"></circle>`,
        style.TypeHead:  `<circle cx="50%" cy="50%" r="35%" fill="#FILL0;"></circle>`,
        style.TypeClo:   `<rect x="25%" y="65%" width="50%" height="30%" fill="#FILL0;"></rect>`,
        style.TypeEyes:  `<circle cx="40%" cy="45%" r="5%" fill="#FILL0;"></circle><circle cx="60%" cy="45%" r="5%" fill="#FILL1;"></circle>`,
        style.TypeMouth: `<path d="M 40% 60% Q 50% 70% 60% 60%" stroke="#FILL0;" stroke-width="2" fill="none"></path>`,
        style.TypeTop:   `<path d="M 30% 30% L 50% 10% L 70% 30%" stroke="#FILL0;" stroke-width="4" fill="#FILL1;"></path>`,
    },
}

// Apply custom style
pn2.WithCustomizeStyle(customStyles)
// Custom theme
customThemes := []theme.Theme{
{
    theme.ThemePart{
        // Environment part colors
        "env": []string{"#FF5733", "#C70039"},
        // Head colors
        "head": []string{"#FFC300", "#FF5733"},
        // Clothes colors
        "clo": []string{"#2E86C1", "#1A5276"},
        // Eyes colors
        "eyes": []string{"#000000", "#FFFFFF"},
        // Mouth colors
        "mouth": []string{"#E74C3C"},
        // Top decoration colors
        "top": []string{"#884EA0", "#7D3C98"},
        },
    },
}

pn.WithCustomizeTheme(customTheme)
```

</details>

### Using SVGBuilder Chainable API

<details open>
<summary><strong>Click to Expand/Collapse Code Example</strong></summary>

```go
pn := NewPixelNebula().WithDefaultCache()
pn.Generate("my-avatar", false).
    SetStyle(style.GirlStyle).
    SetTheme(0).
    SetSize(231, 231).
    SetRotateAnimation("env", 0, 360, 10, -1).
    SetGradientAnimation("env", []string{"#3498db", "#2ecc71", "#f1c40f", "#e74c3c", "#9b59b6"}, 8, -1, true).
    Build().
    ToSVG()
```

</details>

<hr/>

## 💫 Animation Effects

PixelNebula supports multiple animation effects to bring your avatars to life:

<div>
  <table>
    <tr>
      <td align="center" width="150">
        <img src="./assets/example_avatar_2.svg" height="120" width="120" alt="Rotation Animation" /><br/>
        <a href="./assets/example_avatar_2.svg">Rotation</a>
      </td>
      <td align="center" width="150">
        <img src="./assets/example_avatar_3.svg" height="120" width="120" alt="Gradient Animation" /><br/>
        <a href="./assets/example_avatar_3.svg">Gradient</a>
      </td>
      <td align="center" width="150">
        <img src="./assets/example_avatar_4.svg" height="120" width="120" alt="Fade In/Out" /><br/>
        <a href="./assets/example_avatar_4.svg">Fade In/Out</a>
      </td>
      <td align="center" width="150">
        <img src="./assets/example_avatar_5.svg" height="120" width="120" alt="Path Animation" /><br/>
        <a href="./assets/example_avatar_5.svg">Path</a>
      </td>
    </tr>
  </table>
</div>

<details>
<summary><strong>Click to View Animation Code Examples</strong></summary>

```go
pn := NewPixelNebula()

// Set style
pn.WithStyle(style.AfrohairStyle)
pn.WithTheme(0)

// 1. Rotation animation - Rotate environment and head
pn.WithRotateAnimation("env", 0, 360, 10, -1) // Infinite loop environment rotation

// 2. Gradient animation - Environment gradient
pn.WithGradientAnimation("env", []string{"#3498db", "#2ecc71", "#f1c40f", "#e74c3c", "#9b59b6"}, 8, -1, true)
// 2. Gradient animation - Eyes gradient
pn.WithGradientAnimation("eyes", []string{"#3498db", "#2ecc71", "#f1c40f", "#e74c3c", "#9b59b6"}, 8, -1, true)

// 3. Fade in/out animation - Eyes blinking
pn.WithFadeAnimation("eyes", "1", "0.3", 2, -1)

// 4. Transform animation - Mouth scaling
//pn.WithTransformAnimation("mouth", "scale", "1 1", "1.2 1.2", 1, -1)

// 5. Color animation - Hair color change
pn.WithColorAnimation("top", "fill", "#9b59b6", "#e74c3c", 3, -1)
// 5. Color animation - Clothes color change
pn.WithColorAnimation("clo", "fill", "#9b59b6", "#e74c3c", 3, -1)

// 6. Bounce animation - Mouth bouncing
pn.WithBounceAnimation("mouth", "transform", "0,0", "0,-10", 5, 2.5, -1)
// 6. Rotation animation - Mouth rotation
pn.WithRotateAnimation("mouth", 0, 360, 10, -1) // Infinite loop mouth rotation

//// 7. Wave animation - Clothes wave effect
//pn.WithWaveAnimation("clo", 5, 0.2, "horizontal", 4, -1)

// 8. Blink animation - Top decoration blinking
//pn.WithBlinkAnimation("head", 0.3, 1.0, 4, 6, -1)
// 8. Wave animation - Environment wave effect
//pn.WithWaveAnimation("clo", 5, 2, "horizontal", 4, -1)

// 9. Path animation - Eyes moving along path
//pn.WithPathAnimation("eyes", "M 0,0 C 10,-10 -10,-10 0,0", 3, -1)

pn.WithBounceAnimation("eyes", "transform", "0,0", "0,-5", 5, 2, -1)

// 10. Path animation with rotation - Eyes rotating while moving
//pn.WithPathAnimationRotate("eyes", "M 0,0 C 5,5 -5,5 0,0", "auto", 4, -1)
```

</details>

<hr/>

## ⚡ Cache System

PixelNebula has a built-in smart cache system to improve the efficiency of repeated generation:

<details>
<summary><strong>Click to View Cache Configuration Code Examples</strong></summary>

```go
// Use default cache configuration
pn.WithDefaultCache()

// Custom cache configuration
customCacheOptions := cache.CacheOptions{
    Enabled:    true,
    DirectorySize:       100,           // Maximum cache entries
    Expiration: 1 * time.Hour, // Cache expiration time
//... Other configuration options
}
// Create a PixelNebula instance with custom cache
pn := pixelnebula.NewPixelNebula().WithCache(customCacheOptions)

// Enable cache monitoring
pn.WithMonitoring(cache.MonitorOptions{
   Enabled:        true,
    SampleInterval: 5 * time.Second,
//... Other configuration options
})

// Enable cache compression
pn.WithCompression(cache.CompressOptions{
    Enabled:      true,
    Level:        6,
    MinSizeBytes: 100, // Minimum compression size (bytes)
    //... Other configuration options
})
```

</details>

<hr/>

## 📊 Benchmarks

We have conducted comprehensive benchmark tests on PixelNebula to ensure high performance in various scenarios. Below are sample test results (Test environment: Intel i7-12700K, 32GB RAM):

<div>
  <table>
    <tr>
      <th>Operation</th>
      <th>Time per Operation</th>
      <th>Memory Allocation</th>
      <th>Allocations</th>
    </tr>
    <tr>
      <td>Basic Avatar Generation</td>
      <td>3.5 ms/op</td>
      <td>328 KB/op</td>
      <td>52 allocs/op</td>
    </tr>
    <tr>
      <td>No-Environment Avatar</td>
      <td>2.8 ms/op</td>
      <td>256 KB/op</td>
      <td>48 allocs/op</td>
    </tr>
    <tr>
      <td>With Rotation Animation</td>
      <td>4.2 ms/op</td>
      <td>384 KB/op</td>
      <td>62 allocs/op</td>
    </tr>
    <tr>
      <td>Using Cache (Hit)</td>
      <td>0.3 ms/op</td>
      <td>48 KB/op</td>
      <td>12 allocs/op</td>
    </tr>
    <tr>
      <td>Concurrent Generation (10)</td>
      <td>5.7 ms/op</td>
      <td>392 KB/op</td>
      <td>58 allocs/op</td>
    </tr>
  </table>
</div>

### Running Benchmarks

To run benchmarks in your own environment, execute:

```bash
cd benchmark
go test -bench=. -benchmem
```

For more detailed benchmark information, see [benchmark/README.md](benchmark/README.md).

<hr/>

## 💡 Examples and Demos

### 📚 Example Code

We've prepared a complete set of example code covering all core features to help you get started quickly.

<div>
  <table>
    <tr>
      <td align="center" width="200">
        <img src="./assets/base_use.svg" width="100" height="100" alt="Basic Usage" /><br/>
        <strong>Basic Usage</strong><br/>
        <a href="examples/01_basic_usage.go">View Code</a>
      </td>
      <td align="center" width="200">
        <img src="./assets/style.svg" width="100" height="100" alt="Styles & Themes" /><br/>
        <strong>Styles & Themes</strong><br/>
        <a href="examples/02_styles_and_themes.go">View Code</a>
      </td>
      <td align="center" width="200">
        <img src="./assets/theme.svg" width="100" height="100" alt="Custom Theme" /><br/>
        <strong>Custom Theme</strong><br/>
        <a href="examples/03_custom_theme_and_style.go">View Code</a>
      </td>
    </tr>
    <tr>
      <td align="center" width="200">
        <img src="./assets/animation.svg" width="100" height="100" alt="Animations" /><br/>
        <strong>Animations</strong><br/>
        <a href="examples/04_all_animations.go">View Code</a>
      </td>
      <td align="center" width="200">
        <img src="./assets/api.svg" width="100" height="100" alt="Chain API" /><br/>
        <strong>Chain API</strong><br/>
        <a href="examples/05_svg_builder_chain.go">View Code</a>
      </td>
      <td align="center" width="200">
        <img src="./assets/cache.svg" width="100" height="100" alt="Cache System" /><br/>
        <strong>Cache System</strong><br/>
        <a href="examples/06_cache_system.go">View Code</a>
      </td>
    </tr>
    <tr>
      <td align="center" width="200">
        <img src="./assets/convert.svg" width="100" height="100" alt="Format Conversion" /><br/>
        <strong>Format Conversion</strong><br/>
        <a href="examples/07_format_conversion.go">View Code</a>
      </td>
      <td align="center" width="200">
        <img src="./assets/random.svg" width="100" height="100" alt="Random Avatar Generator" /><br/>
        <strong>Random Avatar Generator</strong><br/>
        <a href="examples/08_random_avatar_generator.go">View Code</a>
      </td>
      <td align="center" width="200">
        <h1>...</h1>
      </td>
    </tr>
  </table>
</div>

<details>
<summary><strong>📋 Detailed Example Descriptions</strong></summary>

1. **Basic Usage** [01_basic_usage.go](examples/01_basic_usage.go)
   - Create basic PixelNebula avatars
   - Generate regular avatars and no-environment avatars
   - Demonstrate basic configuration and error handling

2. **Styles and Themes** [02_styles_and_themes.go](examples/02_styles_and_themes.go)
   - Use different styles to generate avatars
   - Apply multiple theme combinations
   - Demonstrate cycling through styles and themes

3. **Custom Themes and Styles** [03_custom_theme_and_style.go](examples/03_custom_theme_and_style.go)
   - Create custom themes
   - Define custom styles
   - Demonstrate combining themes and styles

4. **Animation Effects** [04_all_animations.go](examples/04_all_animations.go)
   - Rotation animation
   - Gradient animation
   - Fade-in/out effects
   - Transform animation
   - Color transformation
   - Bounce effects
   - Wave animation
   - Blink effects
   - Path animation

5. **SVG Builder Chainable API** [05_svg_builder_chain.go](examples/05_svg_builder_chain.go)
   - Basic chainable calls
   - Chainable calls with animation
   - Direct save to file
   - Base64 conversion

6. **Cache System** [06_cache_system.go](examples/06_cache_system.go)
   - Using default cache
   - Custom cache configuration
   - Cache monitoring functionality
   - Compressed cache examples

7. **Format Conversion** [07_format_conversion.go](examples/07_format_conversion.go)
   - Base64 encoding
   - We haven't found a perfect solution for other formats yet, please feel free to make a PR

8. **Random Avatar Generator** [08_random_avatar_generator.go](examples/08_random_avatar_generator.go)
   - Generate random avatars with different styles and themes

</details>

### 🎮 Running Examples

<details>
<summary><strong>📋 How to Run Examples</strong></summary>

1. **Prerequisites**
   - Ensure Go is installed (Go 1.16+ recommended)
   - Make sure GOPATH is properly set

2. **Clone the Repository**
   ```bash
   git clone github.com/landaiqing/go-pixelnebula.git
   cd go-pixelnebula
   ```

3. **Run a Single Example**
   ```bash
   go run examples/01_basic_usage.go
   ```

4. **Run All Examples**
   ```bash
   for file in examples/*_*.go; do
     echo "🚀 Running: $file"
     go run $file
     echo "------------"
   done
   ```

5. **Run Benchmarks**
   ```bash
   cd benchmark
   go test -bench=. -benchmem
   ```

**💡 Tip:** Check the top comments in each example for detailed functionality and customization options.

</details>

<hr/>

## 👥 Contribution Guide

Contributions of code, issue reports, or suggestions are welcome! Please follow these steps:

<div>
  <table>
    <tr>
      <td align="center">
        <img src="./assets/step_1.svg" width="60" height="60" alt="Step 1" /><br/>
        <strong>Fork the Repository</strong>
      </td>
      <td align="center">
        <img src="./assets/step_2.svg" width="60" height="60" alt="Step 2" /><br/>
        <strong>Create a Branch</strong><br/>
        <code>git checkout -b feature/amazing-feature</code>
      </td>
      <td align="center">
        <img src="./assets/step_3.svg" width="60" height="60" alt="Step 3" /><br/>
        <strong>Commit Changes</strong><br/>
        <code>git commit -m 'Add some feature'</code>
      </td>
      <td align="center">
        <img src="./assets/step_4.svg" width="60" height="60" alt="Step 4" /><br/>
        <strong>Push Branch</strong><br/>
        <code>git push origin feature/amazing-feature</code>
      </td>
      <td align="center">
        <img src="./assets/step_5.svg" width="60" height="60" alt="Step 5" /><br/>
        <strong>Open PR</strong>
      </td>
    </tr>
  </table>
</div>

<hr/>

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

<hr/>

<div align="center">
  <p><strong>Made with ❤️ and Go</strong></p>
  <p>© 2025 landaiqing</p>
  
  <br/>
</div> 