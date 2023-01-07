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

const (
	YYYY_MM_DD = "2006-01-02"
	MERCHANT = "merchant"
	QUEUE = "queue"
	QUEUE_COUNT = "queue_count"
)

func validateGenerateNewQueue(
	trx *firestore.Transaction, 
	docRef *firestore.DocumentRef, 
	tableQrId string,
) (bool) {
	doc, err := trx.Get(docRef)
	if(err != nil) {
		fmt.Println("error get queueId: " + tableQrId, err.Error())
	}
	
	// Note: try to validate using firebase rules
	existingQueueQr := doc.Data()
	if (existingQueueQr != nil) {
		return false
		
	}
	fmt.Printf("response existQueueQr: %#v\n", existingQueueQr)
	return true
}


func getQueueCount(
	trx *firestore.Transaction,
	colMerchant *firestore.CollectionRef,
	merchantId string,
	date time.Time,
) (*int64, error) {
	refQueueCount := colMerchant.Doc(merchantId).Collection(QUEUE_COUNT)
	docRefQueueCount := refQueueCount.Doc(date.Format(YYYY_MM_DD))
	
	queueCountRef, getQueueCountErr := trx.Get(docRefQueueCount)
	var incrementedNum int64 = 1
	
	if (getQueueCountErr == nil) {
		incrementedNum = queueCountRef.Data()["number"].(int64) + 1
	}
	
	err := trx.Set(docRefQueueCount, map[string]interface{} {
		"number": incrementedNum,
	})
	
	if (err != nil) {
		fmt.Printf("[ERROR] Failed to set: %#v \n", docRefQueueCount)
		return nil, err
	}
	return &incrementedNum, nil
}

func AddTableQrToQueue(tableQrId string) (*string, error) {
	now := time.Now().UTC()
	tableQr, err := repository.GetOne(tableQrId)
	if(err != nil) {
		fmt.Println(err.Error(), "[AddQRTableToQueue] error when get queue QR Data")
		return nil, err
	}
	colMerchant := fb.Client.Collection(MERCHANT)
		
	addTableQrToQueueTransaction := func(ctx context.Context, trx *firestore.Transaction) error {
		
		refQueue := colMerchant.Doc(tableQr.MerchantId).Collection(QUEUE)
		docRefQueue := refQueue.Doc(tableQr.Id)
		
		isAbleToGenerateNewQueue := validateGenerateNewQueue(trx, docRefQueue, tableQr.Id)
		if (!isAbleToGenerateNewQueue) {
			return errors.New("The " + tableQrId + " table is on the queue")
		}
		
		queueCountNumber, err := getQueueCount(trx, colMerchant, tableQr.MerchantId, now)
		if (err != nil) {
			return err
		}
		
		var queueTable =  map[string]interface{} {
			"createdAt": now,
			"updatedAt": now,
			"name": tableQr.Name,
			"number": &queueCountNumber,
		}
		fmt.Printf("response queueTable: %#v\n", queueTable)
		
		fbErr := trx.Set(docRefQueue, queueTable)
		if(fbErr != nil) {
			fmt.Printf("[ERROR] Failed to set: %#v \n", docRefQueue)
			return fbErr
		}
			
		return nil
	}
	
	errTrx := fb.Client.RunTransaction(context.Background(), addTableQrToQueueTransaction)
		
	if(errTrx != nil) {
		fmt.Println(errTrx, "[AddQRTableToQueue] error when insert to firebase")
		return nil, errTrx
	}
	return &tableQr.Id, nil
}


func RemoveTableFromQueue(tableQrId string) (*string, error) {
	tableQr, err := repository.GetOne(tableQrId)
	if(err != nil) {
		fmt.Println(err.Error(), "[AddQRTableToQueue] error when get queue QR Data")
		return nil, err
	}
	
	ref := fb.Client.Collection(MERCHANT)
	refQueue := ref.Doc(tableQr.MerchantId).Collection(QUEUE)
	_, fbErr := refQueue.Doc(tableQr.Id).
		Delete(context.Background())
		
	if(fbErr != nil) {
		fmt.Println(err, "[AddQRTableToQueue] error when delete data from firebase")
		return nil, err
	}
	
	return &tableQr.Id, nil
}