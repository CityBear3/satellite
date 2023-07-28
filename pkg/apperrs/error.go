package apperrs

const (
	NotFound = iota
	BadRequest
)

const (
	/*
		BadRequestError Message
	*/
	InvalidMetaInfoMsg = "File meta info is invalid or empty."
	InvalidDataMsg     = "File data is invalid or empty."
	InvalidFileExtMsg  = "File extension must be jpeg or png."
	InvalidFileSizeMSG = "File size limitation is %s."
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
