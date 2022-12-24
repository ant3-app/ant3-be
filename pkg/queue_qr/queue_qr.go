package queueqr

import (
	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
)

func testGet(c *gin.Context) ***REMOVED***
	png, err := qrcode.Encode("https://facebook.com", qrcode.High, 256)
	if err != nil ***REMOVED***
		c.JSON(500, gin.H***REMOVED***
			"message": "Something went wrong",
			"error": err,
		***REMOVED***)
	***REMOVED***
	
	fileId, err := SaveImageToGDrive(png, "table_qr_code.png")
	
	if err != nil ***REMOVED***
		c.JSON(500, gin.H***REMOVED***
			"message": "Something went wrong",
			"error": err,
		***REMOVED***)
	***REMOVED***
	
	
	c.JSON(200, gin.H***REMOVED***
		"message": "success to get QR Code",
		"file_id": fileId,
	***REMOVED***)
***REMOVED***

func QueueQR() ***REMOVED***
	r := gin.Default()
	r.GET("/test", testGet)
	r.Run()
***REMOVED***