package i18n

import "context"

type Logger interface {
	Warningf(c context.Context, format string, args ...interface{})
}

var log Logger

func warningf(c context.Context, format string, args ...interface{}) {
	if log != nil {
		log.Warningf(c, format, args...)
	}
}
