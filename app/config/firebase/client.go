package fb

import (
	"context"
	"fmt"

	firestore "cloud.google.com/go/firestore"
)

var Client *firestore.Client

func InitClient() error ***REMOVED***
	InitFirebase()
	client, err := App.Firestore(context.Background())
	if(err != nil) ***REMOVED***
		fmt.Println(err, "[InitClient] error init firebase client")
		return err
	***REMOVED***
	Client = client
	return nil
***REMOVED***