package middle

import (
    "gee/framework/httpContext"
    "gee/tool"
)

type CsrfAuth struct {
}

func (auth *CsrfAuth) Handle(io *httpContext.Io, next func(io *httpContext.Io) func()) (func(), bool) {
    params := io.GetRequestParams()

    if !tool.Isset(params, "csrf_token") {
        data := make(map[string]interface{})
        data["code"] = -1
        data["cost"] = 0
        data["message"] = "参数缺失csrf_token"
        httpContext.ServerJson(data)
        return nil, false
    }

    return next(io), true
}
func (auth *CsrfAuth) GetName() string {
    return "CsrfAuth"
}
