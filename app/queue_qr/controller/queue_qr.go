package controller

import (
	"ant3/app/queue_qr/models"
	repository "ant3/app/queue_qr/repository"
	service "ant3/app/queue_qr/service"
	"fmt"

	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	
	qrTableId := primitive.NewObjectID()
	qrTableWebLink := fmt.Sprintf("%s/qr-table/%s", viper.Get("ANT3_WEB_URL"), qrTableId.Hex())
	
	png, err := qrcode.Encode(qrTableWebLink, qrcode.High, 256)
	handleErrResp(err, c)
	
	fmt.Printf("request: %#v\n", request)
	var fileName = fmt.Sprintf("%s_%d_table_qr_code", request.MerchantId, request.TableName)
	
	fileId, err := service.SaveImageToGDrive(png, fileName)
	handleErrResp(err, c)
	oid, err := primitive.ObjectIDFromHex(request.MerchantId)
	handleErrResp(err, c)
	
	var queueQr *models.QueueQR = &models.QueueQR***REMOVED***
		MerchantId: oid,
		Name: request.TableName,
		FileId: fileId,
		Id: qrTableId,
	***REMOVED***
	
	fmt.Printf("trying to save with %#v\n", queueQr)
	repository.Save(queueQr)
	
	c.JSON(200, gin.H***REMOVED***
		"message": "success to get QR Code",
		"fileId": fileId,
		"merchantId": oid,
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
***REMOVED***