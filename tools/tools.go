package tools

import "github.com/gin-gonic/gin"

type ResultData struct {
    Context *gin.Context `json:"context"`
    Data    any          `json:"data"`
}

func (res *ResultData) ResultDataFunc() {

}
