package api

import (
    "Go-Stirling-PDF/handler"
    "Go-Stirling-PDF/server"
    "Go-Stirling-PDF/tools"
    "bytes"
    ginI18n "github.com/gin-contrib/i18n"
    "github.com/goccy/go-json"
    "github.com/spf13/viper"
    "golang.org/x/text/language"
    "log"
    "net/http"
    "text/template"

    "github.com/gin-gonic/gin"
)

var Router *gin.Engine

func GinI18nLocalize() gin.HandlerFunc {
    return ginI18n.Localize(
        ginI18n.WithBundle(&ginI18n.BundleCfg{
            RootPath:         "./lang",
            AcceptLanguage:   []language.Tag{language.SimplifiedChinese, language.BritishEnglish, language.AmericanEnglish},
            DefaultLanguage:  language.AmericanEnglish,
            FormatBundleFile: "properties",
            UnmarshalFunc: func(data []byte, v interface{}) error {
                //toml.Unmarshal(data, v)
                viper.SetConfigType("properties")
                viper.ReadConfig(bytes.NewBuffer(data))
                err := viper.Unmarshal(&v)
                cdata := make(map[string]interface{})
                marshal, err := json.Marshal(v)
                err = json.Unmarshal(marshal, &cdata)
                if err != nil {
                    log.Println(err)
                }
                for key, val := range cdata {
                    log.Println(key, val)
                }
                return err
            },
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

type CardData struct {
    Id        string `json:"id"`
    CardTitle string `json:"cardTitle"`
    CardText  string `json:"cardText"`
    CardLink  string `json:"cardLink"`
    SvgPath   string `json:"svgPath"`
    Tags      string `json:"tags"`
}

func templateFuncs() {
    Router.SetFuncMap(template.FuncMap{
        "Localize": func(key string, c tools.ResultData) (string, error) {
            for s, a := range c.Context.Keys {
                log.Println(s)
                log.Println(a)
            }
            return ginI18n.GetMessage(c.Context, key)
        },
        "cardData": func(key string, c tools.ResultData) CardData {
            data := CardData{}
            err := json.Unmarshal([]byte(key), &data)
            if err != nil {
                log.Println(err)
            }
            data.CardTitle, _ = ginI18n.GetMessage(c.Context, data.CardTitle)
            data.CardText, _ = ginI18n.GetMessage(c.Context, data.CardText)
            data.Tags, _ = ginI18n.GetMessage(c.Context, data.Tags)
            return data
        },
    })
}
func init() {
    Router = gin.Default()
    Router.Use(GinI18nLocalize())
    // 国际化
    templateFuncs()
    //server.I18n(Router)
    server.LoadTmpl(Router)
    handlers := handler.Handler{}
    homeHandler := handlers.HomeHandler
    accountHandler := handlers.AccountHandler
    //Router.HTMLRender = render()

    views := Router.Group("/")
    views.GET("/", homeHandler.HomeHandler)
    views.GET("/home", homeHandler.HomeHandler)
    views.GET("/account", accountHandler.AccountHandler)
}

func Listen(w http.ResponseWriter, r *http.Request) {
    Router.ServeHTTP(w, r)
}
