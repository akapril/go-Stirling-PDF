package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
}

func (h *AccountHandler) AccountHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "account.html", gin.H{})
}
