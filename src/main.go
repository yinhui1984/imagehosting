package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"time"
)

const (
	maxFileSize = 100 * 1024 * 1024 // 100MB - GitHub limit
	httpTimeout = 30 * time.Second
	// 压缩参数
	maxImageWidth      = 1920       // 最大宽度
	compressionQuality = 75         // JPEG 质量 (0-100, 75 适合技术博客)
	useJsDelivr        = true       // 是否使用 jsDelivr CDN
	minCompressSize    = 500 * 1024 // 小于此大小的文件不压缩 (500KB)
)

// copyLocalFile copies a local file to destination
func copyLocalFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}

	return nil
}

// DownloadFile will download an url to a local file.
func downloadFile(URL, fileName string) error {
	u, err := url.Parse(URL)
	if err != nil || u.Scheme == "" || u.Host == "" {
		// not a valid url, maybe local file
		if _, err := os.Stat(URL); errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("file %s does not exist", URL)
		}
		// Use Go standard library instead of shell command
		return copyLocalFile(URL, fileName)
	}

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: httpTimeout,
	}

	// Get the response bytes from the url
	response, err := client.Get(URL)
	if err != nil {
		return fmt.Errorf("failed to download file: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(response.Body)

	if response.StatusCode != 200 {
		// BUG FIX: Use strconv.Itoa instead of string(rune())
		return fmt.Errorf("got response code %d", response.StatusCode)
	}

	// Check Content-Length for file size
	if response.ContentLength > maxFileSize {
		return fmt.Errorf("file size (%d bytes) exceeds GitHub limit (%d bytes)", response.ContentLength, maxFileSize)
	}

	// Create an empty file
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	// Write the bytes to the file with size limit check
	limitedReader := io.LimitReader(response.Body, maxFileSize+1)
	written, err := io.Copy(file, limitedReader)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	if written > maxFileSize {
		return fmt.Errorf("file size exceeds GitHub limit (%d bytes)", maxFileSize)
	}

	return nil
}

func getCurrentDir() string {
	// Try to get current working directory first
	wd, err := os.Getwd()
	if err == nil {
		// Check if we're in the imagehosting directory
		if strings.HasSuffix(wd, "imagehosting") || strings.Contains(wd, "imagehosting") {
			return wd + "/"
		}
	}

	// Fallback to hardcoded path
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("failed to get home directory: ", err)
	}
	return path.Join(home, "Documents/github/imagehosting/")
}

func getRemoteDir() string {
	currentDir := getCurrentDir()
	cmd := exec.Command("git", "remote", "get-url", "origin")
	cmd.Dir = currentDir
	remote, err := cmd.Output()
	if err != nil {
		log.Fatal("can not get remote url: ", err)
	}

	// BUG FIX: More robust URL processing
	remoteStr := strings.TrimSpace(string(remote))
	remoteStr = strings.TrimSuffix(remoteStr, ".git")

	// Convert SSH URL to HTTPS URL if needed
	if strings.HasPrefix(remoteStr, "git@") {
		// git@github.com:user/repo -> https://github.com/user/repo
		remoteStr = strings.Replace(remoteStr, "git@github.com:", "https://github.com/", 1)
		remoteStr = strings.TrimSuffix(remoteStr, ".git")
	}

	return remoteStr
}

func getCurrentBranch() (string, error) {
	currentDir := getCurrentDir()
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Dir = currentDir
	out, err := cmd.Output()
	if err != nil {
		return "main", err // Default to main if can't detect
	}
	return strings.TrimSpace(string(out)), nil
}

func generateFileName(url string, contentType string) string {
	fileName := path.Base(url)
	extension := path.Ext(fileName)

	// BUG FIX: Handle missing extension
	if extension == "" {
		// Try to infer from Content-Type
		if strings.Contains(contentType, "jpeg") || strings.Contains(contentType, "jpg") {
			extension = ".jpg"
		} else if strings.Contains(contentType, "png") {
			extension = ".png"
		} else if strings.Contains(contentType, "gif") {
			extension = ".gif"
		} else if strings.Contains(contentType, "webp") {
			extension = ".webp"
		} else {
			// Default to .png if can't determine
			extension = ".png"
		}
	}

	return path.Join(getCurrentDir()+"/images", strconv.FormatInt(time.Now().UnixNano(), 10)+extension)
}

