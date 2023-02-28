package transfer

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/zacharygilliom/file-transfer/internal/google"
)

// GetFiles takes in a list of urls and downloads and saves each one.
func GetFiles(c chan google.Photos, dir string) {
	// TODO: Turn this into a goroutine that concurrently downloads the files
	fmt.Println("GetFiles called")
	for n := range c {
		name := dir + "/" + n.Name
		DownloadFile(n.URL, name, "=w2048-h1024")
		fmt.Println("Download Completed")
	}

}

//DownloadFile takes in a pictures url and a name for the new file and downloads it into the directory.
func DownloadFile(URL, fileName, extension string) error {
	response, err := http.Get(URL + extension)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return errors.New("Received non 200 response type code")
	}
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil

}
