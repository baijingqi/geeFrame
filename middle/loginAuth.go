package middle

import (
    "gee/framework/httpContext"
    "gee/tool"
)

type LoginAuth struct {
}

func (auth *LoginAuth) Handle(io *httpContext.Io, next func(io *httpContext.Io) func()) (func(), bool) {
    params := io.GetRequestParams()

    if !tool.Isset(params, "uid") {
        data := make(map[string]interface{})
        data["code"] = -1
        data["cost"] = 0
        data["message"] = "参数缺失uid"
        httpContext.ServerJson(data)
        return nil, false
    }
    return next(io), true
}

func (auth *LoginAuth) GetName() string {
    return "LoginAuth"
}
