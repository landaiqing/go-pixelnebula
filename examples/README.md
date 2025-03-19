# PixelNebula 示例

[英文版](README_EN.md) | 中文版

## 示例文件

这个目录包含了多个示例，展示了如何使用 PixelNebula 库的各种功能。每个示例都是独立的，你可以单独运行它们来了解特定的功能。

| 文件名 | 描述 |
|--------|------|
| [01_basic_usage.go](01_basic_usage.go) | 演示基本的头像生成，包括常规头像和无环境头像 |
| [02_styles_and_themes.go](02_styles_and_themes.go) | 展示如何使用不同的样式和主题 |
| [03_custom_theme_and_style.go](03_custom_theme_and_style.go) | 演示如何创建和使用自定义主题和样式 |
| [04_all_animations.go](04_all_animations.go) | 展示所有支持的动画效果 |
| [05_svg_builder_chain.go](05_svg_builder_chain.go) | 演示如何使用链式API生成SVG |
| [06_cache_system.go](06_cache_system.go) | 展示缓存系统的功能，包括默认缓存、自定义缓存和监控 |
| [07_format_conversion.go](07_format_conversion.go) | 展示如何将SVG转换为其他格式 |
| [08_random_avatar_generator.go](08_random_avatar_generator.go) | 交互式的随机头像生成器，支持多种样式、主题和输出格式 |

## 如何运行示例

确保您已正确安装 Go 环境并设置了 GOPATH。然后，按照以下步骤运行示例：

### 运行单个示例

```bash
# 例如，运行基本用法示例
go run 01_basic_usage.go
```

### 运行所有示例

```bash
for file in *_*.go; do
  echo "🚀 运行示例: $file"
  go run $file
  echo "------------------------"
done
```

## 示例说明

### 01_basic_usage.go

这个示例展示了 PixelNebula 的基本功能，包括：

- 创建一个基本的头像
- 生成一个无环境的头像
- 处理错误和保存文件

### 02_styles_and_themes.go

这个示例展示了如何使用不同的样式和主题：

- 使用预定义的样式生成头像
- 应用不同的主题
- 组合样式和主题

### 03_custom_theme_and_style.go

这个示例展示了如何创建和使用自定义主题和样式：

- 创建自定义颜色主题
- 定义自定义样式
- 组合自定义主题和样式

### 04_all_animations.go

这个示例展示了所有支持的动画效果：

- 旋转动画
- 渐变动画
- 淡入淡出效果
- 变换动画
- 颜色变换
- 弹跳效果
- 波浪动画
- 闪烁效果
- 路径动画

### 05_svg_builder_chain.go

这个示例展示了如何使用链式API：

- 使用链式调用创建简单的SVG
- 添加动画效果
- 直接保存到文件
- 转换为Base64

### 06_cache_system.go

这个示例展示了缓存系统的功能：

- 使用默认缓存
- 配置自定义缓存
- 监控缓存性能
- 使用压缩缓存

### 07_format_conversion.go

这个示例展示了格式转换功能：

- 转换为Base64
- 其他格式暂未找到完美解决方案，欢迎 PR

### 08_random_avatar_generator.go

这个示例是一个交互式的随机头像生成器：

- 随机生成不同样式和主题的头像

## 提示

- 每个示例文件顶部都有详细的注释，解释了该示例所展示的功能
- 如果遇到任何问题，请检查文件中的错误处理部分
- 生成的头像会保存在示例代码指定的位置
- 有些示例可能需要创建目录来保存生成的文件 