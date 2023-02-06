package transfer

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

// GetFiles takes in a list of urls and downloads and saves each one.
func GetFiles(URLs []string, dir string) {
	for i, v := range URLs {
		fmt.Println(v)
		name := dir + "/photo_" + strconv.Itoa(i) + ".pdf"
		DownloadFile(v, name)
	}
}

//DownloadFile takes in a pictures url and a name for the new file and downloads it into the directory.
func DownloadFile(URL, fileName string) error {
	response, err := http.Get(URL)
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
