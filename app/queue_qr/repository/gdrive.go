package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func serviceAccount(credentialFile string) *http.Client {
	cred, err := ioutil.ReadFile(credentialFile)
	if err != nil {
		log.Fatal(err)
	}
	var credStruct = struct {
		Email      string `json:"client_email"`
		PrivateKey string `json:"private_key"`
	}{}
	json.Unmarshal(cred, &credStruct)
	config := &jwt.Config{
		Email:      credStruct.Email,
		PrivateKey: []byte(credStruct.PrivateKey),
		Scopes: []string{
			drive.DriveScope,
		},
		TokenURL: google.JWTTokenURL,
	}
	client := config.Client(context.Background())
	return client
}

func createFile(client *http.Client, imgByte []byte, fileName string) (*drive.File, error){
	service, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to create Drive service: %v", err)
	}

	gdriveFolderId := viper.Get("GDRIVE_FOLDER_ID").(string)

	img := bytes.NewReader(imgByte)
	if err != nil {
		log.Fatalf("error opening %q: %v", fileName, err)
	}
	driveFile, err := service.Files.
		Create(&drive.File{Name: fileName, Parents: []string{gdriveFolderId}}).
		Media(img).
		Do()
	if(err!=nil) {
		fmt.Println(err, "error ma bro")
	}
	log.Printf("Got drive.File, response: %#v", driveFile)
	return driveFile, err
}

func SaveImageToGDrive(imgByte []byte, fileName string) (string, error) {
	client := serviceAccount("service-account.json")
	driveFile, err := createFile(client, imgByte, fileName)
	return driveFile.Id, err
}