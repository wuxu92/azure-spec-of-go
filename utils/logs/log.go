package logs

import "fmt"

type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
}

type fmtLogger struct{}

func (f fmtLogger) Debug(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	fmt.Println("")
}

func (f fmtLogger) Info(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	fmt.Println("")
}

func (f fmtLogger) Warn(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	fmt.Println("")
}

func (f fmtLogger) Error(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	fmt.Println("")
}

var _ Logger = (*fmtLogger)(nil)

type sysLogger struct{}

func (f sysLogger) Debug(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	fmt.Println("")
}

func (f sysLogger) Info(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	fmt.Println("")
}

func (f sysLogger) Warn(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	fmt.Println("")
}

func (f sysLogger) Error(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	fmt.Println("")
}

var defaultLogger Logger

func InitLogger(logger Logger) {
	if logger != nil {
		defaultLogger = logger
	} else {
		defaultLogger = fmtLogger{}
	}
}

func init() {
	InitLogger(nil)
}

func Debug(format string, args ...interface{}) {
	defaultLogger.Debug(format, args...)
}

func Info(format string, args ...interface{}) {
	defaultLogger.Info(format, args...)
}

func Warn(format string, args ...interface{}) {
	defaultLogger.Warn(format, args...)
}

func Error(format string, args ...interface{}) {
	defaultLogger.Error(format, args...)
}
