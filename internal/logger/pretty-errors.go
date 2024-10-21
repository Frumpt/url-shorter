package logger

import "log/slog"

func ErrAttr(err any) slog.Attr {
	return slog.Any("error", err)
}
