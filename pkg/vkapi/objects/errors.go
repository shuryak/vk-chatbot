package objects

import "fmt"

type ErrorType int

// https://dev.vk.com/reference/errors
const (
	ErrorNoType                          ErrorType = 0
	ErrorUnknown                         ErrorType = 1
	ErrorDisabled                        ErrorType = 2
	ErrorMethod                          ErrorType = 3
	ErrorSignature                       ErrorType = 4
	ErrorAuthorization                   ErrorType = 5
	ErrorTooManyRequests                 ErrorType = 6
	ErrorPermission                      ErrorType = 7
	ErrorInvalidRequest                  ErrorType = 8
	ErrorFlood                           ErrorType = 9
	ErrorInternal                        ErrorType = 10
	ErrorEnabledInTest                   ErrorType = 11
	ErrorCaptcha                         ErrorType = 14
	ErrorAccessDenied                    ErrorType = 15
	ErrorHTTPS                           ErrorType = 16
	ErrorValidationRequired              ErrorType = 17
	ErrorUserDeleted                     ErrorType = 18
	ErrorNonStandalonePermission         ErrorType = 20
	ErrorOnlyStandaloneAndOpenAPIAllowed ErrorType = 21
	ErrorMethodDisabled                  ErrorType = 23
	ErrorNeedConfirmation                ErrorType = 24
	ErrorGroupKeyInvalid                 ErrorType = 27
	ErrorAppKeyInvalid                   ErrorType = 28
	ErrorRateLimit                       ErrorType = 29
	// TODO: ...
)

type Error struct {
	Code    ErrorType `json:"error_code"`
	Message string    `json:"error_msg"`
	Text    string    `json:"error_text"`
	// TODO: complete errors
}

func (e Error) Error() string {
	return "api: " + e.Message
}

type ExecuteError struct {
	Method string `json:"method"`
	Code   int    `json:"error_code"`
	Msg    string `json:"error_msg"`
}

type ExecuteErrors []ExecuteError

func (e ExecuteErrors) Error() string {
	return fmt.Sprintf("api: execute errors (%d)", len(e))
}

type UploadError struct {
	Err      string `json:"error"`
	Code     int    `json:"error_code"`
	Descr    string `json:"error_descr"`
	IsLogged bool   `json:"error_is_logged"`
}

func (e UploadError) Error() string {
	if e.Err != "" {
		return "api: " + e.Err
	}

	return fmt.Sprintf("api: upload code %d", e.Code)
}
