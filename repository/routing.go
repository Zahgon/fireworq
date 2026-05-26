package repository

// QueueNotFoundError is an error returned when a non-existing queue
// is specifie as the destination of a routing.
type QueueNotFoundError struct {
	QueueName string
}

func (qe *QueueNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }
