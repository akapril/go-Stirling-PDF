package main

import (
    "Go-Stirling-PDF/api"
    _ "Go-Stirling-PDF/config"
    "fmt"
    "net/http"
    "os"
    "path/filepath"
    "strings"
)

func main() {
    err := http.ListenAndServe(":8888", api.Router)
    if err != nil {
        fmt.Println("服务器开启错误: ", err)
    }
    //renameFile()
}

func renameFile() {
    folderpath := "./lang" //指定文件夹

    oldExtension := "_" //旧后缀

    newExtension := "-" //新后缀

    err := filepath.Walk(folderpath, func(filename string, fi os.FileInfo, err error) error {

        if err != nil {
            return err
        }

        if !fi.IsDir() {

            basename := filepath.Base(filename)
            //ext := filepath.Ext(basename)
            ext := strings.Contains(basename, oldExtension)

            if ext {

                newFilename := strings.ReplaceAll(basename, oldExtension, newExtension)

                return os.Rename(filename, filepath.Join(filepath.Dir(filename), newFilename))
            }
        }

        return nil
    })

    if err != nil {
        panic(err)
    }
}
