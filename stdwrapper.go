package log

import (
	stdlog "log"
)

type stdLogWriter struct {
	severity Severity
}

func (w stdLogWriter) Write(data []byte) (n int, err error) {
	switch w.severity {
	case TRACE:
		Trace(string(data[:len(data)-1]))
	case INFO:
		Info(string(data[:len(data)-1]))
	case WARNING:
		Warn(string(data[:len(data)-1]))
	case ERROR:
		Error(string(data[:len(data)-1]))
	case FATAL:
		Fatal(string(data[:len(data)-1]))
	}
	return len(data), nil
}

func NewStdLogger(severity Severity) *stdlog.Logger {
	return stdlog.New(stdLogWriter{
		severity: severity,
	}, "", 0)
}
