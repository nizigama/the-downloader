package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/fatih/color"
)

var downloadLink string
var fileName string
var chunkSize int

func init() {
	flag.StringVar(&downloadLink, "downloadLink", "https://avatars.githubusercontent.com/u/65668685?s=40&v=4", "The link of the file you want to download")
	flag.StringVar(&fileName, "fileName", "gopher.png", "The name of the file you want to download")
	flag.IntVar(&chunkSize, "chunkSize", 500, "The size of a single chunk of data to download")
	flag.Parse()
}

func main() {

	client := &http.Client{}

	req, _ := http.NewRequest("GET", downloadLink, nil)

	sizeInBytes, err := getFileDownloadSizeInBytes(client, req)

	if err != nil {
		color.Red(err.Error())
		return
	}

	err = downloadFileInChunks(client, req, sizeInBytes-1)

	if err != nil {
		color.Red(err.Error())
		return
	}
	color.Green("Download completed successfully")

}

func downloadFileInChunks(client *http.Client, req *http.Request, bytesNeeded int) error {

	received := 0
	nextChunkSize := chunkSize
	start := 0
	end := start + nextChunkSize

	consolidatedFile, err := os.Create(fileName)

	if err != nil {
		return fmt.Errorf("failed creating local consolidated file\nError: %s", err.Error())
	}

	defer consolidatedFile.Close()

	for received < bytesNeeded {
		req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", start, end))

		resp, err := client.Do(req)

		if err != nil {
			return fmt.Errorf("failed download a portion of the file\nError: %s", err.Error())
		}

		if resp.StatusCode >= 300 {
			return fmt.Errorf("invalid response code: %v", resp.StatusCode)
		}

		_, err = io.Copy(consolidatedFile, resp.Body)

		if err != nil {
			return fmt.Errorf("failed reading body content\nError: %s", err.Error())
		}

		resp.Body.Close()

		received += nextChunkSize

		start = end + 1

		if bytesNeeded-received < nextChunkSize {
			nextChunkSize = bytesNeeded - received
		}

		end = received + nextChunkSize
	}

	return nil
}

func getFileDownloadSizeInBytes(client *http.Client, req *http.Request) (int, error) {
	resp, err := client.Do(req)

	if err != nil {
		return 0, fmt.Errorf("failed connecting to the download link\nError: %s", err.Error())
	}

	defer resp.Body.Close()

	fileSize, err := strconv.Atoi(resp.Header.Get("Content-Length"))

	if err != nil {
		return 0, fmt.Errorf("invalid file content length")
	}

	return fileSize, nil
}
