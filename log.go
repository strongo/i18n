package i18n

import "context"

type Logger interface {
	Debugf(c context.Context, format string, args ...any)
	Errorf(c context.Context, format string, args ...any)
	Warningf(c context.Context, format string, args ...any)
}

var log Logger

func warningf(c context.Context, format string, args ...any) {
	if log != nil {
		log.Warningf(c, format, args...)
	}
}

func errorf(c context.Context, format string, args ...any) {
	if log != nil {
		log.Debugf(c, format, args...)
	}
}

func debugf(c context.Context, format string, args ...any) {
	if log != nil {
		log.Debugf(c, format, args...)
	}
}
