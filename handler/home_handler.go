package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeHandler struct {
}

func (h *HomeHandler) HomeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "home", gin.H{"title": "go-Stirling-PDF"})
}
