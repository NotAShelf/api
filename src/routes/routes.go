package routes

import (
	"github.com/fasthttp/router"
	"notashelf.dev/api/src/handlers/index"
)

func InitRoutes(r *router.Router) {
	r.GET("/", index.IndexHandler)
}
