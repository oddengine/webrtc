package log

import (
	"fmt"
	"io"
	"os"
	"sync"
	"unsafe"

	"gitlab.xthktech.cn/xthk-media/rawrtc/log/level"
)

var (
	std = NewDefaultLogger(os.Stdout, level.TRACE, "LOG")
)

// DefaultLogger encapsulates functionality for providing logging at user-defined levels.
type DefaultLogger struct {
	sync.RWMutex

	fd    unsafe.Pointer
	level level.Level
	scope string
	trace io.Writer
	debug io.Writer
	info  io.Writer
	warn  io.Writer
	error io.Writer
}

// Init this class.
func (me *DefaultLogger) Init(n level.Level, scope string) *DefaultLogger {
	me.level = n
	me.scope = scope
	return me
}

// WithTrace is a chainable configuration function which sets the trace-level logger.
func (me *DefaultLogger) WithTrace(out io.Writer) *DefaultLogger {
	me.level |= level.TRACE
	me.trace = out
	return me
}

// WithDebug is a chainable configuration function which sets the debug-level logger.
func (me *DefaultLogger) WithDebug(out io.Writer) *DefaultLogger {
	me.debug = out
	return me
}

// WithInfo is a chainable configuration function which sets the info-level logger.
func (me *DefaultLogger) WithInfo(out io.Writer) *DefaultLogger {
	me.info = out
	return me
}

// WithWarn is a chainable configuration function which sets the warn-level logger.
func (me *DefaultLogger) WithWarn(out io.Writer) *DefaultLogger {
	me.warn = out
	return me
}

// WithError is a chainable configuration function which sets the error-level logger.
func (me *DefaultLogger) WithError(out io.Writer) *DefaultLogger {
	me.error = out
	return me
}

// Trace emits the preformatted message if the logger is at or below trace-level.
func (me *DefaultLogger) Trace(s string) {
	if me.trace != nil && (me.level&level.TRACE) != 0 {
		me.trace.Write([]byte(s + "\n"))
	}
}

// Tracef formats and emits a message if the logger is at or below trace-level.
func (me *DefaultLogger) Tracef(format string, args ...interface{}) {
	if me.trace != nil && (me.level&level.TRACE) != 0 {
		s := fmt.Sprintf(format, args...)
		me.trace.Write([]byte(s + "\n"))
	}
}

// Debug emits the preformatted message if the logger is at or below debug-level.
func (me *DefaultLogger) Debug(n uint32, s string) {
	if me.debug != nil && (me.level&level.DEBUG) <= level.DEBUG0<<n {
		me.debug.Write([]byte(s))
	}
}

// Debugf formats and emits a message if the logger is at or below debug-level.
func (me *DefaultLogger) Debugf(n uint32, format string, args ...interface{}) {
	if me.debug != nil && (me.level&level.DEBUG) <= level.DEBUG0<<n {
		s := fmt.Sprintf(format, args...)
		me.debug.Write([]byte(s))
	}
}

// Info emits the preformatted message if the logger is at or below info-level.
func (me *DefaultLogger) Info(s string) {
	if me.info != nil && (me.level&level.INFO) != 0 {
		me.info.Write([]byte(s))
	}
}

// Infof formats and emits a message if the logger is at or below info-level.
func (me *DefaultLogger) Infof(format string, args ...interface{}) {
	if me.info != nil && (me.level&level.INFO) != 0 {
		s := fmt.Sprintf(format, args...)
		me.info.Write([]byte(s))
	}
}

// Warn emits the preformatted message if the logger is at or below warn-level.
func (me *DefaultLogger) Warn(s string) {
	if me.warn != nil && (me.level&level.WARN) != 0 {
		me.warn.Write([]byte(s))
	}
}

// Warnf formats and emits a message if the logger is at or below warn-level.
func (me *DefaultLogger) Warnf(format string, args ...interface{}) {
	if me.warn != nil && (me.level&level.WARN) != 0 {
		s := fmt.Sprintf(format, args...)
		me.warn.Write([]byte(s))
	}
}

// Error emits the preformatted message if the logger is at or below error-level.
func (me *DefaultLogger) Error(s string) {
	if me.error != nil && (me.level&level.ERROR) != 0 {
		me.error.Write([]byte(s))
	}
}

// Errorf formats and emits a message if the logger is at or below error-level.
func (me *DefaultLogger) Errorf(format string, args ...interface{}) {
	if me.error != nil && (me.level&level.ERROR) != 0 {
		s := fmt.Sprintf(format, args...)
		me.error.Write([]byte(s))
	}
}

// NewDefaultLogger returns a configured ILogger.
func NewDefaultLogger(out io.Writer, n level.Level, scope string) *DefaultLogger {
	return new(DefaultLogger).Init(n, scope).
		WithTrace(os.Stdout).
		WithDebug(out).
		WithInfo(out).
		WithWarn(out).
		WithError(out)
}
