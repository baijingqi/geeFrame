package router

import (
    "fmt"
    "gee/framework/controller"
    "gee/framework/httpContext"
    "gee/framework/middle"
    "github.com/liudng/godump"
    "net/http"
    "reflect"
    "strings"
)

var instance *Router

type Router struct {
    PostRouter         map[string]Node
    GetRouter          map[string]Node
    PutRouter          map[string]Node
    DeleteRouter       map[string]Node
    Groups             map[string]Node
    DynamicRouter      TreeNode
    Middles            map[string]middle.Interface
    currentMiddles     []string
    currentGroupPrefix string
}

type Node struct {
    Controller controller.Interface
    Action     string
    Middles    []string
}

func GetInstance() *Router {
    if instance != nil {
        return instance
    }
    instance = &Router{
        PostRouter:    make(map[string]Node),
        GetRouter:     make(map[string]Node),
        PutRouter:     make(map[string]Node),
        DeleteRouter:  make(map[string]Node),
        DynamicRouter: TreeNode{Pattern: "/", Children: make(map[string]*TreeNode)},
        Middles:       make(map[string]middle.Interface),
    }
    return instance
}

func (Router *Router) addRoute(method string, urlPath string, controller controller.Interface, action string) {
    if instance == nil {
        GetInstance()
    }
    urlPath = strings.TrimLeft(urlPath, "/")

    if Router.currentGroupPrefix != "" {
        urlPath = Router.currentGroupPrefix + "/" + urlPath
    }

    isDynamic := isDynamicRouter(urlPath)
    if isDynamic {
        instance.DynamicRouter.Add(method, urlPath, controller, action, Router.currentMiddles)
    } else {
        switch method {
        case "POST":
            instance.PostRouter[urlPath] = Node{controller, action, Router.currentMiddles}
        case "GET":
            instance.GetRouter[urlPath] = Node{controller, action, Router.currentMiddles}
        case "PUT":
            instance.PutRouter[urlPath] = Node{controller, action, Router.currentMiddles}
        case "DELETE":
            instance.DeleteRouter[urlPath] = Node{controller, action, Router.currentMiddles}
        }
    }

    godump.Dump(instance.DynamicRouter)
}

func isDynamicRouter(urlPath string) bool {
    urlPath = strings.TrimLeft(urlPath, "/")
    arr := strings.Split(urlPath, "/")
    isDynamic := false
    for _, value := range arr {
        if value[0] == ':' {
            isDynamic = true
            break
        }
    }
    return isDynamic
}

func (Router *Router) GET(urlPath string, controller controller.Interface, action string) {
    Router.addRoute("GET", urlPath, controller, action)
}

func (Router *Router) POST(urlPath string, controller controller.Interface, action string) {
    Router.addRoute("POST", urlPath, controller, action)
}

func (Router *Router) PUT(urlPath string, controller controller.Interface, action string) {
    Router.addRoute("PUT", urlPath, controller, action)
}

func (Router *Router) DELETE(urlPath string, controller controller.Interface, action string) {
    Router.addRoute("DELETE", urlPath, controller, action)
}

func (Router *Router) Group(name string, middleAlias []string, function func()) {
    for _,i := range middleAlias {
        Router.currentMiddles = append(Router.currentMiddles, i)
    }

    if Router.currentGroupPrefix == "" {
        Router.currentGroupPrefix = name
    } else {
        Router.currentGroupPrefix += "/" + name
    }

    function()
    Router.currentMiddles = []string{}
    Router.currentGroupPrefix = ""
}

func (Router *Router) RouterMiddle(name string, middle middle.Interface) {
    Router.Middles[name] = middle
}

func Handle() {
    httpIo := httpContext.GetIo()
    httpIo.Init()
    params := parseParams(httpIo.R)
    path := parseUrl(httpIo.R)
    method := parseMethod(httpIo.R)
    httpIo.BindRequestParams(params)
    handleByController(method, path, params)
}

func parseParams(req *http.Request) map[string]string {
    _ = req.ParseForm()
    params := make(map[string]string)

    for k, v := range req.Form {
        params[k] = strings.Join(v, "")
    }
    return params
}

func parseUrl(req *http.Request) string {
    return req.URL.Path
}

func parseMethod(req *http.Request) string {
    return req.Method
}

func handleByController(method string, path string, params map[string]string) {
    if path[0] == '/' {
        path = strings.TrimLeft(path, "/")
    }

    findRouter := make(map[string]Node)
    switch method {
    case "POST":
        findRouter = instance.PostRouter
    case "GET":
        findRouter = instance.GetRouter
    case "PUT":
        findRouter = instance.PutRouter
    case "DELETE":
        findRouter = instance.DeleteRouter
    }

    if _, ok := findRouter[path]; ok {
        node := findRouter[path]
        HandleNode(node)
    } else {
        node, routerParams := findDynamicRouter(method, path)
        if node != nil {
            for key, value := range routerParams {
                params[key] = value
            }
            httpContext.GetIo().BindRequestParams(params)
            HandleDynamicRouteNode(node)
        } else {
            httpContext.NotFound()
            fmt.Println("path=", path, " 404 ")
        }
    }
}

func findDynamicRouter(method string, path string) (node *TreeNode, params map[string]string) {
    routerNode, params := instance.DynamicRouter.FindRouter(method, path)
    return routerNode, params
}

func HandleNode(node Node) {
    function := func(node Node) func() {
        return func() {
            node.Controller.Init()
            reflect.ValueOf(node.Controller).MethodByName(node.Action).Call([]reflect.Value{})
        }
    }(node)

    ThroughMiddlesToController(node.Middles, function)
}

func HandleDynamicRouteNode(node *TreeNode) {
    function := func(node *TreeNode) func() {
        return func() {
            node.Controller.Init()
            reflect.ValueOf(node.Controller).MethodByName(node.Action).Call([]reflect.Value{})
        }
    }(node)

    ThroughMiddlesToController(node.Middles, function)
}

func ThroughMiddlesToController(middles []string, then func()) {
    callback := func(io *httpContext.Io) func() {
        return nil
    }
    var validMiddles []middle.Interface
    for _, alias := range middles {
        if _, ok := instance.Middles[alias]; ok {
            validMiddles = append(validMiddles, instance.Middles[alias])
        } else {
            panic("别名为：" + alias + "的中间件不存在")
        }
    }
    io := httpContext.GetIo()

    middleOk := true
    for i := len(validMiddles) - 1; i >= 0; i-- {
        func(middle middle.Interface, middleOk *bool) {
            tmpCallback := callback
            callback = func(io *httpContext.Io) func() {
                callable, ok := middle.Handle(io, tmpCallback)
                if ok {
                    return callable
                } else {
                    *middleOk = ok
                    return nil
                }
            }
        }(validMiddles[i], &middleOk)
    }
    callback(io)
    if !middleOk {
        return
    }
    then()

}
