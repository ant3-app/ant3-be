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

func AddTableQrToQueue(tableQrId string) (*string, error) ***REMOVED***
	tableQr, err := repository.GetOne(tableQrId)
	if(err != nil) ***REMOVED***
		fmt.Println(err.Error(), "[AddQRTableToQueue] error when get queue QR Data")
		return nil, err
	***REMOVED***
	
	now := time.Now().UTC()
	
	var queueTable = models.TableQueue***REMOVED***
		CreatedAt: now,
		UpdatedAt: now,
	***REMOVED***	
	
	fmt.Printf("response queueTable: %#v\n", queueTable)
	// Noted: try to validate using firebase rules
	var existQueueQr models.TableQueue
	
	ref := fb.Client.NewRef("queue_table")
	myErr := ref.
		Child(tableQr.Id).
		Get(context.Background(), &existQueueQr)
		
	if(myErr != nil) ***REMOVED***
		fmt.Println("error get queueId: " + tableQr.Id, myErr.Error())
	***REMOVED***
	
	fmt.Printf("response existQueueQr: %#v\n", existQueueQr)
	
	if (existQueueQr != models.TableQueue***REMOVED******REMOVED***) ***REMOVED***
		return nil, errors.New("The " + tableQr.Id + " table is on the queue")
	***REMOVED***
	
	
	fbErr := ref.Child(tableQr.Id).
		Set(context.Background(), &queueTable)
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
	
	ref := fb.Client.NewRef("queue_table")
	fbErr := ref.Child(tableQr.Id).
		Delete(context.Background())
		
	if(fbErr != nil) ***REMOVED***
		fmt.Println(err, "[AddQRTableToQueue] error when delete data from firebase")
		return nil, err
	***REMOVED***
	
	return &tableQr.Id, nil
***REMOVED***