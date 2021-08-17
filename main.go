package main

import (
    "gee"
)
import _ "frame/routerconf"

func main() {
    r := gee.New()
    r.Run(":9999")

}
