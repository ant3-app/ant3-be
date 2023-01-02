package main

import (
	mongo "ant3/app/config/mongo"
	queue_qr "ant3/app/queue_qr/controller"
	"log"

	"github.com/gin-gonic/gin"
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
	r := gin.Default()
	r.POST("/qr-table", queue_qr.CreateQRTable)
	r.Run()
***REMOVED***