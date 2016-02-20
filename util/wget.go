package util

import (
	"io"
	"log"
	"net/http"
	"os"
)

func Wget(url, fileName string) error {
	if !Exists(fileName) {
		log.Println("Downloading", url, "to", fileName)

		output, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer output.Close()

		response, err := http.Get(url)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		n, err := io.Copy(output, response.Body)
		if err != nil {
			return err
		}

		log.Println(n, "bytes downloaded.")
	}
	return nil
}
