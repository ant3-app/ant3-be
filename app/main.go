package main

import (
	fb "ant3/app/config/firebase"
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
	fb.InitClient()
	
	r := gin.Default()
	r.POST("/qr-table", queue_qr.CreatetableQr)
	r.GET("/qr-table/:id", queue_qr.GetTableQrQueueInfo)
	r.POST("qr-table/queue", queue_qr.InsertTableQrToQueue)
	r.DELETE("qr-table/queue", queue_qr.RemoveTableQrFromQueue)
	r.Run()
***REMOVED***