package transfer

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/zacharygilliom/file-transfer/internal/google"
)

// GetFiles takes in a list of urls and downloads and saves each one.
func GetFiles(c chan google.Photos, dir string) {
	// TODO: Turn this into a goroutine that concurrently downloads the files
	fmt.Println("GetFiles called")
	var wg sync.WaitGroup
	for n := range c {
		wg.Add(1)
		go DownloadFile(n, &wg, dir)
	}
	wg.Wait()
}

//DownloadFile takes in a pictures url and a name for the new file and downloads it into the directory.
func DownloadFile(n google.Photos, wg *sync.WaitGroup, dir string) error {
	defer wg.Done()
	response, err := http.Get(n.URL + "=w2048-h1024")
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return errors.New("Received non 200 response type code")
	}
	file, err := os.Create(dir + "/" + n.CreationTime + "_" + n.Name)
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
