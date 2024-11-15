package sl

import "log/slog"

func Error(err error) slog.Attr {
	return slog.Attr{
		Key:   "Error",
		Value: slog.StringValue(err.Error()),
	}
}
