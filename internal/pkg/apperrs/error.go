package apperrs

// Bad Request
var (
	InvalidFileSizeError             = NewError(BadRequest, "file size is invalid.")
	InvalidFileError                 = NewError(BadRequest, "file data or content type is invalid.")
	InvalidMetaInfoError             = NewError(BadRequest, "file meta info is invalid or empty.")
	InvalidEventDeviceIDError        = NewError(BadRequest, "event device id is invalid.")
	InvalidEventClientIDError        = NewError(BadRequest, "event client id is invalid.")
	InvalidClientCallingArchiveError = NewError(BadRequest, "this client calling archive is invalid.")
	EventTimeOutError                = NewError(BadRequest, "this event is timeout")
)

// NotFound
var (
	NotFoundArchiveError = NewError(NotFound, "The Archive is not exist.")
	NotFoundDeviceError  = NewError(NotFound, "The Device is not exist.")
)

// Unauthenticated
var UnauthenticatedError = NewError(Unauthenticated, "authentication error occurred. id or secret is invalid.")

// Forbidden
var ForbiddenError = NewError(Forbidden, "authorization error occurred. access denied.")

// Unexpected
var UnexpectedError = NewError(Unexpected, "unexpected error occurred.")

const (
	NotFound = iota
	BadRequest
	Unauthenticated
	Unexpected
	Forbidden
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
