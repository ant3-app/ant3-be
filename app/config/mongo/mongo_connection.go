package mongo

import (
	"context"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDatabase() ***REMOVED***
	uri := viper.Get("MONGODB_URI").(string)
	println(uri, "uri")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	
	DB = client.Database("ant3")
***REMOVED***