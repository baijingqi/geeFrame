package main

import (
    "fmt"
    "github.com/liudng/godump"
)

type Middle interface {
    Handle(request map[string]string, next func(request map[string]string) func()) func()
    GetName() string
}
type VerifyCsrf struct {
}

func (VerifyCsrf *VerifyCsrf) Handle(request map[string]string, next func(request map[string]string) func()) func() {
    fmt.Println("调用了 VerifyCsrf")
    request["VerifyCsrf"] = "1"
    return next(request)
}
func (VerifyCsrf *VerifyCsrf) GetName() string {
    return "VerifyCsrf"
}

type VerifyAuth struct {
}

func (VerifyAuth *VerifyAuth) Handle(request map[string]string, next func(request map[string]string) func()) func() {
    fmt.Println("调用了 VerifyAuth")
    request["VerifyAuth"] = "1"

    return next(request)
}
func (VerifyAuth *VerifyAuth) GetName() string {
    return "VerifyAuth"
}

type SetCookie struct {
}

func (SetCookie *SetCookie) Handle(request map[string]string, next func(request map[string]string) func()) func() {
    fmt.Println("调用了 setCookie")
    request["SetCookie"] = "1"
    return next(request)
}
func (SetCookie *SetCookie) GetName() string {
    return "SetCookie"
}
func main() {
    arr := []Middle{&VerifyAuth{}, &VerifyCsrf{}, &SetCookie{}}
    request := make(map[string]string)
    request["uid"] = "10000"
    callback := func(request map[string]string) func() {
       return nil
    }

    godump.Dump(request)
    for i := len(arr) -1; i >= 0; i-- {
       func(middle Middle) {
           //fmt.Println(middle.GetName())
           tmpCallback := callback
           callback = func(request map[string]string) func() {
                middle.Handle(request, tmpCallback)
                return nil
           }
       }(arr[i])
    }
    callback(request)
    godump.Dump(request)
}
