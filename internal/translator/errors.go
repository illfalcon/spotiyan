package translator

type BadRequestError struct {
	err error
	msg string
}

func NewBadRequestError(err error, msg string) *BadRequestError {
	return &BadRequestError{err: err, msg: msg}
}

func (e *BadRequestError) Error() string {
	return e.msg
}

func (e *BadRequestError) Unwrap() error {
	return e.err
}
