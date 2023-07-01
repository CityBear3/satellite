package apperrs

const (
	NotFound = iota
	BadRequest
)

type Error struct {
	Code int
	Msg  string
}

func NewError(code int, msg string) *Error {
	return &Error{Code: code, Msg: msg}
}

func (a *Error) Error() string {
	return a.Msg
}
