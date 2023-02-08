package transfer

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
)

// GetFiles takes in a list of urls and downloads and saves each one.
func GetFiles(URLs map[string][]string, dir string) {
	// TODO: Implement goroutines here
	for i, v := range URLs["photos"] {
		name := dir + "/photo_" + strconv.Itoa(i) + ".jpeg"
		DownloadFile(v, name, "=w2048-h1024")
	}
	for j, p := range URLs["videos"] {
		name := dir + "/photo_" + strconv.Itoa(j) + ".mp4"
		DownloadFile(p, name, "=dv")
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
