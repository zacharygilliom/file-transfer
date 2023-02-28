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
	items := google.GetPhotos(pl)
	c := make(chan google.Photos, len(items))
	for _, a := range items {
		c <- a
	}
	close(c)
	dirName := "/home/zach/googlephotos"
	system.CreatePhotoDirectory(dirName)
	transfer.GetFiles(c, dirName)
	fmt.Println("program closed")
}
