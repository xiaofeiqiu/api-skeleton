package logger

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

var jsonLogger *logrus.Logger
var textLogger *logrus.Logger

func init() {
	jsonLogger = initJsonLogger()
	textLogger = initTextLogger()
}

type StructuredLogger struct {
	Logger *logrus.Logger
}

type StructuredLoggerEntry struct {
	Logger logrus.FieldLogger
}

func initJsonLogger() *logrus.Logger {
	logger := logrus.New()
	logLevel := logrus.InfoLevel

	tmp := os.Getenv("LOG_LEVEL")
	if strings.EqualFold(tmp, "info") {
		logLevel = logrus.InfoLevel
	} else if strings.EqualFold(tmp, "debug") {
		logLevel = logrus.DebugLevel
	}

	filedMap := logrus.FieldMap{
		logrus.FieldKeyTime:  "time",
		logrus.FieldKeyLevel: "loglevel",
	}

	logger.Formatter = new(logrus.JSONFormatter)
	logger.Formatter.(*logrus.JSONFormatter).DisableTimestamp = false
	logger.Formatter.(*logrus.JSONFormatter).TimestampFormat = time.RFC3339Nano
	logger.Formatter.(*logrus.JSONFormatter).FieldMap = filedMap
	logger.Formatter.(*logrus.JSONFormatter).DisableHTMLEscape = true
	logger.Level = logLevel
	logger.Out = os.Stdout
	return logger
}

func initTextLogger() *logrus.Logger {
	logger := logrus.New()
	logLevel := logrus.InfoLevel

	tmp := os.Getenv("LOG_LEVEL")
	if strings.EqualFold(tmp, "info") {
		logLevel = logrus.InfoLevel
	} else if strings.EqualFold(tmp, "debug") {
		logLevel = logrus.DebugLevel
	} else if strings.EqualFold(tmp, "error") {
		logLevel = logrus.ErrorLevel
	}

	logger.Formatter = new(logrus.TextFormatter)
	logger.Formatter.(*logrus.TextFormatter).DisableTimestamp = false
	logger.Formatter.(*logrus.TextFormatter).TimestampFormat = time.RFC3339
	logger.Level = logLevel
	logger.Out = os.Stdout
	return logger
}

func (l *StructuredLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &StructuredLoggerEntry{logrus.NewEntry(l.Logger)}
	logFields := logrus.Fields{}

	if reqID := middleware.GetReqID(r.Context()); reqID != "" {
		logFields["requestId"] = reqID
	}

	logFields["method"] = r.Method
	logFields["path"] = r.URL.Path

	if r.URL.RawQuery != "" {
		logFields["path"] = fmt.Sprintf("%s?%s", r.URL.Path, r.URL.RawQuery)
	}

	entry.Logger = entry.Logger.WithFields(logFields)
	entry.Logger.Info("request started")
	return entry
}

func (l *StructuredLoggerEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	l.Logger = l.Logger.WithFields(logrus.Fields{
		"status": status,
	})

	l.Logger.Infoln("request complete")
}

func (l *StructuredLoggerEntry) Panic(v interface{}, stack []byte) {
	l.Logger = l.Logger.WithFields(logrus.Fields{
		"stack": string(stack),
		"panic": fmt.Sprintf("%+v", v),
	})
}

func NewMiddlewareLogger() func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&StructuredLogger{jsonLogger})
}

func NewJsonLogger() *logrus.Logger {
	return jsonLogger
}

func NewTextLogger() *logrus.Logger {
	return textLogger
}

func Info(event string, details string) {
	jsonLogger.WithField("event", event).Info(details)
}

func Error(event string, err error, details string) {
	jsonLogger.WithField("event", event).Error(fmt.Sprintf("%s, %s", details, err.Error()))
}

func Fatal(event string, details string) {
	jsonLogger.WithField("event", event).Fatal(details)
}
