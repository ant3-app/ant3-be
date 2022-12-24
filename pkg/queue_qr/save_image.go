package queueqr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func ServiceAccount(credentialFile string) *http.Client ***REMOVED***
	cred, err := ioutil.ReadFile(credentialFile)
	if err != nil ***REMOVED***
		log.Fatal(err)
	***REMOVED***
	var credStruct = struct ***REMOVED***
		Email      string `json:"client_email"`
		PrivateKey string `json:"private_key"`
	***REMOVED******REMOVED******REMOVED***
	json.Unmarshal(cred, &credStruct)
	config := &jwt.Config***REMOVED***
		Email:      credStruct.Email,
		PrivateKey: []byte(credStruct.PrivateKey),
		Scopes: []string***REMOVED***
			drive.DriveScope,
		***REMOVED***,
		TokenURL: google.JWTTokenURL,
	***REMOVED***
	client := config.Client(context.Background())
	return client
***REMOVED***

func createFile(client *http.Client, imgByte []byte) (*drive.File, error)***REMOVED***
	// service, err := drive.New(client)
	service, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil ***REMOVED***
		log.Fatalf("Unable to create Drive service: %v", err)
	***REMOVED***

	filename := "my_qr_code.jpeg"

	img := bytes.NewReader(imgByte)
	if err != nil ***REMOVED***
		log.Fatalf("error opening %q: %v", filename, err)
	***REMOVED***
	driveFile, err := service.Files.Create(&drive.File***REMOVED***Name: filename, Parents: []string***REMOVED***"id"***REMOVED******REMOVED***).Media(img).Do()
	if(err!=nil) ***REMOVED***
		fmt.Println(err, "error ma bro")
	***REMOVED***
	log.Printf("Got drive.File, response: %#v", driveFile)
	return driveFile, err
***REMOVED***

func SaveImageToGDrive(imgByte []byte) (*drive.File, error) ***REMOVED***
	client := ServiceAccount("client_secret_gdrive.json")
	return createFile(client, imgByte)
***REMOVED***