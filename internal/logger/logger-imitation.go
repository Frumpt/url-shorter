package logger

import (
	"context"
	"log/slog"
)

func NewImitationLogger() *slog.Logger {
	return slog.New(loggerImitation{})
}

type loggerImitation struct{}

func (l loggerImitation) Enabled(_ context.Context, _ slog.Level) bool {
	return false
}

func (l loggerImitation) Handle(_ context.Context, _ slog.Record) error {
	return nil
}

func (l loggerImitation) WithAttrs(_ []slog.Attr) slog.Handler {
	return l
}

func (l loggerImitation) WithGroup(_ string) slog.Handler {
	return l
}
