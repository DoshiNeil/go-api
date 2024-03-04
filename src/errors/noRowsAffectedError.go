package errors

type NoRowsAffectedError struct{}

func (e *NoRowsAffectedError) Error() string {
	return "No rows affected"
}
