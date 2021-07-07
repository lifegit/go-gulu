package koaMiddleware

import (
	"encoding/json"
	"fmt"
	"github.com/lifegit/go-gulu/v2/nice/koa"
	"github.com/lifegit/go-gulu/v2/pkg/logging"
	"io"
	"log"
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
	TimeStamp   time.Time `json:"time"`
	ServiceName string    `json:"service_name"`
	Result      koa.Data  `json:"result"`
}

var stdoutLogFormatter = func(param LogFormatterParams) string {
	statusCode, statusColor, statusContent := func() (string, string, interface{}) {
		if param.Result.Err != nil {
			return "fail", yellow, param.Result.Err.Error()
		}
		return "success", green, param.Result.Data
	}()

	return fmt.Sprintf("[MEDIEM %s] %v |%s %s %s| %v\n",
		param.ServiceName,
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, statusCode, reset,
		statusContent,
	)
}

func NewLoggerMiddlewareSmoothFail(isStdout, isWriter bool, serviceName string, writerDir string) koa.HandlerFunc {
	res, err := NewLoggerMiddleware(isStdout, isWriter, serviceName, writerDir)
	if err != nil {
		log.Fatalf("NewLoggerMiddlewareSmoothFail: %s. err: %s", writerDir, err)
	}

	return res
}

func NewLoggerMiddleware(isStdout, isWriter bool, serviceName string, writerDir string) (koa.HandlerFunc, error) {
	var writer io.Writer
	if isWriter {
		w, err := logging.NewRotateIO(writerDir, 5)
		if err != nil {
			return nil, err
		}
		writer = w
	}

	return func(c *koa.Context) {
		c.Next()

		timestamp := time.Now()

		param := LogFormatterParams{
			TimeStamp:   timestamp,
			ServiceName: serviceName,
			Result:      c.Result,
		}
		bytes, _ := json.Marshal(param)

		if isStdout {
			fmt.Fprint(os.Stdout, stdoutLogFormatter(param))
		}

		if isWriter {
			fmt.Fprint(writer, fmt.Sprintln(string(bytes)))
		}
	}, nil
}
