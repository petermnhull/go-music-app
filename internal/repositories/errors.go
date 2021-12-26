package repositories

type ErrNoRecords struct {
	Message string
}

func (e *ErrNoRecords) Error() string {
	return e.Message
}
