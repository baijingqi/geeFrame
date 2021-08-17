package controller

import (
    "fmt"
    "gee/framework/controller"
)

// LangController Operations about object
type LangController struct {
    controller.Controller
}

func (c *LangController) Doc() {
    c.Data["req"] = c.GetRequestParams()
    c.ServerJson()
}

func (c *LangController) Tutorial() {
    fmt.Println(c.GetRequestParams())
}
