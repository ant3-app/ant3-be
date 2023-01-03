package fb

import (
	"context"
	"fmt"

	db "firebase.google.com/go/v4/db"
	"github.com/spf13/viper"
)

var Client *db.Client

func InitClient() error ***REMOVED***
	InitFirebase()
	client, err := App.DatabaseWithURL(context.Background(), viper.GetString("FIREBASE_REALTIME_DATABASE"))
	if(err != nil) ***REMOVED***
		fmt.Println(err, "[InitClient] error init firebase client")
		return err
	***REMOVED***
	Client = client
	return nil
***REMOVED***