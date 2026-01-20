.PHONY: help build run install clean test

# 默认目标
help:
	@echo "图床工具 - Makefile 命令"
	@echo ""
	@echo "可用命令:"
	@echo "  make build      - 编译程序到 bin/imagehosting"
	@echo "  make run URL    - 运行程序（需要提供图片URL或路径）"
	@echo "  make install    - 安装到 /usr/local/bin"
	@echo "  make clean      - 清理编译文件"
	@echo "  make test       - 测试编译"
	@echo ""
	@echo "示例:"
	@echo "  make run https://example.com/image.png"
	@echo "  make run ~/Downloads/image.png"

# 编译
build:
	@echo "正在编译..."
	@mkdir -p bin
	@cd src && go build -o ../bin/imagehosting main.go
	@echo "编译完成: bin/imagehosting"

# 运行（需要提供参数，先编译再运行）
run: build
	@if [ -z "$(URL)" ]; then \
		echo "错误: 请提供图片URL或路径"; \
		echo "用法: make run URL=<图片URL或路径>"; \
		exit 1; \
	fi
	@echo "运行程序..."
	@./bin/imagehosting $(URL)

# 安装到系统路径
install: build
	@echo "正在安装到 /usr/local/bin..."
	@sudo cp bin/imagehosting /usr/local/bin/
	@echo "安装完成! 现在可以使用 'imagehosting <图片URL或路径>' 命令"

# 清理编译文件
clean:
	@echo "正在清理..."
	@rm -rf bin/
	@echo "清理完成"

# 测试编译
test:
	@echo "测试编译..."
	@cd src && go build -o /dev/null main.go
	@echo "编译测试通过!"
