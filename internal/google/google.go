package google

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gphotosuploader/googlemirror/api/photoslibrary/v1"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Photos holds all the data returned from getMediaItems
type Photos struct {
	Name     string
	URL      string
	MimeType string
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func getClient(config *oauth2.Config) *http.Client {
	tokFile := "config/token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		log.Fatalf("Unable to cache OAuth token: %v", err)
	}
	json.NewEncoder(f).Encode(token)
}

// Requests a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

//VerifyPhotosService will start our api connection to google photos
func VerifyPhotosService() (*photoslibrary.Service, error) {
	configFile, err := os.ReadFile("config/credentials.json")
	if err != nil {
		log.Fatal(err)
	}
	config, err := google.ConfigFromJSON(configFile, "https://www.googleapis.com/auth/photoslibrary.readonly")
	if err != nil {
		log.Fatal("Unable to parse credentials.json file")
	}

	client := getClient(config)

	photosService, err := photoslibrary.New(client)

	return photosService, err

}

func getMediaItems(p *photoslibrary.Service, searchFilters photoslibrary.SearchMediaItemsRequest, rA []Photos) (string, []Photos) {
	// TODO: fix how the goroutine is unbuffered and wont accept more data until it is taken
	var nextPageToken string
	mItems := p.MediaItems
	searchParams := mItems.Search(&searchFilters)
	result, err := searchParams.Do()
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range result.MediaItems {
		var photo = Photos{v.Filename, v.BaseUrl, v.MimeType}
		rA = append(rA, photo)
	}
	nextPageToken = result.NextPageToken
	return nextPageToken, rA
}

// GetPhotos returns an array of downloaded urls
func GetPhotos(pl *photoslibrary.Service) []Photos {
	var resultsArray []Photos
	var nextPageToken string
	searchFilters := photoslibrary.SearchMediaItemsRequest{PageSize: 50}
	nextPageToken, resultsArray = getMediaItems(pl, searchFilters, resultsArray)
	for nextPageToken != "" {
		searchFilters.PageToken = nextPageToken
		nextPageToken, resultsArray = getMediaItems(pl, searchFilters, resultsArray)
	}
	return resultsArray
}
