package router

//import (
//    "fmt"
//    "gee/framework/Controller"
//    "gee/framework/httpContext"
//    "net/http"
//    "reflect"
//    "strings"
//)
//
//var instance *Router
//
//type Router struct {
//    PostRouter   map[string]Node
//    GetRouter    map[string]Node
//    PutRouter    map[string]Node
//    DeleteRouter map[string]Node
//    Groups       map[string]Node
//}
//
//type Node struct {
//    Controller Controller.Interface
//    Action string
//    Group  string
//}
//
//func GetInstance() *Router {
//    if instance != nil {
//        return instance
//    }
//    instance = &Router{PostRouter: make(map[string]Node), GetRouter: make(map[string]Node), PutRouter: make(map[string]Node), DeleteRouter: make(map[string]Node)}
//    return instance
//}
//
//func (Router *Router) addRoute(method string, urlPath string, Controller Controller.Interface, Action string) {
//    if instance == nil {
//        GetInstance()
//    }
//    if urlPath[0] == '/' {
//        urlPath = strings.TrimLeft(urlPath, "/")
//    }
//    switch method {
//    case "POST":
//        instance.PostRouter[urlPath] = Node{Controller, Action, ""}
//    case "GET":
//        instance.GetRouter[urlPath] = Node{Controller, Action, ""}
//    case "PUT":
//        instance.PutRouter[urlPath] = Node{Controller, Action, ""}
//    case "DELETE":
//        instance.DeleteRouter[urlPath] = Node{Controller, Action, ""}
//    }
//
//}
//
//func (Router *Router) GET(urlPath string, Controller Controller.Interface, Action string) {
//    Router.addRoute("GET", urlPath, Controller, Action)
//}
//
//func (Router *Router) POST(urlPath string, Controller Controller.Interface, Action string) {
//    Router.addRoute("POST", urlPath, Controller, Action)
//}
//
//func (Router *Router) PUT(urlPath string, Controller Controller.Interface, Action string) {
//    Router.addRoute("PUT", urlPath, Controller, Action)
//}
//
//func (Router *Router) DELETE(urlPath string, Controller Controller.Interface, Action string) {
//    Router.addRoute("DELETE", urlPath, Controller, Action)
//}
//
//func Handle() {
//    httpIo := httpContext.GetIo()
//    params := parseParams(httpIo.R)
//    path := parseUrl(httpIo.R)
//    method := parseMethod(httpIo.R)
//    handleByController(method, path, params)
//}
//
//func parseParams(req *http.Request) map[string]interface{} {
//    _ = req.ParseForm()
//    params := make(map[string]interface{})
//
//    for k, v := range req.Form {
//        //fmt.Println("v=", v, "typeof(v)=", reflect.TypeOf(v))
//        params[k] = v
//    }
//    return params
//}
//
//func parseUrl(req *http.Request) string {
//    return req.URL.Path
//}
//
//func parseMethod(req *http.Request) string {
//    return req.Method
//}
//
//func handleByController(method string, path string, params map[string]interface{}) {
//
//    if path[0] == '/' {
//        path = strings.TrimLeft(path, "/")
//    }
//    findRouter := make(map[string]Node)
//    switch method {
//    case "POST":
//        findRouter = instance.PostRouter
//    case "GET":
//        findRouter = instance.GetRouter
//    case "PUT":
//        findRouter = instance.PutRouter
//    case "DELETE":
//        findRouter = instance.DeleteRouter
//    }
//    fmt.Println("findRouter=", findRouter)
//    if _, ok := findRouter[path]; ok {
//        routerNode := findRouter[path]
//        fmt.Println("routerNode=", routerNode)
//
//        //reflect.ValueOf(routerNode).MethodByName(routerNode.Action) != 0
//        fmt.Println("Action=", routerNode.Action)
//        reflect.ValueOf(routerNode.Controller).MethodByName(routerNode.Action).Call([]reflect.Value{})
//    } else {
//        fmt.Print("路由未找到", path)
//        httpContext.NotFound()
//    }
//}