func ensureDir(dir string) error {
	return os.MkdirAll(dir, 0755)
}

func execCommand(command string) error {
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func execCommandAndReturnOutput(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	out, err := cmd.Output()
	return string(out), err
}

// copyToClipboard safely copies text to clipboard
func copyToClipboard(text string) error {
	cmd := exec.Command("pbcopy")
	cmd.Stdin = strings.NewReader(text)
	return cmd.Run()
}

// compressImage 压缩图片，直接覆盖原文件（时间戳部分不变，但扩展名可能变为 .jpg）
func compressImage(imagePath string) (string, error) {
	// 检查文件是否存在
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		return imagePath, fmt.Errorf("image file does not exist: %s", imagePath)
	}

	// 获取文件信息
	fileInfo, err := os.Stat(imagePath)
	if err != nil {
		return imagePath, fmt.Errorf("failed to get file info: %w", err)
	}

	// 如果文件小于 minCompressSize，不压缩（小文件压缩意义不大）
	if fileInfo.Size() < minCompressSize {
		log.Println("File is small, skipping compression")
		return imagePath, nil
	}

	// 获取原文件扩展名
	originalExt := path.Ext(imagePath)
	// 生成新的文件名（时间戳部分不变，但扩展名改为 .jpg）
	newPath := strings.TrimSuffix(imagePath, originalExt) + ".jpg"

	// 如果已经是 .jpg 格式，直接使用原文件名
	if originalExt == ".jpg" || originalExt == ".jpeg" {
		newPath = imagePath
	}

	// 使用 macOS 自带的 sips 工具进行压缩
	// sips 参数说明：
	// -s format jpeg: 转换为 JPEG 格式（更好的压缩比）
	// -s formatOptions 75: JPEG 质量 75%
	// --resampleHeightWidthMax 1920: 最大尺寸限制为 1920px
	cmd := exec.Command("sips",
		"-s", "format", "jpeg",
		"-s", "formatOptions", strconv.Itoa(compressionQuality),
		"--resampleHeightWidthMax", strconv.Itoa(maxImageWidth),
		imagePath,
		"--out", newPath)

	err = cmd.Run()
	if err != nil {
		return imagePath, fmt.Errorf("compression failed: %w", err)
	}

	// 检查压缩后的文件大小
	compressedInfo, err := os.Stat(newPath)
	if err != nil {
		os.Remove(newPath)
		return imagePath, fmt.Errorf("failed to get compressed file info: %w", err)
	}

	// 如果压缩后文件更大，使用原文件
	if compressedInfo.Size() >= fileInfo.Size() {
		log.Println("Compressed file is larger, keeping original")
		os.Remove(newPath)
		return imagePath, nil
	}

	// 压缩成功，删除原文件（如果扩展名变了）
	if newPath != imagePath {
		os.Remove(imagePath)
		log.Printf("Image format changed: %s -> %s", originalExt, ".jpg")
	}

	log.Printf("Image compressed: %d KB -> %d KB (saved %.1f%%)",
		fileInfo.Size()/1024,
		compressedInfo.Size()/1024,
		float64(fileInfo.Size()-compressedInfo.Size())*100/float64(fileInfo.Size()))

	return newPath, nil
}

