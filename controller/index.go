package controller

import "github.com/kataras/iris/v12"

func IndexPage(ctx iris.Context) {
	ctx.WriteString("hello world in controller")
}

func IndexPage1(ctx iris.Context) {
	ctx.View("index.html")
}
