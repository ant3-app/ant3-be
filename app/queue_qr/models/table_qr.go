package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoTimeStamp struct {
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

type TableQr struct {
	Id primitive.ObjectID `bson:"_id"`
	MerchantId primitive.ObjectID `bson:"merchantId"`
	Name string `bson:"name"`
	FileId string `bson:"fileId"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

type TableQrDTO struct {
	Id string `bson:"_id" json:"id"`
	MerchantId string `json:"merchantId"`
	Name string `json:"name"`
	FileId string `json:"fileId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *TableQr) MarshalBSON() ([]byte, error) {
	if u.CreatedAt.IsZero() {
			u.CreatedAt = time.Now()
	}
	u.UpdatedAt = time.Now()
	
	type my TableQr
	return bson.Marshal((*my)(u))
}

