package context

import (
	"context"

	"github.com/sirupsen/logrus"
)

type Context struct {
	Logger *logrus.Logger
	fields map[string]interface{}
	req    interface{}
	ctx    context.Context
}

func NewContext(requestID string, logger *logrus.Logger) *Context {
	return &Context{
		Logger: logger.WithFields(logrus.Fields{
			"request_id": requestID,
		}).Logger,
		fields: map[string]interface{}{},
		ctx:    context.WithValue(context.Background(), "request_id", requestID),
	}
}

func (ctx *Context) AddLogField(key string, value interface{}) {
	ctx.fields[key] = value
}

func (ctx *Context) AddLogFields(fields map[string]interface{}) {
	for key, value := range fields {
		ctx.fields[key] = value
	}
}

func (ctx *Context) SetReq(req interface{}) {
	ctx.req = req
}

func (ctx *Context) Req() interface{} {
	return ctx.req
}

func (ctx *Context) LogFields() map[string]interface{} {
	return ctx.fields
}

func (ctx *Context) Ctx() context.Context {
	if ctx.ctx == nil {
		ctx.ctx = context.Background()
	}
	return ctx.ctx
}
