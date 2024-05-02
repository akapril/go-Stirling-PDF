package api

import (
	"Go-Stirling-PDF/handler"
	"Go-Stirling-PDF/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
	server.LoadTmpl(Router)
	handlers := handler.Handler{}

	homeHandler := handlers.HomeHandler
	views := Router.Group("/")
	views.GET("/home.html", homeHandler.HomeHandler)
}

func Listen(w http.ResponseWriter, r *http.Request) {
	Router.ServeHTTP(w, r)
}
