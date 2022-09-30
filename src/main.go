package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"time"
)

// DownloadFile will download an url to a local file.
func downloadFile(URL, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(response.Body)

	if response.StatusCode != 200 {
		return errors.New("got response code " + string(rune(response.StatusCode)))
	}
	//Create an empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	//Write the bytes to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func getCurrentDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return path.Join(home, "Documents/github/imagehosting/")
}

func getRemoteDir() string {
	remote, err := execCommandAndReturnOutput("git", "remote", "get-url", "origin")
	if err != nil {
		log.Fatal("can not get remote url: ", err)
	}
	remote = strings.TrimSuffix(remote, ".git\n")
	return remote
}

func generateFileName(url string) string {

	fileName := path.Base(url)

	return path.Join(getCurrentDir()+"/images", strconv.FormatInt(time.Now().UnixNano(), 10)+"-"+fileName)
}

//func execCommand(command string, args ...string) error {
//	cmd := exec.Command(command, args...)
//	cmd.Stdout = os.Stdout
//	cmd.Stderr = os.Stderr
//	return cmd.Run()
//}

func execCommand(command string) error {
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func execCommandAndReturnOutput(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	out, err := cmd.Output()
	return string(out), err
}

// example: go run main.go https://i.stack.imgur.com/5W3rG.png
func main() {

	if len(os.Args) != 2 {
		log.Println("Usage: main.go <image_file>")
		os.Exit(1)
	}
	url := os.Args[1]
	tempFilePath := generateFileName(url)
	err := downloadFile(url, tempFilePath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Downloaded file to: " + tempFilePath)

	cmd := "cd " + getCurrentDir() + " && git add . && git commit -m 'new image' && git push origin main"

	err = execCommand(cmd)
	if err != nil {
		log.Fatal(err)
	}

	uploadedImage := getRemoteDir() + "/image/" + path.Base(tempFilePath)

	//MD format of the image
	md := "![image](" + uploadedImage + ")"
	fmt.Println()
	fmt.Println()
	fmt.Println(md)
	fmt.Println()
	fmt.Println()

	//copy to clipboard
	err = execCommand("echo " + md + " | pbcopy")
	if err != nil {
		log.Fatal("can not copy to clipboard ", err)
	}

}
