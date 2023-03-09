package routes

import (
	"github.com/fasthttp/router"
	"github.com/jckli/api/src/handlers/index"
)

func InitRoutes(r *router.Router) {
	r.GET("/", index.IndexHandler)
}
