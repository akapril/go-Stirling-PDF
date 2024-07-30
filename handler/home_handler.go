package handler

import (
    "Go-Stirling-PDF/tools"
    "net/http"

    "github.com/gin-gonic/gin"
)

type HomeHandler struct {
}

func (h *HomeHandler) HomeHandler(c *gin.Context) {
    //log.Println()
    //message, err := ginI18n.GetMessage(c, "home.pipeline.title")
    //log.Printf("%s: %s", message, err)
    //{{ Localize 'home.pipeline.title' .}}
    c.Set("title", "message")
    c.HTML(http.StatusOK, "home.html", tools.ResultData{Context: c, Data: gin.H{"title": "message"}})
}
