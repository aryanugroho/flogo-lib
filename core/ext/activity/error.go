package activity


//todo add error code
type Error struct {
	errorStr string
}

func NewError(errorText string) error {
	return &Error{errorStr:errorText}
}

func (e *Error) Error() string {
	return e.errorStr
}