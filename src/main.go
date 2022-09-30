package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strconv"
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

func generateFileName(url string) string {

	fileName := path.Base(url)

	return path.Join(getCurrentDir()+"/images", strconv.FormatInt(time.Now().UnixNano(), 10)+"-"+fileName)
}

func execCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {

	//if len(os.Args) != 2 {
	//	log.Println("Usage: main.go <image_file>")
	//	os.Exit(1)
	//}
	//url := os.Args[1]
	//tempFilePath := generateFileName(url)
	//err := downloadFile(url, tempFilePath)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println("Downloaded file to: " + tempFilePath)

	//commit images to github
	//git add .
	//git commit -m "add images"
	//git push origin master

	err := execCommand("cd", getCurrentDir()+"/images")
	if err != nil {
		log.Fatal(err)
	}
	err = execCommand("git", "add", ".")
	if err != nil {
		log.Fatal(err)
	}
	err = execCommand("git", "commit", "-m", "add images")
	if err != nil {
		log.Fatal(err)
	}
	err = execCommand("git", "push", "origin", "master")
	if err != nil {
		log.Fatal(err)
	}

}
