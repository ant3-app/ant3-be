package queueqr

import (
	"ant3/models"
	"context"
)

func Save(queueQr *models.QueueQR) ***REMOVED***
	collection := models.DB.Collection("queue_qr")

	collection.InsertOne(context.Background(), queueQr)
	
	println("success to insert?")
***REMOVED***