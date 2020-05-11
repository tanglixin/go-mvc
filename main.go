package main

import (
	"code.tlx.com/go-mvc/routers"
	"github.com/echo"
)
func main()  {
	//实例化echo对象
	engine := echo.New()

	//注册路由
	routers.RegisterRouter(engine)

	//启动http server服务器，监听9090端口
	engine.Start(":9090")
}


