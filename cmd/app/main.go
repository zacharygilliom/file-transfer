package main

import (
	"log"

	"github.com/zacharygilliom/file-transfer/internal/google"
	"github.com/zacharygilliom/file-transfer/internal/system"
	"github.com/zacharygilliom/file-transfer/internal/transfer"
)

func main() {
	pl, err := google.VerifyPhotosService()
	if err != nil {
		log.Fatal(err)
	}
	photos := google.GetPhotos(pl)
	dirName := "/home/zach/googlephotos"
	system.CreatePhotoDirectory(dirName)
	transfer.GetFiles(photos, dirName)
	//transfer.DownloadFile(photos)
}
