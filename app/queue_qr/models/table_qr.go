package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoTimeStamp struct ***REMOVED***
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
***REMOVED***

type TableQr struct ***REMOVED***
	Id primitive.ObjectID `bson:"_id"`
	MerchantId primitive.ObjectID `bson:"merchantId"`
	Name string `bson:"name"`
	FileId string `bson:"fileId"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
***REMOVED***

type TableQrDTO struct ***REMOVED***
	Id string `bson:"_id" json:"id"`
	MerchantId string `json:"merchantId"`
	Name string `json:"name"`
	FileId string `json:"fileId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
***REMOVED***

func (u *TableQr) MarshalBSON() ([]byte, error) ***REMOVED***
	if u.CreatedAt.IsZero() ***REMOVED***
			u.CreatedAt = time.Now()
	***REMOVED***
	u.UpdatedAt = time.Now()
	
	type my TableQr
	return bson.Marshal((*my)(u))
***REMOVED***

