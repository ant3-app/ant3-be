package fb

import (
	"context"
	"fmt"

	firestore "cloud.google.com/go/firestore"
)

var Client *firestore.Client

func InitClient() error {
	InitFirebase()
	client, err := App.Firestore(context.Background())
	if(err != nil) {
		fmt.Println(err, "[InitClient] error init firebase client")
		return err
	}
	Client = client
	return nil
}