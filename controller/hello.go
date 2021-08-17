package controller

import (
    "gee/framework/controller"
)

// Operations about object
type HelloController struct {
    controller.Controller
}

func (c *HelloController) Hello() {
    c.Data["name"] = "gee"
    c.Data["version"] = "1.0.0"
    c.ServerJson()
}