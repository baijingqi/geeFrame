package httpContext

import (
    "encoding/json"
    "net/http"
)

type Io struct {
    W             http.ResponseWriter
    R             *http.Request
    requestParams map[string]string
}

var HttpOk = 200
var HttpNotFound = 404
var HttpServerError = 500

var httpIo Io

func BindIo(w http.ResponseWriter, r *http.Request) {
    httpIo = Io{w, r, make(map[string]string)}
}

func (httpIo Io) BindRequestParams(params map[string]string) {
    for key, value := range params {
        httpIo.requestParams[key] = value
    }
}
func (httpIo Io) GetRequestParams() map[string]string {
    return httpIo.requestParams
}
func (httpIo Io) Init() {
    httpIo.requestParams = nil
}

func GetIo() *Io {
    return &httpIo
}

func ServerJson(data interface{}) {
    httpIo.W.Header().Set("Content-Type", "application/json; charset=utf-8")
    httpIo.W.WriteHeader(HttpOk)
    content, _ := json.Marshal(data)
    _, _ = httpIo.W.Write(content)
}

func Response(code int, data map[interface{}]interface{}) {

}

func NotFound() {
    httpIo.W.Header().Set("Content-Type", "text/html")
    httpIo.W.WriteHeader(HttpNotFound)
    msg := "GeeMsg:" + httpIo.R.URL.Path + " Not Found"
    _, _ = httpIo.W.Write([]byte(msg))
}
