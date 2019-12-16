package requestid

import (
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	xRequestIDKey = "X-Request-ID"
)

// generator a function type that returns string.
type generator func() string

var (
	random = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
)

func uuid(len int) string {
	bytes := make([]byte, len)
	random.Read(bytes)
	return base64.StdEncoding.EncodeToString(bytes)[:len]
}

//RequestID is a middleware that injects a 'RequestID' into the context and header of each request.
func RequestID(gen generator) gin.HandlerFunc {
	return func(c *gin.Context) {
		var xRequestID string
		if gen != nil {
			xRequestID = gen()
		} else {
			xRequestID = uuid(16)
		}
		c.Request.Header.Set(xRequestIDKey, xRequestID)
		c.Set(xRequestIDKey, xRequestID)
		fmt.Printf("[GIN-debug] %s [%s] - \"%s %s\"\n", time.Now().Format(time.RFC3339), xRequestID, c.Request.Method, c.Request.URL.Path)
		c.Next()
	}
}

// GetLoggerConfig return gin.LoggerConfig which will write the logs to specified io.Writer with given gin.LogFormatter.
// By default gin.DefaultWriter = os.Stdout
// reference: https://github.com/gin-gonic/gin#custom-log-format
func GetLoggerConfig(formatter gin.LogFormatter, output io.Writer, skipPaths []string) gin.LoggerConfig {
	if formatter == nil {
		formatter = GetDefaultLogFormatterWithRequestID()
	}
	return gin.LoggerConfig{
		Formatter: formatter,
		Output:    output,
		SkipPaths: skipPaths,
	}
}

//GetDefaultLogFormatterWithRequestID returns gin.LogFormatter with 'RequestID'
func GetDefaultLogFormatterWithRequestID() gin.LogFormatter {
	return func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[GIN-debug] %s [%s] - [%s] \"%s %s %s %d %s\" %s %s\n",
			param.TimeStamp.Format(time.RFC3339),
			param.Request.Header.Get(xRequestIDKey),
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}
}

// GetRequestIDFromContext returns 'RequestID' from the given context if present.
func GetRequestIDFromContext(c *gin.Context) string {
	if v, ok := c.Get(xRequestIDKey); ok {
		if requestID, ok := v.(string); ok {
			return requestID
		}
	}
	return ""
}

// GetRequestIDFromHeaders returns 'RequestID' from the headers if present.
func GetRequestIDFromHeaders(c *gin.Context) string {
	return c.Request.Header.Get(xRequestIDKey)
}
