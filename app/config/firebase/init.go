package fb

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
)

var App *firebase.App

func InitFirebase() ***REMOVED***
	config := &firebase.Config***REMOVED***ProjectID: viper.GetString("FIREBASE_PROJECT_ID")***REMOVED***
	opt := option.WithCredentialsFile("./service-account.json")

	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil ***REMOVED***
		log.Fatalf("error initializing app: %v\n", err)
	***REMOVED***
	
	App = app
***REMOVED***