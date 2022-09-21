package controller

import (
	"time"

	"github.com/kataras/iris/v12"
)

func IndexPage(ctx iris.Context) {
	ctx.WriteString("hello world in controller")
}

func IndexPage1(ctx iris.Context) {
	nowstamptime := time.Now().Unix()
	ctx.ViewData("nowstamptime", nowstamptime)
	ctx.View("index.html")
}
