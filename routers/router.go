package routers

import (
	"code.tlx.com/go-mvc/controllers"
	"code.tlx.com/go-mvc/controllers/person"
	"github.com/echo"
)

func RegisterRouter(engine *echo.Echo){
	//首页
	engine.GET("index",controllers.IndexHandle)
	//人
	engine.GET("person",person.PersonHandle)
}

