package controller

import "gee/framework/controller"

type UserLevelController struct {
    controller.Controller
}

func (o *UserLevelController) Add() {
    o.Data["msg"] = "进入到 userLevel 的 Add方法"
    o.ServerJson()
}
func (o *UserLevelController) GetLevel() {
    o.Data["msg"] = "进入到 userLevel 的 GetLevel 方法"
    o.Data["level"] =  o.GetRequestParams()["level"]
    o.ServerJson()
}
