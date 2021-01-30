package fasthttpserver

import (
	"github.com/valyala/fasthttp"
)

func ReadinessHandler(requestCtx *fasthttp.RequestCtx) {
	requestCtx.SetStatusCode(200)
	requestCtx.SetContentType("application/json")
	requestCtx.SetBodyString("OK")
}
