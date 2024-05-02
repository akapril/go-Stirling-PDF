package server

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
}

//go:embed views/*
var EmbededFiles embed.FS

func LoadTmpl(param *gin.Engine) {

	// param.StaticFS("/views/static", http.Dir("server/views/static"))
	param.StaticFS("/static", http.Dir("server/views/static"))
	param.StaticFS("/css", http.Dir("server/views/static/css"))
	param.StaticFS("/files", http.Dir("server/views/static/files"))
	param.StaticFS("/fonts", http.Dir("server/views/static/fonts"))
	param.StaticFS("/images", http.Dir("server/views/static/images"))
	param.StaticFS("/js", http.Dir("server/views/static/js"))
	param.StaticFS("/pdfjs", http.Dir("server/views/static/pdfjs"))

	param.Delims("{{{", "}}}")
	param.LoadHTMLGlob("server/views/templates/**/*")
}
