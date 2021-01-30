package fasthttpserver

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type Response struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error,omitempty"`
}

type ResponseBuilder struct {
	ctx  *fasthttp.RequestCtx
	resp Response
}

func NewResponseBuilder(ctx *fasthttp.RequestCtx) *ResponseBuilder {
	return &ResponseBuilder{
		ctx: ctx,
		resp: Response{
			Data: make(map[int]struct{}),
		},
	}
}

func (r *ResponseBuilder) SetStatusCode(code int) *ResponseBuilder {
	r.ctx.Response.SetStatusCode(code)

	return r
}

func (r *ResponseBuilder) SetData(data interface{}) *ResponseBuilder {
	if data == nil {
		return r
	}

	r.resp.Data = data

	return r
}

func (r *ResponseBuilder) SetError(errorMessage string) *ResponseBuilder {
	r.resp.Error = errorMessage

	return r
}

func (r *ResponseBuilder) Send() {
	body, _ := json.Marshal(r.resp)

	r.ctx.Response.SetBody(body)
	r.ctx.Response.Header.SetContentType("application/json")
}
