package tool

import (
    "bytes"
    "encoding/json"
    "fmt"
    "runtime"
)

func Print(data interface{}) {
    bs, _ := json.Marshal(data)
    var out bytes.Buffer
    _ = json.Indent(&out, bs, "", "  ")
    _, codePath, codeLine, _ := runtime.Caller(1)
    fmt.Println(codePath, codeLine, "行：")
    fmt.Printf(" %v\n", out.String())
}


func Isset(data map[string]string, key string) bool {
    _, ok := data[key]
    return ok
}
