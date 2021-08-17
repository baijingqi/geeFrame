package controller

import "C"
import (
    "gee/framework/httpContext"
    "time"
)

type Interface interface {
    Init()
    BindData(key string, value interface{})
}

type Controller struct {
    Data           map[string]interface{}
    ControllerName string
    ActionName     string
    Io             *httpContext.Io
    startDealTime  int
    cost           float64
}

func (c *Controller) Init() {
    c.startDealTime = time.Now().Nanosecond()
    c.Data = make(map[string]interface{})
    c.Data["code"] = 200
    c.Data["cost"] = 0
    c.Data["message"] = ""
}

func (c *Controller) ServerJson() {
    c.cost = float64(time.Now().Nanosecond()-c.startDealTime) / float64(100000)
    c.Data["cost"] = c.cost
    httpContext.ServerJson(c.Data)
}

func (c *Controller) BindData(key string, value interface{}) {
    c.Data[key] = value
}

func (c *Controller) GetRequestParams() map[string]string {
    if c.Io == nil {
        c.Io = httpContext.GetIo()
    }
    return c.Io.GetRequestParams()
}
func (c *Controller) Result(data map[string]interface{}, code int, message string) {
    c.Data["data"] = data
    c.Data["code"] = code
    c.Data["message"] = message
    httpContext.ServerJson(c.Data)
    c.Data = nil
}
