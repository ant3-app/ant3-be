package models

import (
	"time"

	enum "ant3/app/queue_qr/enum"

	"github.com/google/uuid"
)

type QueueTable struct ***REMOVED***
	Id uuid.UUID
	TableName string
	CreatedAt time.Time
	UpdatedAt time.Time
	Status enum.QUEUE_STATUS
	QueueQR QueueQrDTO
***REMOVED***