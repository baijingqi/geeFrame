package controller

import (
    "gee/framework/controller"
)

// UserController Operations about object
type UserController struct {
    controller.Controller
}

func (o *UserController) Info() {
    res := o.GetRequestParams()
    o.Data["id"] = res["id"]
    o.Data["name"] = "xxxx"
    o.ServerJson()
}
func (o *UserController) List() {
    o.Data["msg"] = "进入到userlist"
    o.ServerJson()
}
