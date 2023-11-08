package log

import "context"

func Debug(ctx context.Context, msg string, fields ...interface{}) {
	FromContext(ctx).Debug(msg, fields...)
}

func Info(ctx context.Context, msg string, fields ...interface{}) {
	FromContext(ctx).Info(msg, fields...)
}

func Error(ctx context.Context, err error, msg string, fields ...interface{}) {
	FromContext(ctx).Error(err, msg, fields...)
}

func Derive(ctx context.Context, options ...Option) context.Context {
	return WithLogger(ctx, FromContext(ctx).Derive(options...))
}

type ctxKey int

const keyLogger ctxKey = 0x123ABCF

var appLog = DefaultLogger

func FromContext(ctx context.Context) Logger {
	if l, ok := ctx.Value(keyLogger).(Logger); ok && l != nil {
		return l
	}
	return appLog
}

func WithLogger(ctx context.Context, l Logger) context.Context {
	return context.WithValue(ctx, keyLogger, l)
}

func SetLevel(level Level) {
	appLog = appLog.Derive(WithLevel(level))
}
