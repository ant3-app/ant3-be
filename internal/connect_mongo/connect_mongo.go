package connectmongo

import (
	"context"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Database ***REMOVED***
	uri := viper.Get("MONGODB_URI").(string)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	
	return client.Database("ant3")
	
***REMOVED***