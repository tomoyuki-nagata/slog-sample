package middleware

import (
	"fmt"
	"log"
	"log/slog"
	"time"
	"todo-app/core/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func WithRequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := uuid.New().String()
		c.Set(config.REQUEST_ID, requestId)
		c.Header(config.REQUEST_ID_HEADER, requestId)
		c.Next()
	}
}

func WithExecutorId() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		userId := ""
		if token != "" {
			parser := &jwt.Parser{
				SkipClaimsValidation: true,
			}
			token, _, err := parser.ParseUnverified(token, jwt.MapClaims{})
			if err != nil {
				log.Fatal(err)
			}
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				userId = claims["sub"].(string)
			}
		}
		c.Set(config.USER_ID, userId)
		c.Next()
	}
}

func WithAccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		beforeAttributes := []slog.Attr{
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
			slog.String("route", c.FullPath()),
			slog.String("ip", c.ClientIP()),
			slog.String("user-agent", c.Request.UserAgent()),
		}

		slog.LogAttrs(c, slog.LevelInfo, "access log", beforeAttributes...)

		c.Next()
	}
}

func WithCustomGinLogger() gin.HandlerFunc {
	conf := gin.LoggerConfig{}
	conf.Formatter = customLogFormatter
	return gin.LoggerWithConfig(conf)
}

// ログフォーマッタ
var customLogFormatter = func(param gin.LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	if param.Latency > time.Minute {
		param.Latency = param.Latency.Truncate(time.Second)
	}

	return fmt.Sprintf("[GIN] %v| %s |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		param.Keys[config.REQUEST_ID],
		statusColor,
		param.StatusCode,
		resetColor,
		param.Latency,
		param.ClientIP,
		methodColor,
		param.Method,
		resetColor,
		param.Path,
		param.ErrorMessage,
	)
}
