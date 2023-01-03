package enum

type QUEUE_STATUS string
const (
	PENDING QUEUE_STATUS = "pending"
	RESOLVED QUEUE_STATUS = "resolved"
)