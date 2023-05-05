package i18n

import "context"

type Logger interface {
	Debugf(c context.Context, format string, args ...interface{})
	Errorf(c context.Context, format string, args ...interface{})
	Warningf(c context.Context, format string, args ...interface{})
}

var log Logger

func warningf(c context.Context, format string, args ...interface{}) {
	if log != nil {
		log.Warningf(c, format, args...)
	}
}

func errorf(c context.Context, format string, args ...interface{}) {
	if log != nil {
		log.Debugf(c, format, args...)
	}
}

func debugf(c context.Context, format string, args ...interface{}) {
	if log != nil {
		log.Debugf(c, format, args...)
	}
}
