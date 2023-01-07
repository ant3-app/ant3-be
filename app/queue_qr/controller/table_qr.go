package controller

import (
	service "ant3/app/queue_qr/service"

	"github.com/gin-gonic/gin"
)

func CreatetableQr(c *gin.Context) {
	request := tableQrRequest{}
	var err = c.BindJSON(&request)
	if(err != nil) {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Something went wrong",
			"error": err.Error(),
		})
		return
	}
	
	res, err := service.SaveTableQr(service.SaveTableQrRequest{
		MerchantId: request.MerchantId,
		TableName: request.TableName,
	})
	if(err != nil) {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Something went wrong",
			"error": err.Error(),
		})
		return
	}
	
	c.JSON(200, gin.H{
		"message": "success to get QR Code",
		"fileId": res.FileId,
		"merchantId": res.MerchantId,
		"id": res.Id,
	})
}

func GetTableQrQueueInfo(c *gin.Context) {
	// tableQr := c.Param("id")
	
}

func InsertTableQrToQueue(c *gin.Context) {
	request := tableQueueRequest{}
	var err = c.BindJSON(&request)
	if(err != nil) {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Something went wrong",
			"error": err.Error(),
		})
		return
	}
	
	key, err := service.AddTableQrToQueue(request.TableQrId)
	if(err != nil) {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "Something went wrong",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success to add the table to the queue",
		"key": key,
	})
}

func RemoveTableQrFromQueue(c *gin.Context) {
	request := tableQueueRequest{}
	var err = c.BindJSON(&request)
	if(err != nil) {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Something went wrong",
			"error": err.Error(),
		})
		return
	}
	key, err := service.RemoveTableFromQueue(request.TableQrId)
	if(err != nil) {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "Something went wrong",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success to remove the table from the queue",
		"key": key,
	})
}