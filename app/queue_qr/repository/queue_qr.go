package repository

import (
	mongo "ant3/app/config/mongo"
	"ant3/app/queue_qr/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// var collection *mongodriver.Collection

func Save(queueQr *models.QueueQR) ***REMOVED***
	collection := mongo.DB.Collection("queue_qr")

	collection.InsertOne(context.Background(), queueQr)
	
***REMOVED***

func GetOne(queueQrId string) (*models.QueueQrDTO, error) ***REMOVED***
	collection := mongo.DB.Collection("queue_qr")
	fmt.Println("[GetOne] mongo collection with id:", queueQrId)
	id, err := primitive.ObjectIDFromHex(queueQrId)
	if(err != nil) ***REMOVED***
		fmt.Println(err, "[GetOne] error when get the data from mongo")
		return nil, err
	***REMOVED***
	res := collection.FindOne(context.Background(), bson.M***REMOVED***"_id": id***REMOVED***)
	var queueQr models.QueueQrDTO
	res.Decode(&queueQr)
	return &queueQr, nil
***REMOVED***  