// generateJsDelivrURL 生成 jsDelivr CDN 链接
func generateJsDelivrURL(remoteURL, branch, filePath string) string {
	// 从远程 URL 提取 user/repo
	var userRepo string

	remoteStr := strings.TrimSpace(remoteURL)
	remoteStr = strings.TrimSuffix(remoteStr, ".git")

	if strings.HasPrefix(remoteStr, "https://github.com/") {
		userRepo = strings.TrimPrefix(remoteStr, "https://github.com/")
	} else if strings.HasPrefix(remoteStr, "git@github.com:") {
		userRepo = strings.TrimPrefix(remoteStr, "git@github.com:")
		userRepo = strings.TrimSuffix(userRepo, ".git")
	} else {
		// 如果无法解析，返回空字符串，后续会使用 GitHub raw
		return ""
	}

	// jsDelivr 格式: https://cdn.jsdelivr.net/gh/user/repo@branch/path
	return fmt.Sprintf("https://cdn.jsdelivr.net/gh/%s@%s/images/%s",
		userRepo, branch, path.Base(filePath))
}

// generateImageURL 生成图片链接（支持 jsDelivr 和 GitHub raw）
func generateImageURL(remoteURL, branch, filePath string) string {
	if useJsDelivr {
		jsDelivrURL := generateJsDelivrURL(remoteURL, branch, filePath)
		if jsDelivrURL != "" {
			return jsDelivrURL
		}
		// 如果 jsDelivr 生成失败，fallback 到 GitHub raw
		log.Println("Warning: Failed to generate jsDelivr URL, using GitHub raw")
	}

	// GitHub raw 链接
	return remoteURL + "/blob/" + branch + "/images/" + path.Base(filePath) + "?raw=true"
}

// example: go run main.go https://i.stack.imgur.com/5W3rG.png
func main() {
	if len(os.Args) != 2 {
		log.Println("Usage: go run main.go <image_file_or_url>")
		os.Exit(1)
	}

	imageURL := os.Args[1]

	// Ensure images directory exists
	imagesDir := path.Join(getCurrentDir(), "images")
	if err := ensureDir(imagesDir); err != nil {
		log.Fatal("failed to create images directory: ", err)
	}

	// Get content type if it's a URL
	var contentType string
	if u, err := url.Parse(imageURL); err == nil && u.Scheme != "" && u.Host != "" {
		client := &http.Client{Timeout: 5 * time.Second}
		resp, err := client.Head(imageURL)
		if err == nil {
			contentType = resp.Header.Get("Content-Type")
			resp.Body.Close()
		}
	}

	tempFilePath := generateFileName(imageURL, contentType)
	err := downloadFile(imageURL, tempFilePath)
	if err != nil {
		log.Fatal("failed to download file: ", err)
	}

	log.Println("Downloaded file to: " + tempFilePath)

	// 压缩图片（直接覆盖原文件，时间戳部分不变，但扩展名可能变为 .jpg）
	tempFilePath, err = compressImage(tempFilePath)
	if err != nil {
		log.Printf("Warning: image compression failed: %v (continuing with original)", err)
		// 压缩失败不影响后续流程，继续使用原图
	}

	// Get current branch dynamically
	branch, err := getCurrentBranch()
	if err != nil {
		log.Println("Warning: failed to get current branch, using 'main'")
		branch = "main"
	}

	// BUG FIX: Use dynamic branch name
	currentDir := getCurrentDir()
	cmd := fmt.Sprintf("cd %s && git add . && git commit -m 'new image' && git push origin %s", currentDir, branch)

	err = execCommand(cmd)
	if err != nil {
		log.Fatal("failed to push to git: ", err)
	}

	// 生成图片链接（支持 jsDelivr 和 GitHub raw）
	remoteURL := getRemoteDir()
	uploadedImage := generateImageURL(remoteURL, branch, tempFilePath)

	// MD format of the image
	md := "![image](" + uploadedImage + ")"
	fmt.Println()
	fmt.Println()
	fmt.Println(md)
	fmt.Println()
	fmt.Println()

	// BUG FIX: Safe clipboard copy
	err = copyToClipboard(md)
	if err != nil {
		log.Fatal("can not copy to clipboard: ", err)
	}

	log.Println("Markdown link copied to clipboard!")
	if useJsDelivr {
		log.Println("Using jsDelivr CDN for faster loading")
	}
}
