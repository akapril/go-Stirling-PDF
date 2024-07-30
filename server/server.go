package server

import (
    "bytes"
    "embed"
    ginI18n "github.com/gin-contrib/i18n"
    "github.com/spf13/viper"
    "golang.org/x/text/language"
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

    param.LoadHTMLGlob("server/views/templates/**/*")
}

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
