package main

import (
	"Go-Stirling-PDF/api"
	_ "Go-Stirling-PDF/config"
	"fmt"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8888", api.Router)
	if err != nil {
		fmt.Println("服务器开启错误: ", err)
	}
}
