package system

import (
	"fmt"
	"log"
	"os"
)

// CreatePhotoDirectory takes in a user's specified directory and checks if it exists and if not it will create it.
func CreatePhotoDirectory(directory string) {
	folderInfo, err := os.Stat(directory)
	if os.IsNotExist(err) {
		err := os.Mkdir(directory, 0755)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Directory %v Successfully Created\n", directory)
	} else {
		fmt.Printf("Directory %v Already Exists\n", folderInfo.Name())
	}
}
