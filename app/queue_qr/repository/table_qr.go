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

func Save(queueQr *models.TableQr) (*string, error) {
	collection := mongo.DB.Collection("table_qr")

	res, err := collection.InsertOne(context.Background(), queueQr)
	if (err != nil) {
		return nil, err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return &id, nil
}

func GetOne(queueQrId string) (*models.TableQrDTO, error) {
	collection := mongo.DB.Collection("table_qr")
	fmt.Println("[GetOne] mongo collection with id:", queueQrId)
	id, err := primitive.ObjectIDFromHex(queueQrId)
	
	if(err != nil) {
		fmt.Println(err.Error(), "[GetOne] error when get the data from mongo")
		return nil, err
	}
	res := collection.FindOne(context.Background(), bson.M{"_id": id})
	var queueQr models.TableQrDTO
	res.Decode(&queueQr)
	if(queueQr.Id == "") {
		return nil, errors.New("data not found")
	}
	return &queueQr, nil
}  