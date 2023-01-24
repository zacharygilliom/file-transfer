package gmail

import (
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

//VerifyGmail will start our api connection to google photos
func VerifyGmail() {

	content, err := ioutil.ReadFile("../../config/credentials.json")
	if err != nil {
		log.Fatal(err)
	}

	//TODO: Actually figure out how to do authorization.  Something is wrong here and not allowing the authorization to go through
	conf, err := google.ConfigFromJSON(content)
	url := conf.AuthCodeURL("state")
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	tok, err := conf.Exchange(oauth2.NoContext, "authorization-code")
	if err != nil {
		log.Fatal(err)
	}
	client := conf.Client(oauth2.NoContext, tok)
	client.Get("...")

}
