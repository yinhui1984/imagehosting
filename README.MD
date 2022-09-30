# 图床

1. 图片存储在`images`下
2. `src`目录下是一个小脚本,功能:
    1. 上传图片到图床
    2. 生成markdown格式的图片链接
   ```bash
    go run ./src/main.go https://i.stack.imgur.com/5W3rG.png
    ```