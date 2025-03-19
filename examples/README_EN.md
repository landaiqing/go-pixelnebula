# PixelNebula Examples

[中文版](README.md) | English

## Example Files

This directory contains multiple examples showcasing various features of the PixelNebula library. Each example is standalone and can be run separately to understand specific functionalities.

| Filename | Description |
|----------|-------------|
| [01_basic_usage.go](01_basic_usage.go) | Demonstrates basic avatar generation, including regular and no-environment avatars |
| [02_styles_and_themes.go](02_styles_and_themes.go) | Shows how to use different styles and themes |
| [03_custom_theme_and_style.go](03_custom_theme_and_style.go) | Demonstrates how to create and use custom themes and styles |
| [04_all_animations.go](04_all_animations.go) | Showcases all supported animation effects |
| [05_svg_builder_chain.go](05_svg_builder_chain.go) | Demonstrates how to use the chainable API to generate SVGs |
| [06_cache_system.go](06_cache_system.go) | Shows the cache system functionality, including default, custom, and monitored caching |
| [07_format_conversion.go](07_format_conversion.go) | Shows how to convert SVGs to other formats |
| [08_random_avatar_generator.go](08_random_avatar_generator.go) | Interactive random avatar generator with support for multiple styles, themes, and output formats |

## How to Run Examples

Ensure you have Go properly installed and GOPATH set correctly. Then, follow these steps to run the examples:

### Run a Single Example

```bash
# For example, run the basic usage example
go run 01_basic_usage.go
```

### Run All Examples

```bash
for file in *_*.go; do
  echo "🚀 Running example: $file"
  go run $file
  echo "------------------------"
done
```

## Example Details

### 01_basic_usage.go

This example demonstrates the basic functionality of PixelNebula, including:

- Creating a basic avatar
- Generating a no-environment avatar
- Handling errors and saving files

### 02_styles_and_themes.go

This example shows how to use different styles and themes:

- Using predefined styles to generate avatars
- Applying different themes
- Combining styles and themes

### 03_custom_theme_and_style.go

This example shows how to create and use custom themes and styles:

- Creating custom color themes
- Defining custom styles
- Combining custom themes and styles

### 04_all_animations.go

This example showcases all supported animation effects:

- Rotation animation
- Gradient animation
- Fade-in/out effects
- Transform animation
- Color transformation
- Bounce effects
- Wave animation
- Blink effects
- Path animation

### 05_svg_builder_chain.go

This example shows how to use the chainable API:

- Using chain calls to create simple SVGs
- Adding animation effects
- Saving directly to file
- Converting to Base64

### 06_cache_system.go

This example demonstrates the cache system functionality:

- Using the default cache
- Configuring custom caches
- Monitoring cache performance
- Using compressed caching

### 07_format_conversion.go

This example shows the format conversion capabilities:

- Converting to Base64
- We haven't found a perfect solution for other formats yet, please feel free to make a PR

### 08_random_avatar_generator.go

This example is an interactive random avatar generator:

- Randomly generating avatars with different styles and themes

## Tips

- Each example file has detailed comments at the top explaining the functionality it demonstrates
- Check the error handling sections in the files if you encounter any issues
- Generated avatars will be saved in the locations specified in the example code
- Some examples may require creating directories to save generated files 