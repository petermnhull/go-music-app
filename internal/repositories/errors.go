package repositories

type ErrNoRecords struct {
	Message string
}

func (e *ErrNoRecords) Error() string {
	return e.Message
}

func NewErrNoRecords(msg string) *ErrNoRecords {
	return &ErrNoRecords{Message: msg}
}
