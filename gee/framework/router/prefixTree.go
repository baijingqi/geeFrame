package router

import (
    "gee/framework/controller"
    "strings"
)

type TreeNode struct {
    Pattern    string               // 待匹配路由，例如 /p/:lang
    Part       string               // 路由中的一部分，例如 :lang
    Children   map[string]*TreeNode // 子节点，例如 [doc, tutorial, intro]
    IsWild     bool                 // 是否精确匹配，Part 含有 : 或 * 时为true
    Controller controller.Interface
    Action     string
    Middles    []string
    Method     string
}

func (root *TreeNode) Add(method string, urlPath string, controller controller.Interface, action string, Middles []string) {
    urlPath = strings.TrimLeft(urlPath, "/")
    arr := strings.Split(urlPath, "/")

    rootCopy := root
    for _, str := range arr {
        if _, ok := rootCopy.Children[str]; !ok {
            node := TreeNode{
                Pattern:    str,
                Part:       str,
                Children:   make(map[string]*TreeNode),
                IsWild:     strings.Contains(str, ":") || strings.Contains(str, "*"),
                Controller: controller,
                Action:     action,
                Middles:    Middles,
                Method:     method,
            }
            rootCopy.Children[str] = &node
        }
        rootCopy = rootCopy.Children[str]
    }
}

func (root *TreeNode) FindInsertPos(str string, tree []TreeNode) TreeNode {
    for _, value := range tree {
        if value.Pattern == str {
            return value
        }
    }
    return TreeNode{}
}

func (root *TreeNode) FindRouter(method string, path string) (node *TreeNode, params map[string]string) {
    path = strings.TrimLeft(path, "/")
    if path == "" {
        path = "/"
    }
    arr := strings.Split(path, "/")

    for _, node := range root.Children {
        params = make(map[string]string)

        node := find(node, method, arr, 0, params)
        if node != nil {
            if len(node.Children) <= 0 {
                return node, params
            }
        }
        params = nil
    }
    return nil, make(map[string]string)
}

func find(root *TreeNode, method string, arr []string, currentHeight int, params map[string]string) *TreeNode {
    pattern := arr[currentHeight]
    if (root.Method == "") || (root.Method == method && (root.Pattern == pattern || root.IsWild)) {
        if root.IsWild {
            params[strings.TrimLeft(root.Pattern, ":")] = pattern
        }

        if currentHeight >= len(arr)-1 {
            return root
        }

        root.Part = pattern
        pattern = arr[currentHeight]

        for key, node := range root.Children {
            if ((node.Method == "") || node.Method == method) || (key == pattern || node.IsWild) {
                if node.IsWild {
                    params[strings.TrimLeft(key, ":")] = pattern
                }
                if currentHeight >= len(arr)-1 {
                    return node
                } else {
                    node := find(node, method, arr, currentHeight+1, params)
                    if node != nil {
                        return node
                    }
                }
            }
        }
    }
    return nil
}
