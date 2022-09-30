# 图床

1. 图片存储在`images`下
2. `src`目录下是一个小脚本,功能:
    1. 上传图片到图床
    2. 生成markdown格式的图片链接
   ```bash
    go run ./src/main.go https://i.stack.imgur.com/5W3rG.png
    # 或者
    go run ./src/main.go ~/Downloads/5W3rG.png
    ```
   输出
   ```markdown
    ![image](https://github.com/yinhui1984/imagehosting/blob/main/images/1664515782371748000-5W3rG.png?raw=true)
    ```