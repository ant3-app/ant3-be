package main

import (
	queueqr "ant3/pkg/queue_qr"
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
	queueqr.QueueQR()
***REMOVED***