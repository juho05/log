package log

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
)

type color string

const (
	colorTrace   color = "\033[2m\033[37m"
	colorInfo    color = "\033[36m"
	colorWarning color = "\033[33m"
	colorError   color = "\033[1m\033[31m"
	colorFatal   color = "\033[4m\033[1m\033[35m"
)

type Severity int

const (
	NONE Severity = iota
	FATAL
	ERROR
	WARNING
	INFO
	TRACE
)

var (
	out          = colorable.NewColorableStderr()
	colorEnabled = true
	minSeverity  = INFO
)

func log(severity Severity, args ...interface{}) {
	if severity <= minSeverity {
		fmt.Fprint(out, logPrefix(severity), fmt.Sprintln(args...))
	}
}

func logf(severity Severity, format string, args ...interface{}) {
	if severity <= minSeverity {
		fmt.Fprint(out, logPrefix(severity), fmt.Sprintf(format+"\n", args...))
	}
}

func Trace(args ...interface{}) {
	setColor(colorTrace)
	log(TRACE, args...)
	resetColor()
}

func Tracef(format string, args ...interface{}) {
	setColor(colorTrace)
	logf(TRACE, format, args...)
	resetColor()
}

func Info(args ...interface{}) {
	setColor(colorInfo)
	log(INFO, args...)
	resetColor()
}

func Infof(format string, args ...interface{}) {
	setColor(colorInfo)
	logf(INFO, format, args...)
	resetColor()
}

func Warn(args ...interface{}) {
	setColor(colorWarning)
	log(WARNING, args...)
	resetColor()
}

func Warnf(format string, args ...interface{}) {
	setColor(colorWarning)
	logf(WARNING, format, args...)
	resetColor()
}

func Error(args ...interface{}) {
	setColor(colorError)
	log(ERROR, args...)
	resetColor()
}

func Errorf(format string, args ...interface{}) {
	setColor(colorError)
	logf(ERROR, format, args...)
	resetColor()
}

func Fatal(args ...interface{}) {
	setColor(colorFatal)
	log(FATAL, args...)
	resetColor()
	os.Exit(1)
}

func Fatalf(format string, args ...interface{}) {
	setColor(colorFatal)
	logf(FATAL, format, args...)
	resetColor()
	os.Exit(1)
}

func SetSeverity(severity Severity) {
	minSeverity = severity
}

func SetOutput(file *os.File) {
	if isatty.IsTerminal(file.Fd()) {
		out = colorable.NewColorable(file)
	} else {
		out = colorable.NewNonColorable(file)
	}
}

func Output() io.Writer {
	return out
}

func logPrefix(severity Severity) string {
	return fmt.Sprintf("%s %s", time.Now().UTC().Format("2006-01-02 15:04:05.000"), severityTag(severity))
}

func setColor(color color) {
	fmt.Fprint(out, color)
}

func resetColor() {
	fmt.Fprint(out, "\033[0m")
}

func severityTag(severity Severity) string {
	switch severity {
	case TRACE:
		return "[TRACE]: "
	case INFO:
		return "[INFO] : "
	case WARNING:
		return "[WARN] : "
	case ERROR:
		return "[ERROR]: "
	case FATAL:
		return "[FATAL]: "
	default:
		return "[???]  : "
	}
}
