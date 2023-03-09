package index

import (
	"github.com/valyala/fasthttp"
)

func IndexHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html")
	ctx.WriteString("Hello World!")
}
