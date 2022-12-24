package queueqr

import (
	"fmt"

	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
)

type qrTableRequest struct ***REMOVED***
	MerchantId string `json:merchantId`
	TableNumber int `json:tableNumber`
***REMOVED***

func handleErrResp(err error, c *gin.Context) ***REMOVED***
	if(err != nil) ***REMOVED***
		c.JSON(500, gin.H***REMOVED***
			"message": "Something went wrong",
			"error": err,
		***REMOVED***)
	***REMOVED***
***REMOVED***

func createQRTable(c *gin.Context) ***REMOVED***
	request := qrTableRequest***REMOVED******REMOVED***
	var err = c.BindJSON(&request)
	handleErrResp(err, c)
	
	png, err := qrcode.Encode("https://facebook.com", qrcode.High, 256)
	handleErrResp(err, c)
	
	fmt.Printf("request: %#v\n", request)
	var fileName = fmt.Sprintf("%s_%d_table_qr_code", request.MerchantId, request.TableNumber)
	
	fileId, err := SaveImageToGDrive(png, fileName)
	handleErrResp(err, c)
	
	c.JSON(200, gin.H***REMOVED***
		"message": "success to get QR Code",
		"file_id": fileId,
	***REMOVED***)
***REMOVED***

func QueueQR() ***REMOVED***
	r := gin.Default()
	r.POST("/qr-table", createQRTable)
	r.Run()
***REMOVED***