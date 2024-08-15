package logger

import (
	"context"
	"log/slog"
	"os"
	"todo-app/core/config"
)

// slogのデフォルトのハンドラーを設定する。
func Initialize() {
	handler := &ApiHandler{
		parent: slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level:     slog.LevelDebug,
				AddSource: true,
			},
		),
	}
	slog.SetDefault(slog.New(handler))
}

// API用のカスタムハンドラー
type ApiHandler struct {
	parent slog.Handler
}

// カスタムハンドラの処理。
// ここではrequestIdとuserIdをログに出力するようにしている。
func (h *ApiHandler) Handle(c context.Context, record slog.Record) error {
	requestId := c.Value(config.REQUEST_ID)
	if id, ok := requestId.(string); ok {
		record.Add(slog.String(config.REQUEST_ID, id))
	}
	userId := c.Value(config.USER_ID)
	if id, ok := userId.(string); ok {
		record.Add(slog.String(config.USER_ID, id))
	}

	return h.parent.Handle(c, record)
}

func (h *ApiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.parent.Enabled(ctx, level)
}

func (h *ApiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ApiHandler{h.parent.WithAttrs(attrs)}
}

func (h *ApiHandler) WithGroup(name string) slog.Handler {
	return &ApiHandler{h.parent.WithGroup(name)}
}
