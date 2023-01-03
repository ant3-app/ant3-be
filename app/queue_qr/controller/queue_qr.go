package controller

import (
	service "ant3/app/queue_qr/service"

	"github.com/gin-gonic/gin"
)

func handleErrResp(err error, c *gin.Context) ***REMOVED***
	if(err != nil) ***REMOVED***
		c.JSON(500, gin.H***REMOVED***
			"message": "Something went wrong",
			"error": err,
		***REMOVED***)
	***REMOVED***
***REMOVED***

func CreateQRTable(c *gin.Context) ***REMOVED***
	request := qrTableRequest***REMOVED******REMOVED***
	var err = c.BindJSON(&request)
	handleErrResp(err, c)
	
	res, err := service.SaveQrTable(service.SaveQrTableRequest***REMOVED***
		MerchantId: request.MerchantId,
		TableName: request.TableName,
	***REMOVED***)
	handleErrResp(err, c)
	
	c.JSON(200, gin.H***REMOVED***
		"message": "success to get QR Code",
		"fileId": res.FileId,
		"merchantId": res.MerchantId,
		"id": res.Id,
	***REMOVED***)
***REMOVED***

func GetQRTableQueueInfo(c *gin.Context) ***REMOVED***
	// qrTable := c.Param("id")
	
***REMOVED***

func InsertQRTableToQueue(c *gin.Context) ***REMOVED***
	request := queueTableRequest***REMOVED******REMOVED***
	var err = c.BindJSON(&request)
	handleErrResp(err, c)
	
	service.AddQRTableToQueue(request.QueueQrId)
	// c.JSON(200, gin.H***REMOVED***
	// 	"message": "success to add the table to the queue",
	// 	"fileId": res.FileId,
	// 	"merchantId": res.MerchantId,
	// 	"id": res.Id,
	// ***REMOVED***)
***REMOVED***