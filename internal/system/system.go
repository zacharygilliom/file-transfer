package system

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// PhotosDateTime holds the latest photo date
type PhotosDateTime struct {
	Year  int64
	Month int64
	Day   int64
}

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

// GetCurrentPhotoIds will return the last day of creation that photos are in the directory
func GetCurrentPhotoIds(directory string) PhotosDateTime {
	items, _ := ioutil.ReadDir(directory)
	var latestDate = PhotosDateTime{1900, 1, 1}
	for _, item := range items {
		fileparts := strings.Split(item.Name(), "_")
		filePartsCreationTime := fileparts[0]
		dayTimePartSplit := strings.Split(filePartsCreationTime, "T")
		dayPart := dayTimePartSplit[0]
		dayPartSplit := strings.Split(dayPart, "-")
		year, err := strconv.ParseInt(dayPartSplit[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		month, err := strconv.ParseInt(dayPartSplit[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		day, err := strconv.ParseInt(dayPartSplit[2], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		var currentDate = PhotosDateTime{
			year,
			month,
			day,
		}
		fmt.Printf("%v-%v-%v\n", currentDate.Year, currentDate.Month, currentDate.Day)
		if currentDate.Year >= latestDate.Year {
			if currentDate.Month >= latestDate.Month {
				if currentDate.Day > latestDate.Day {
					latestDate.Year = currentDate.Year
					latestDate.Month = currentDate.Month
					latestDate.Day = currentDate.Day
				} else if currentDate.Day == latestDate.Day {
					latestDate = currentDate
				} else {
					latestDate = currentDate
				}
			} else {
				latestDate = currentDate
			}
		}
	}
	return latestDate
}
