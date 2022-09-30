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

// DownloadFile will download an url to a local file.
func downloadFile(URL, fileName string) error {

	u, err := url.Parse(URL)
	if err != nil || u.Scheme == "" || u.Host == "" {
		//not a valid url, maybe local file
		if _, err := os.Stat(URL); errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("file %s does not exist", URL)
		} else {
			err := execCommand("cp " + URL + " " + fileName)
			if err != nil {
				return err
			}
			return nil
		}
	}

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
	cmd := "cd " + getCurrentDir() + " && git remote get-url origin"
	remote, err := execCommandAndReturnOutput(cmd)
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

func execCommandAndReturnOutput(command string) (string, error) {
	cmd := exec.Command("bash -c", command)
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

	uploadedImage := getRemoteDir() + "/blob/main/images/" + path.Base(tempFilePath) + "?raw=true"

	//MD format of the image
	md := "![image](" + uploadedImage + ")"
	fmt.Println()
	fmt.Println()
	fmt.Println(md)
	fmt.Println()
	fmt.Println()

	//copy to clipboard
	err = execCommand("echo \"" + md + "\" | pbcopy")
	if err != nil {
		log.Fatal("can not copy to clipboard ", err)
	}

}
