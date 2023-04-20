package system

import (
	"fmt"
	"io/fs"
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

func getPhotosDateTime(item fs.FileInfo) (PhotosDateTime, error) {
	fileParts := strings.Split(item.Name(), "_")
	filePartsCreationTime := fileParts[0]
	dayTimePartSplit := strings.Split(filePartsCreationTime, "T")
	dayPart := dayTimePartSplit[0]
	dayPartSplit := strings.Split(dayPart, "-")
	var currentDate PhotosDateTime
	year, err := strconv.ParseInt(dayPartSplit[0], 10, 64)
	if err != nil {
		return currentDate, err
	}
	currentDate.Year = year
	month, err := strconv.ParseInt(dayPartSplit[1], 10, 64)
	if err != nil {
		return currentDate, err
	}
	currentDate.Month = month
	day, err := strconv.ParseInt(dayPartSplit[2], 10, 64)
	if err != nil {
		return currentDate, err
	}
	currentDate.Day = day
	return currentDate, nil
}

// GetLatestDate finds the date of the most recent file in the directory
func GetLatestDate(currentDate, latestDate PhotosDateTime) PhotosDateTime {
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
	return latestDate
}

// GetCurrentPhotoIds will return the last day of creation that photos are in the directory
func GetCurrentPhotoIds(directory string) PhotosDateTime {
	items, _ := ioutil.ReadDir(directory)
	var latestDate = PhotosDateTime{1900, 1, 1}
	for _, item := range items {
		currentDate, err := getPhotosDateTime(item)
		fmt.Printf("%v-%v-%v\n", currentDate.Year, currentDate.Month, currentDate.Day)
		if err != nil {
			log.Fatal(err)
		}
		// currentDate = {1901, 1, 1}, latestDate = {1900, 1, 1} should return {1901, 1, 1}
		// currentDate = {1899, 1, 1}, latestDate = {1900, 1, 1} should return {1900, 1, 1}
		latestDate = GetLatestDate(currentDate, latestDate)
	}
	return latestDate
}
