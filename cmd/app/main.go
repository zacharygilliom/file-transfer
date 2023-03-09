package main

import (
	"fmt"
	"log"

	"github.com/zacharygilliom/file-transfer/internal/google"
	"github.com/zacharygilliom/file-transfer/internal/system"
	"github.com/zacharygilliom/file-transfer/internal/transfer"
)

func main() {
	pl, err := google.VerifyPhotosService()
	fmt.Println("verified")
	if err != nil {
		log.Fatal(err)
	}
	var latestPhotoDate system.PhotosDateTime
	dirName := "/home/zach/googlephotos"
	latestPhotoDate = system.GetCurrentPhotoIds(dirName)
	system.CreatePhotoDirectory(dirName)
	items := google.GetPhotos(pl, latestPhotoDate)
	transfer.GetFiles(items, dirName)
	fmt.Println("Program Completed")
}
