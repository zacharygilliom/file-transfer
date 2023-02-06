package main

import (
	"log"

	"github.com/zacharygilliom/file-transfer/internal/google"
	"github.com/zacharygilliom/file-transfer/internal/system"
)

func main() {
	pl, err := google.VerifyPhotosService()
	if err != nil {
		log.Fatal(err)
	}
	google.GetPhotos(pl)
	system.CreatePhotoDirectory("/home/zach/googlephotos")
}
