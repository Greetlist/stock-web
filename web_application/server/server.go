package server

import (
    "github.com/beego/beego/v2/server/web"
)

func RunServer(bindAddr string, bindPort int) {
    StaticDir["/static"] = "static"
    web.Run()
}
