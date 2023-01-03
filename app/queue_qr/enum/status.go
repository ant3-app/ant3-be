package enum

type QUEUE_STATUS int
const (
	PENDING QUEUE_STATUS = iota
	RESOLVED 
)

func (status QUEUE_STATUS) String() string ***REMOVED***
	return []string***REMOVED***"pending", "resolved"***REMOVED***[status]
***REMOVED***