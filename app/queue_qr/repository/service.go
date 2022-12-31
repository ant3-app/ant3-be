package repository

import (
	mongo "ant3/app/config/mongo"
	"ant3/app/queue_qr/models"
	"context"
)

func Save(queueQr *models.QueueQR) ***REMOVED***
	collection := mongo.DB.Collection("queue_qr")

	collection.InsertOne(context.Background(), queueQr)
	
	println("success to insert?")
***REMOVED***