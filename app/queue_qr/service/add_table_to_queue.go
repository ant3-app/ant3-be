package service

import (
	fb "ant3/app/config/firebase"
	"ant3/app/queue_qr/enum"
	"ant3/app/queue_qr/models"
	"ant3/app/queue_qr/repository"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func AddQRTableToQueue(queueQrId string) (error) ***REMOVED***
	queueQr, err := repository.GetOne(queueQrId)
	if(err != nil) ***REMOVED***
		fmt.Println(err, "[AddQRTableToQueue] error when get queue QR Data")
		return err
	***REMOVED***
	
	now := time.Now().UTC()
	
	var queueTable = models.QueueTable***REMOVED***
		Id: uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Status: enum.PENDING,
		QueueQR: *queueQr,
	***REMOVED***	
	
	fmt.Printf("response queueQR: %#v", *queueQr)
	
	ref := fb.Client.NewRef("queue_table")
	
	result, err := ref.Push(context.Background(), queueTable)
	if(err != nil) ***REMOVED***
		fmt.Println(err, "[AddQRTableToQueue] error when insert to firebase")
		return err
	***REMOVED***
	var _ = result
	return nil
***REMOVED***