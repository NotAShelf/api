package index

import (
	"github.com/valyala/fasthttp"
    "encoding/json"
)

func IndexHandler(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.SetCanonical([]byte("Content-Type"), []byte("application/json"))
	ctx.Response.SetStatusCode(200)
	response := &IndexResponse{
		Status: 200,
		Data: &IndexData{
			Message: "NotAnAPI v1",
			Link: "https://notashelf.dev",
		},
	}
	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}
