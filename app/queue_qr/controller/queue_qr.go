package controller

import (
	service "ant3/app/queue_qr/service"

	"github.com/gin-gonic/gin"
)

func CreateQRTable(c *gin.Context) ***REMOVED***
	request := qrTableRequest***REMOVED******REMOVED***
	var err = c.BindJSON(&request)
	if(err != nil) ***REMOVED***
		c.AbortWithStatusJSON(500, gin.H***REMOVED***
			"message": "Something went wrong",
			"error": err.Error(),
		***REMOVED***)
		return
	***REMOVED***
	
	res, err := service.SaveQrTable(service.SaveQrTableRequest***REMOVED***
		MerchantId: request.MerchantId,
		TableName: request.TableName,
	***REMOVED***)
	if(err != nil) ***REMOVED***
		c.AbortWithStatusJSON(500, gin.H***REMOVED***
			"message": "Something went wrong",
			"error": err.Error(),
		***REMOVED***)
		return
	***REMOVED***
	
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
	if(err != nil) ***REMOVED***
		c.AbortWithStatusJSON(500, gin.H***REMOVED***
			"message": "Something went wrong",
			"error": err.Error(),
		***REMOVED***)
		return
	***REMOVED***
	
	key, err := service.AddQRTableToQueue(request.QueueQrId)
	if(err != nil) ***REMOVED***
		c.AbortWithStatusJSON(400, gin.H***REMOVED***
			"message": "Something went wrong",
			"error": err.Error(),
		***REMOVED***)
		return
	***REMOVED***
	c.JSON(200, gin.H***REMOVED***
		"message": "success to add the table to the queue",
		"key": key,
	***REMOVED***)
***REMOVED***