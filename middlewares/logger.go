package middlewares

import (
	"github.com/manyminds/api2go"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type AccessLogger struct {
	ZapLogger *zap.SugaredLogger
}

func (ac *AccessLogger) AccessLogMiddleware(ctx api2go.APIContexter, w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer ac.ZapLogger.Info(r.URL.Path,
		zap.String("method", r.Method),
		zap.String("remote_addr", r.RemoteAddr),
		zap.String("url", r.URL.Path),
		zap.String("work_time", time.Since(start).String()),
	)
}
