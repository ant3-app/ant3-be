package repository

import (
	mongo "ant3/app/config/mongo"
	"ant3/app/queue_qr/models"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// var collection *mongodriver.Collection

func Save(queueQr *models.QueueQR) (*string, error) ***REMOVED***
	collection := mongo.DB.Collection("queue_qr")

	res, err := collection.InsertOne(context.Background(), queueQr)
	if (err != nil) ***REMOVED***
		return nil, err
	***REMOVED***
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return &id, nil
***REMOVED***

func GetOne(queueQrId string) (*models.QueueQrDTO, error) ***REMOVED***
	collection := mongo.DB.Collection("queue_qr")
	fmt.Println("[GetOne] mongo collection with id:", queueQrId)
	id, err := primitive.ObjectIDFromHex(queueQrId)
	
	if(err != nil) ***REMOVED***
		fmt.Println(err.Error(), "[GetOne] error when get the data from mongo")
		return nil, err
	***REMOVED***
	res := collection.FindOne(context.Background(), bson.M***REMOVED***"_id": id***REMOVED***)
	var queueQr models.QueueQrDTO
	res.Decode(&queueQr)
	if(queueQr.Id == "") ***REMOVED***
		return nil, errors.New("data not found")
	***REMOVED***
	return &queueQr, nil
***REMOVED***  