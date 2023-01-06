package service

import (
	fb "ant3/app/config/firebase"
	"ant3/app/queue_qr/repository"
	"context"
	"errors"
	"fmt"
	"time"
)

func AddTableQrToQueue(tableQrId string) (*string, error) ***REMOVED***
	tableQr, err := repository.GetOne(tableQrId)
	if(err != nil) ***REMOVED***
		fmt.Println(err.Error(), "[AddQRTableToQueue] error when get queue QR Data")
		return nil, err
	***REMOVED***
	
	now := time.Now().UTC()
	
	var queueTable =  map[string]interface***REMOVED******REMOVED*** ***REMOVED***
		"createdAt": now,
		"updatedAt": now,
		"name": tableQr.Name,
	***REMOVED***
	
	fmt.Printf("response queueTable: %#v\n", queueTable)
	// Noted: try to validate using firebase rules
	
	ref := fb.Client.Collection("merchant")
	refQueue := ref.Doc(tableQr.MerchantId).Collection("queue")
	doc, myErr := refQueue.Doc(tableQr.Id).Get(context.Background())
	existingQueueQr := doc.Data()
		
		
	if(myErr != nil) ***REMOVED***
		fmt.Println("error get queueId: " + tableQr.Id, myErr.Error())
	***REMOVED***
	
	fmt.Printf("response existQueueQr: %#v\n", existingQueueQr)
	
	if (existingQueueQr != nil) ***REMOVED***
		return nil, errors.New("The " + tableQr.Id + " table is on the queue")
	***REMOVED***
	
	_, fbErr := refQueue.Doc(tableQr.Id).
		Set(context.Background(), queueTable)
		
	if(fbErr != nil) ***REMOVED***
		fmt.Println(err, "[AddQRTableToQueue] error when insert to firebase")
		return nil, err
	***REMOVED***
	return &tableQr.Id, nil
***REMOVED***


func RemoveTableFromQueue(tableQrId string) (*string, error) ***REMOVED***
	tableQr, err := repository.GetOne(tableQrId)
	if(err != nil) ***REMOVED***
		fmt.Println(err.Error(), "[AddQRTableToQueue] error when get queue QR Data")
		return nil, err
	***REMOVED***
	
	ref := fb.Client.Collection("merchant" )
	refQueue := ref.Doc(tableQr.MerchantId).Collection("queue")
	_, fbErr := refQueue.Doc(tableQr.Id).
		Delete(context.Background())
		
	if(fbErr != nil) ***REMOVED***
		fmt.Println(err, "[AddQRTableToQueue] error when delete data from firebase")
		return nil, err
	***REMOVED***
	
	return &tableQr.Id, nil
***REMOVED***