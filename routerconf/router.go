package routerconf

import (
    "frame/controller"
    "frame/middle"
    "gee/framework/router"
)

func init() {
    ins := router.GetInstance()

    //中间件别名
    ins.RouterMiddle("login", &middle.LoginAuth{})
    ins.RouterMiddle("csrf", &middle.CsrfAuth{})

    ins.GET(":lang/doc", &controller.LangController{}, "Doc")
    ins.GET(":lang/tutorial", &controller.LangController{}, "Tutorial")
    ins.GET("hello", &controller.HelloController{}, "Hello")

    loginGroup := []string{"csrf", "login"}
    var emptyGroup []string

    ins.Group("user", loginGroup, func() {
        ins.GET(":id", &controller.UserController{}, "Info")
        ins.GET("list", &controller.UserController{}, "List")
        ins.Group("level", emptyGroup, func() {
            ins.GET("add", &controller.UserLevelController{}, "Add")
            ins.GET(":level", &controller.UserLevelController{}, "GetLevel")
        })
    })
}
