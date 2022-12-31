package models

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoTimeStamp struct ***REMOVED***
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
***REMOVED***

type QueueQR struct ***REMOVED***
	// Id primitive.ObjectID `bson:"_id"`
	MerchantId primitive.ObjectID `bson:"merchantId"`
	Number int32 `bson:"number"`
	FileId string `bson:"fileId"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
***REMOVED***

func (u *QueueQR) MarshalBSON() ([]byte, error) ***REMOVED***
	fmt.Println("wewewewewewe aneeh")
	if u.CreatedAt.IsZero() ***REMOVED***
			u.CreatedAt = time.Now()
	***REMOVED***
	u.UpdatedAt = time.Now()
	
	type my QueueQR
	return bson.Marshal((*my)(u))
***REMOVED***