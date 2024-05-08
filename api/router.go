package api

import (
	"Go-Stirling-PDF/handler"
	"Go-Stirling-PDF/server"
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/pelletier/go-toml/v2"
	"golang.org/x/text/language"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

/*
	func render() multitemplate.Renderer {
	    r := multitemplate.NewRenderer()
	    r.AddFromFiles("home", "./server/views/home/home.html", "./server/views/fragments/common.html")
	    return r
	}
*/
func init() {
	Router = gin.Default()
	server.LoadTmpl(Router)
	// 国际化
	Router.Use(GinI18nLocalize())
	Router.SetFuncMap(template.FuncMap{
		"Localize": ginI18n.GetMessage,
	})

	handlers := handler.Handler{}
	homeHandler := handlers.HomeHandler
	accountHandler := handlers.AccountHandler
	//Router.HTMLRender = render()

	views := Router.Group("/")
	views.GET("/", homeHandler.HomeHandler)
	views.GET("/home", homeHandler.HomeHandler)
	views.GET("/account", accountHandler.AccountHandler)
}

func GinI18nLocalize() gin.HandlerFunc {
	return ginI18n.Localize(
		ginI18n.WithBundle(&ginI18n.BundleCfg{
			RootPath:         "./lang",
			AcceptLanguage:   []language.Tag{language.Chinese, language.English},
			DefaultLanguage:  language.Chinese,
			FormatBundleFile: "properties",
			UnmarshalFunc:    toml.Unmarshal,
		}),
		ginI18n.WithGetLngHandle(
			func(context *gin.Context, defaultLng string) string {
				lng := context.Query("lang")
				if lng == "" {
					return defaultLng
				}
				return lng
			},
		),
	)
}
func Listen(w http.ResponseWriter, r *http.Request) {
	Router.ServeHTTP(w, r)
}
