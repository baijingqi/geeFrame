package middle

import "gee/framework/httpContext"

type Interface interface {
    Handle(io *httpContext.Io, next func(io *httpContext.Io) func()) (func(), bool)
    GetName() string
}

type Struct struct {
}
