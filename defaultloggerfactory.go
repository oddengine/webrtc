package rawrtc

/*
#cgo CFLAGS : -g -I../../include
#cgo linux LDFLAGS: -L../../lib -ldl
#cgo windows LDFLAGS: -L../../../lib

#include "api.h"
*/
import "C"
import (
	"io"
	"strings"
	"unsafe"

	"gitlab.xthktech.cn/xthk-media/rawrtc/level"
)

// DefaultLoggerFactory creates new DefaultLogger.
type DefaultLoggerFactory struct {
	fd    unsafe.Pointer
	out   io.Writer
	level level.Level
}

// Init this class.
func (me *DefaultLoggerFactory) Init(out io.Writer, n level.Level) *DefaultLoggerFactory {
	me.out = out
	me.level = n
	return me
}

// NewLogger returns a configured ILogger for the given scope.
func (me *DefaultLoggerFactory) NewLogger(scope string) ILogger {
	l := NewDefaultLogger(me.out, me.level, strings.ToUpper(scope))
	l.fd = C.CreateDefaultLogger(unsafe.Pointer(l), me.fd, C.CString(scope)) // Ignore memory overflow of "scope".
	return l
}

// NewDefaultLoggerFactory creates a new DefaultLoggerFactory.
func NewDefaultLoggerFactory(constraints *DefaultWriterConstraints) *DefaultLoggerFactory {
	w := new(DefaultWriter).Init(constraints)
	n := level.Parse(constraints.Level, "|")
	f := new(DefaultLoggerFactory).Init(w, n)
	f.fd = C.CreateDefaultLoggerFactory(unsafe.Pointer(f), w.fd, C.int(n))
	return f
}
