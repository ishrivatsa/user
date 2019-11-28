package tracer

import (
	//"fmt"

	"github.com/gin-gonic/gin"
	stdopentracing "github.com/opentracing/opentracing-go"
	tracelog "github.com/opentracing/opentracing-go/log"
	"github.com/vmwarecloudadvocacy/user/pkg/logger"
)

func CreateTracerAndSpan(spanName string, c *gin.Context) (stdopentracing.Span, error) {

	tracer := stdopentracing.GlobalTracer()

	userSpanCtx, err := tracer.Extract(stdopentracing.HTTPHeaders, stdopentracing.HTTPHeadersCarrier(c.Request.Header))
	
	if err !=nil {
		logger.Logger.Infof(err.Error())
	}

	userSpan := tracer.StartSpan(spanName, stdopentracing.ChildOf(userSpanCtx))
	defer userSpan.Finish()
	return userSpan, nil
}

func OnErrorLog(receivedSpan stdopentracing.Span, err error) {

	receivedSpan.LogFields(
		tracelog.String("event", "error"),
		tracelog.String("message", err.Error()),
	)
	logger.Logger.Infof(err.Error())
}