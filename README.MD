# 图床工具

一个基于 GitHub 的图床工具，支持自动图片压缩和 jsDelivr CDN 加速。

## 功能特性

- ✅ **自动上传**：支持 URL 和本地文件路径
- ✅ **智能压缩**：自动压缩大图片（>500KB），转换为 JPEG 格式，质量 75%
- ✅ **CDN 加速**：使用 jsDelivr CDN 提供更快的访问速度
- ✅ **自动提交**：自动 Git 提交并推送到 GitHub
- ✅ **剪贴板**：自动复制 Markdown 格式的图片链接到剪贴板
- ✅ **文件名管理**：使用时间戳命名，避免文件名冲突

## 使用方法

### 方式一：直接运行（推荐）

```bash
# 从 URL 上传
go run ./src/main.go https://i.stack.imgur.com/5W3rG.png

# 从本地文件上传
go run ./src/main.go ~/Downloads/image.png
```

### 方式二：使用 Makefile

```bash
# 编译
make build

# 运行（需要提供图片URL或路径）
make run URL=https://example.com/image.png
make run URL=~/Downloads/image.png

# 安装到系统（可选）
make install
# 安装后可以直接使用: imagehosting <图片URL或路径>
```

### 方式三：编译后使用

```bash
# 编译
cd src && go build -o ../bin/imagehosting main.go

# 运行
./bin/imagehosting https://example.com/image.png
```

## 输出示例

运行后会输出 Markdown 格式的图片链接：

```markdown
![image](https://cdn.jsdelivr.net/gh/yinhui1984/imagehosting@main/images/1768877495030214000.jpg)
```

链接会自动复制到剪贴板，可以直接粘贴使用。

## 配置说明

压缩和 CDN 相关配置在 `src/main.go` 中可以调整：

```go
const (
    maxImageWidth      = 1920       // 最大宽度（像素）
    compressionQuality = 75         // JPEG 质量 (0-100)
    useJsDelivr        = true       // 是否使用 jsDelivr CDN
    minCompressSize    = 500 * 1024 // 小于此大小的文件不压缩
)
```

### 压缩规则

- 文件小于 500KB：不压缩
- 文件大于 500KB：自动压缩为 JPEG 格式
- 最大宽度限制：1920px
- 如果压缩后文件更大：保留原图
- 时间戳部分保持不变，扩展名可能变为 `.jpg`

### CDN 说明

- 默认使用 **jsDelivr CDN** 加速
- 如果 jsDelivr 生成失败，自动回退到 GitHub raw 链接
- jsDelivr 格式：`https://cdn.jsdelivr.net/gh/用户名/仓库名@分支/images/文件名`
- GitHub raw 格式：`https://github.com/用户名/仓库名/blob/分支/images/文件名?raw=true`

## 项目结构

```
.
├── images/          # 图片存储目录
├── src/             # 源代码目录
│   ├── main.go      # 主程序
│   └── go.mod       # Go 模块文件
├── .gitignore       # Git 忽略文件
├── Makefile         # 构建脚本
└── README.MD        # 说明文档
```

## 依赖要求

- Go 1.19+
- Git（已配置远程仓库）
- macOS（使用 `sips` 工具进行图片压缩）

## 注意事项

1. **Git 配置**：确保已配置 Git 远程仓库
2. **分支名称**：自动检测当前分支，默认使用 `main`
3. **文件大小**：GitHub 限制单文件最大 100MB
4. **图片格式**：压缩后统一转换为 JPEG 格式（`.jpg`）

## 工作流程

1. 下载/复制图片到 `images` 目录
2. 自动压缩图片（如果大于 500KB）
3. Git 提交并推送到 GitHub
4. 生成 jsDelivr CDN 链接
5. 复制 Markdown 格式链接到剪贴板

## 许可证

MIT
