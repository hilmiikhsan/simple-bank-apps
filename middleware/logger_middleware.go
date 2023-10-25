package middleware

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/simple-bank-apps/config"
	"github.com/sirupsen/logrus"
)

type RequestLog struct {
	StartTime  time.Time
	Username   string
	Activity   string
	StatusCode int
}

type LogMiddleware interface {
	LogRequestMiddleware(log *logrus.Logger) gin.HandlerFunc
	LogRequestAuth(log *logrus.Logger) gin.HandlerFunc
}

func NewLogMiddleware(log *logrus.Logger, cfg *config.Config) LogMiddleware {
	return &logMiddleware{
		Log: log,
		cfg: cfg,
	}
}

type logMiddleware struct {
	Log *logrus.Logger
	cfg *config.Config
}

func (l *logMiddleware) LogRequestMiddleware(log *logrus.Logger) gin.HandlerFunc {
	customFormatter := &logrus.JSONFormatter{}
	customFormatter.TimestampFormat = time.RFC3339

	log.SetFormatter(customFormatter)

	file, err := os.OpenFile(l.cfg.Log.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("error open file", err.Error())
	}
	log.SetOutput(file)

	return func(c *gin.Context) {
		currentTime := time.Now()
		requestLog := &RequestLog{
			StartTime:  currentTime,
			Username:   GetTokenMiddlewareFromContext(c.Request.Context()).Username,
			Activity:   c.Request.Method + " " + c.Request.URL.Path,
			StatusCode: c.Writer.Status(),
		}

		switch {
		case requestLog.StatusCode >= 500:
			log.WithFields(logrus.Fields{
				"time":        currentTime.Format(time.RFC3339),
				"username":    requestLog.Username,
				"activity":    requestLog.Activity,
				"status_code": requestLog.StatusCode,
			}).Error("Request log")
		case requestLog.StatusCode >= 400:
			log.WithFields(logrus.Fields{
				"time":        currentTime.Format(time.RFC3339),
				"username":    requestLog.Username,
				"activity":    requestLog.Activity,
				"status_code": requestLog.StatusCode,
			}).Warn("Request log")
		default:
			log.WithFields(logrus.Fields{
				"time":        currentTime.Format(time.RFC3339),
				"username":    requestLog.Username,
				"activity":    requestLog.Activity,
				"status_code": requestLog.StatusCode,
			}).Info("Request log")
		}
	}
}

func (l *logMiddleware) LogRequestAuth(log *logrus.Logger) gin.HandlerFunc {
	customFormatter := &logrus.JSONFormatter{}
	customFormatter.TimestampFormat = time.RFC3339

	log.SetFormatter(customFormatter)

	file, err := os.OpenFile(l.cfg.Log.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("error open file", err.Error())
	}
	log.SetOutput(file)

	return func(c *gin.Context) {
		currentTime := time.Now()
		requestLog := &RequestLog{
			StartTime:  currentTime,
			Username:   c.GetString("username"),
			Activity:   c.Request.Method + " " + c.Request.URL.Path,
			StatusCode: c.Writer.Status(),
		}

		switch {
		case requestLog.StatusCode >= 500:
			log.WithFields(logrus.Fields{
				"time":        currentTime.Format(time.RFC3339),
				"username":    requestLog.Username,
				"activity":    requestLog.Activity,
				"status_code": requestLog.StatusCode,
			}).Error("Request log")
		case requestLog.StatusCode >= 400:
			log.WithFields(logrus.Fields{
				"time":        currentTime.Format(time.RFC3339),
				"username":    requestLog.Username,
				"activity":    requestLog.Activity,
				"status_code": requestLog.StatusCode,
			}).Warn("Request log")
		default:
			log.WithFields(logrus.Fields{
				"time":        currentTime.Format(time.RFC3339),
				"username":    requestLog.Username,
				"activity":    requestLog.Activity,
				"status_code": requestLog.StatusCode,
			}).Info("Request log")
		}
	}
}
