package constants

type ErrorCodeType string
type ErrorLevelType string

const (
	FatalLevel ErrorLevelType = "1"
	ErrorLevel ErrorLevelType = "2"
	WarnLevel  ErrorLevelType = "3"
	InfoLevel  ErrorLevelType = "4"
	DebugLevel ErrorLevelType = "5"
	TraceLevel ErrorLevelType = "6"
)
