package route

import (
	"goblog/controller"

	"github.com/kataras/iris/v12"
)

func Register(app *iris.Application) {
	app.Get("/hello", func(ctx iris.Context) {
		ctx.WriteString("hello world in route")
	})
	app.Get("helloworld", controller.IndexPage)
	app.Get("helloworld1", controller.IndexPage1)
}
