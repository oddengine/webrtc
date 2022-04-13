package level

import (
	"strings"
)

const (
	// NONE completely disables logging of any events.
	NONE Level = 0x0000
	// TRACE is for logging very low-level library information (e.g. network traces).
	TRACE Level = 0x0001
	// DEBUG is for logging low-level library information (e.g. internal operations).
	DEBUG Level = 0x01FE
	// INFO is for logging normal library operation (e.g. state transitions, etc.).
	INFO Level = 0x0200
	// WARN is for logging abnormal, but non-fatal library operation.
	WARN Level = 0x0400
	// ERROR is for fatal errors which should be handled by user code.
	ERROR Level = 0x0800
)

// Debug-levels
const (
	DEBUG0 Level = 0x0002
	DEBUG1 Level = 0x0004
	DEBUG2 Level = 0x0008
	DEBUG3 Level = 0x0010
	DEBUG4 Level = 0x0020
	DEBUG5 Level = 0x0040
	DEBUG6 Level = 0x0080
	DEBUG7 Level = 0x0100
)

var (
	levels = map[string]Level{
		"none":   NONE,
		"trace":  TRACE,
		"debug":  DEBUG0,
		"debug0": DEBUG0,
		"debug1": DEBUG1,
		"debug2": DEBUG2,
		"debug3": DEBUG3,
		"debug4": DEBUG4,
		"debug5": DEBUG5,
		"debug6": DEBUG6,
		"debug7": DEBUG7,
		"info":   INFO,
		"warn":   WARN,
		"error":  ERROR,
	}
)

// Level represents the level at which the logger will emit log messages
type Level uint32

// Parse parses a readable string into Level.
func Parse(s string, sep string) Level {
	var (
		level = NONE
	)

	arr := strings.Split(s, sep)
	for _, v := range arr {
		n, ok := levels[v]
		if ok {
			level |= n
		}
	}
	return level
}
