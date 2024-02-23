package rotatelogs

import (
	"time"
)

const (
	optkeyClock          = "clock"
	optkeyHandler        = "handler"
	optkeyLinkName       = "link-name"
	optkeyMaxAge         = "max-age"
	optKeyMaxSize        = "max-size"
	optkeyRotationTime   = "rotation-time"
	optkeyRotationCount  = "rotation-count"
	optKeyRedirectStdErr = "redirect-stderr"
)

// WithClock creates a new Option that sets a clock
// that the RotateLogs object will use to determine
// the current time.
//
// By default rotatelogs.Local, which returns the
// current time in the local time zone, is used. If you
// would rather use UTC, use rotatelogs.UTC as the argument
// to this option, and pass it to the constructor.
func WithClock(c Clock) Option {
	return NewOpt(optkeyClock, c)
}

// WithLocation creates a new Option that sets up a
// "Clock" interface that the RotateLogs object will use
// to determine the current time.
//
// This optin works by always returning the in the given
// location.
func WithLocation(loc *time.Location) Option {
	return NewOpt(optkeyClock, clockFn(func() time.Time {
		return time.Now().In(loc)
	}))
}

// WithLinkName creates a new Option that sets the
// symbolic link name that gets linked to the current
// file name being used.
func WithLinkName(s string) Option {
	return NewOpt(optkeyLinkName, s)
}

// WithMaxAge creates a new Option that sets the
// max age of a log file before it gets purged from
// the file system.
func WithMaxAge(d time.Duration) Option {
	return NewOpt(optkeyMaxAge, d)
}

// WithMaxSize creates a new Option that sets the
// max size of a log file.
func WithMaxSize(size int64) Option {
	return NewOpt(optKeyMaxSize, size)
}

// WithRotationTime creates a new Option that sets the
// time between rotation.
func WithRotationTime(d time.Duration) Option {
	return NewOpt(optkeyRotationTime, d)
}

// WithRotationCount creates a new Option that sets the
// number of files should be kept before it gets
// purged from the file system.
func WithRotationCount(n uint) Option {
	return NewOpt(optkeyRotationCount, n)
}

// WithHandler creates a new Option that specifies the
// Handler object that gets invoked when an event occurs.
// Currently `FileRotated` event is supported
func WithHandler(h Handler) Option {
	return NewOpt(optkeyHandler, h)
}

// WithRedirectStdErr creates a new Option that specifies whether the current
// logger wants to redirect input of stderr (e.g. panic message) to its own log file.
// NOTE: in a project, there must be only one logger can set this option to true,
// so we recommend that global logger set this Option to true, and make sure other loggers
// do not set this Option (default value is false) or set this Option to false explicitly.
func WithRedirectStdErr(rd bool) Option {
	return NewOpt(optKeyRedirectStdErr, rd)
}
