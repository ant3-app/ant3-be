package service

import (
	fb "ant3/app/config/firebase"
	"ant3/app/queue_qr/models"
	"ant3/app/queue_qr/repository"
	"context"
	"errors"
	"fmt"
	"time"
)

func AddTableQrToQueue(queueQrId string) (*string, error) ***REMOVED***
	queueQr, err := repository.GetOne(queueQrId)
	if(err != nil) ***REMOVED***
		fmt.Println(err.Error(), "[AddQRTableToQueue] error when get queue QR Data")
		return nil, err
	***REMOVED***
	
	now := time.Now().UTC()
	
	var queueTable = models.QueueTable***REMOVED***
		CreatedAt: now,
		UpdatedAt: now,
	***REMOVED***	
	
	fmt.Printf("response queueTable: %#v\n", queueTable)
	// Noted: try to validate using firebase rules
	var existQueueQr models.QueueTable
	
	ref := fb.Client.NewRef("queue_table")
	myErr := ref.
		Child(queueQr.Id).
		Get(context.Background(), &existQueueQr)
		
	if(myErr != nil) ***REMOVED***
		fmt.Println("error get queueId: " + queueQr.Id, myErr.Error())
	***REMOVED***
	
	fmt.Printf("response existQueueQr: %#v\n", existQueueQr)
	
	if (existQueueQr != models.QueueTable***REMOVED******REMOVED***) ***REMOVED***
		return nil, errors.New("The " + queueQr.Id + " table is on the queue")
	***REMOVED***
	
	
	fbErr := ref.Child(queueQr.Id).
		Set(context.Background(), &queueTable)
	if(fbErr != nil) ***REMOVED***
		fmt.Println(err, "[AddQRTableToQueue] error when insert to firebase")
		return nil, err
	***REMOVED***
	return &queueQr.Id, nil
***REMOVED***