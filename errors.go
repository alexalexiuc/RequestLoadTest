package main

import "os"

type LocalErrorCode int

const (
	CodeWarning LocalErrorCode = 2
	CodeError   LocalErrorCode = 1
)

type LocalError interface {
	Print()
	PrintFatal()
	SetAdditionalDetails(string)
}

type LocalErrorStruct struct {
	err        error
	msg        string
	code       LocalErrorCode
	addDetails string
	// hasError bool
}

func (e *LocalErrorStruct) Print() {
	PrintError("Error code: %d\n", e.code)
	PrintError("Error Msg : %s\n", e.msg)
	if e.addDetails != "" {
		PrintError("Details: %s\n", e.addDetails)
	}
	if e.err != nil {
		PrintError("Error Stack: %s\n", e.err)
	}
}

func (e *LocalErrorStruct) PrintFatal() {
	e.Print()
	os.Exit(1)
}

func (e *LocalErrorStruct) SetAdditionalDetails(d string) {
	e.addDetails = d
}

// WithError add error to struct and returns a copy of LocalErrorStruct
func (e *LocalErrorStruct) WithError(err error) LocalError {
	errCopy := *e
	errCopy.err = err
	return &errCopy
}

var ConfigFileReadErr = LocalErrorStruct{
	code: CodeError,
	msg:  "Unable to read config file",
}

var ConfigFileUnmarshalErr = LocalErrorStruct{
	code: CodeError,
	msg:  "Unable to unmarshal config file",
}

var MissingUserCredentialsErr = LocalErrorStruct{
	code: CodeWarning,
	msg:  "Missing user credentials",
}

var RequestMarshalErr = LocalErrorStruct{
	code: CodeError,
	msg:  "Unable to marshal request",
}

var UserLoginErr = LocalErrorStruct{
	code: CodeError,
	msg:  "Unable to login user",
}

var RespBodyUnmarshalErr = LocalErrorStruct{
	code: CodeError,
	msg:  "Unable to unmarshal response body",
}

var DataParseErr = LocalErrorStruct{
	code: CodeError,
	msg:  "Unable to parse data",
}

var TokenReceiveErr = LocalErrorStruct{
	code: CodeError,
	msg:  "Unable to get token from result. Format may be wrong",
}

var RequestExecErr = LocalErrorStruct{
	code: CodeError,
	msg:  "Unable to do request",
}

var RequestRespWarn = LocalErrorStruct{
	code: CodeWarning,
	msg:  "Request responded with non 200 status",
}
