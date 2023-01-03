package models

import (
	"time"

	enum "ant3/app/queue_qr/enum"

	"github.com/google/uuid"
)

type QueueTable struct ***REMOVED***
	Id uuid.UUID `json:"id"`
	TableName string `json:"tableName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Status enum.QUEUE_STATUS `json:"status"`
	QueueQR QueueQrDTO `json:"queueQr"`
***REMOVED***