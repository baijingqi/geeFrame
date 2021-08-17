package gee

import (
    "fmt"
    "gee/framework/httpContext"
    "gee/framework/router"
    "net/http"
)

type App struct {
    w http.ResponseWriter
    r http.Request
}

func New() *App {
    return &App{}
}

func (app *App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    httpContext.BindIo(w, req)
    router.Handle()
}

func (app *App) Run(s string) {
    fmt.Println("开始运行")
    _ = http.ListenAndServe(s, app)
}
