package router

import (
	"net/http"

	"github.com/ahmadtheswe/carbon-footprint-api/handler"
)

func SetupRoutes() {
	http.HandleFunc("/", handler.HelloHandler)
}
