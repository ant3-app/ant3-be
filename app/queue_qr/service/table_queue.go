package service

import (
	fb "ant3/app/config/firebase"
	"ant3/app/queue_qr/repository"
	"context"
	"errors"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
)

func getLastDocNum(trx *firestore.Transaction, colRef *firestore.CollectionRef) (*int64, error) ***REMOVED***
	query := colRef.OrderBy("createdAt", firestore.Desc).Limit(1)
	lastDoc, err := trx.Documents(query).Next()
	if(err != nil) ***REMOVED***
		return nil, err
	***REMOVED***
	
	numberLastDoc := lastDoc.Data()["number"].(int64)
	
	return &numberLastDoc, nil
***REMOVED***

func validateGenerateNewQueue(
	trx *firestore.Transaction, 
	docRef *firestore.DocumentRef, 
	tableQrId string,
) (bool) ***REMOVED***
	doc, err := trx.Get(docRef)
	if(err != nil) ***REMOVED***
		fmt.Println("error get queueId: " + tableQrId, err.Error())
	***REMOVED***
	
	// Note: try to validate using firebase rules
	existingQueueQr := doc.Data()
	if (existingQueueQr != nil) ***REMOVED***
		return false
		
	***REMOVED***
	fmt.Printf("response existQueueQr: %#v\n", existingQueueQr)
	return true
***REMOVED***

func AddTableQrToQueue(tableQrId string) (*string, error) ***REMOVED***
	tableQr, err := repository.GetOne(tableQrId)
	if(err != nil) ***REMOVED***
		fmt.Println(err.Error(), "[AddQRTableToQueue] error when get queue QR Data")
		return nil, err
	***REMOVED***
		
	addTableQrToQueueTransaction := func(ctx context.Context, trx *firestore.Transaction) error ***REMOVED***
		colMerchant := fb.Client.Collection("merchant")
		refQueue := colMerchant.Doc(tableQr.MerchantId).Collection("queue")
		docRefQueue := refQueue.Doc(tableQr.Id)
		
		lastDocNumber, err := getLastDocNum(trx, refQueue)
		if (err != nil) ***REMOVED***
			return err
		***REMOVED***
		
		now := time.Now().UTC()
	
		var queueTable =  map[string]interface***REMOVED******REMOVED*** ***REMOVED***
			"createdAt": now,
			"updatedAt": now,
			"name": tableQr.Name,
			"number": *lastDocNumber + 1,
		***REMOVED***
		fmt.Printf("response queueTable: %#v\n", queueTable)
		
		isAbleToGenerateNewQueue := validateGenerateNewQueue(trx, docRefQueue, tableQr.Id)
		if (!isAbleToGenerateNewQueue) ***REMOVED***
			return errors.New("The " + tableQrId + " table is on the queue")
		***REMOVED***
		
		fbErr := trx.Set(docRefQueue, queueTable)
		if(fbErr != nil) ***REMOVED***
			return fbErr
		***REMOVED***
			
		return nil
	***REMOVED***
	
	errTrx := fb.Client.RunTransaction(context.Background(), addTableQrToQueueTransaction)
		
	if(errTrx != nil) ***REMOVED***
		fmt.Println(errTrx, "[AddQRTableToQueue] error when insert to firebase")
		return nil, errTrx
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