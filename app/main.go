package main

import (
	mongo "ant3/app/config/mongo"
	controller "ant3/app/queue_qr/controller"
	"log"

	"github.com/spf13/viper"
)

func initLoadEnv() ***REMOVED***
	viper.AddConfigPath("../../")
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil ***REMOVED***
    log.Fatalf("Error while reading config file %s", err)
  ***REMOVED***
***REMOVED***

func main() ***REMOVED***
	initLoadEnv()
	mongo.ConnectDatabase()
	controller.QueueQR()
***REMOVED***