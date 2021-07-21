package mwLogger

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lifegit/go-gulu/v2/pkg/logging"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"time"
)

var (
	green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow  = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset   = string([]byte{27, 91, 48, 109})
)

type LogFormatterParams struct {
	TimeStamp    time.Time              `json:"time"`
	StatusCode   int                    `json:"code"`
	Latency      time.Duration          `json:"latency"`
	ClientIP     string                 `json:"ip"`
	Method       string                 `json:"method"`
	Path         string                 `json:"path"`
	ErrorMessage string                 `json:"errorMessage"`
	BodySize     int                    `json:"bodySize"`
	Keys         map[string]interface{} `json:"keys"`
}

var stdoutLogFormatter = func(param LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	statusColor = param.StatusCodeColor()
	methodColor = param.MethodColor()
	resetColor = param.ResetColor()

	if param.Latency > time.Minute {
		// Truncate in a golang < 1.8 safe way
		param.Latency = param.Latency - param.Latency%time.Second
	}
	return fmt.Sprintf("[GIN] %v |%s %3d %s| %13v | %15s |%s %-7s %s %s\n%s",
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, param.StatusCode, resetColor,
		param.Latency,
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
		param.ErrorMessage,
	)
}

// StatusCodeColor is the ANSI color for appropriately logging http status code to a terminal.
func (p *LogFormatterParams) StatusCodeColor() string {
	code := p.StatusCode

	switch {
	case code >= http.StatusOK && code < http.StatusMultipleChoices:
		return green
	case code >= http.StatusMultipleChoices && code < http.StatusBadRequest:
		return white
	case code >= http.StatusBadRequest && code < http.StatusInternalServerError:
		return yellow
	default:
		return red
	}
}

// MethodColor is the ANSI color for appropriately logging http method to a terminal.
func (p *LogFormatterParams) MethodColor() string {
	method := p.Method

	switch method {
	case "GET":
		return blue
	case "POST":
		return cyan
	case "PUT":
		return yellow
	case "DELETE":
		return red
	case "PATCH":
		return green
	case "HEAD":
		return magenta
	case "OPTIONS":
		return white
	default:
		return reset
	}
}

// ResetColor resets all escape attributes.
func (p *LogFormatterParams) ResetColor() string {
	return reset
}

func NewLoggerMiddlewareSmoothFail(isStdout bool, writerDir string) gin.HandlerFunc {
	res, err := NewLoggerMiddleware(isStdout, writerDir)
	if err != nil {
		logrus.WithError(err).WithField("writerDir", writerDir).Fatal("NewLoggerMiddlewareSmoothFail")
	}

	return res
}

func NewLoggerMiddleware(isStdout bool, writerDir string) (gin.HandlerFunc, error) {
	var writer io.Writer
	if writerDir != "" {
		w, err := logging.NewRotateIO(writerDir, 5)
		if err != nil {
			return nil, err
		}
		writer = w
	}

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		timestamp := time.Now()
		if raw != "" {
			path = path + "?" + raw
		}

		param := LogFormatterParams{
			TimeStamp:    timestamp,
			StatusCode:   c.Writer.Status(),
			Latency:      timestamp.Sub(start),
			ClientIP:     c.ClientIP(),
			Method:       c.Request.Method,
			Path:         path,
			ErrorMessage: c.Errors.ByType(1 << 0).String(),
			BodySize:     c.Writer.Size(),
			Keys:         c.Keys,
		}
		bytes, _ := json.Marshal(param)

		if isStdout {
			fmt.Fprint(os.Stdout, stdoutLogFormatter(param))
		}

		if writer != nil {
			fmt.Fprint(writer, fmt.Sprintln(string(bytes)))
		}
	}, nil
}
