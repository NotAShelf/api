package main

import (
	"fmt"
	"log"
	"os"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	_ "github.com/joho/godotenv/autoload"

	"notashelf.dev/api/src/routes"
)

func main() {
	r := router.New()
	routes.InitRoutes(r)
	port := os.Getenv("PORT")
	fmt.Println("API on port: " + port)
	log.Fatal(fasthttp.ListenAndServe(":" + port, r.Handler))
}
