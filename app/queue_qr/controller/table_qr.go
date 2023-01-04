package controller

import (
	service "ant3/app/queue_qr/service"

	"github.com/gin-gonic/gin"
)

func CreatetableQr(c *gin.Context) ***REMOVED***
	request := tableQrRequest***REMOVED******REMOVED***
	var err = c.BindJSON(&request)
	if(err != nil) ***REMOVED***
		c.AbortWithStatusJSON(500, gin.H***REMOVED***
			"message": "Something went wrong",
			"error": err.Error(),
		***REMOVED***)
		return
	***REMOVED***
	
	res, err := service.SaveTableQr(service.SaveTableQrRequest***REMOVED***
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

func GetTableQrQueueInfo(c *gin.Context) ***REMOVED***
	// tableQr := c.Param("id")
	
***REMOVED***

func InsertTableQrToQueue(c *gin.Context) ***REMOVED***
	request := tableQueueRequest***REMOVED******REMOVED***
	var err = c.BindJSON(&request)
	if(err != nil) ***REMOVED***
		c.AbortWithStatusJSON(500, gin.H***REMOVED***
			"message": "Something went wrong",
			"error": err.Error(),
		***REMOVED***)
		return
	***REMOVED***
	
	key, err := service.AddTableQrToQueue(request.TableQrId)
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

func RemoveTableQrFromQueue(c *gin.Context) ***REMOVED***
	request := tableQueueRequest***REMOVED******REMOVED***
	var err = c.BindJSON(&request)
	if(err != nil) ***REMOVED***
		c.AbortWithStatusJSON(500, gin.H***REMOVED***
			"message": "Something went wrong",
			"error": err.Error(),
		***REMOVED***)
		return
	***REMOVED***
	key, err := service.RemoveTableFromQueue(request.TableQrId)
	if(err != nil) ***REMOVED***
		c.AbortWithStatusJSON(400, gin.H***REMOVED***
			"message": "Something went wrong",
			"error": err.Error(),
		***REMOVED***)
		return
	***REMOVED***
	c.JSON(200, gin.H***REMOVED***
		"message": "success to remove the table from the queue",
		"key": key,
	***REMOVED***)
***REMOVED***