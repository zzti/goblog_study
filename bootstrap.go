package goblog

import (
	"fmt"
	"time"

	"goblog/config"
	"goblog/route"

	"github.com/kataras/iris/v12"
)

type Bootstrap struct {
	Application *iris.Application
	Port        int
	LoggerLevel string
}

func New(port int, loggerLevel string) *Bootstrap {
	var bootstrap Bootstrap
	bootstrap.Application = iris.New()
	bootstrap.Port = port
	bootstrap.LoggerLevel = loggerLevel
	return &bootstrap
}

func (bootstrap *Bootstrap) LoadRoutes() {
	route.Register(bootstrap.Application)
}

func TimestampToDate(in int64, layout string) string {
	t := time.Unix(in, 0)
	return t.Format(layout)
}

func (bootstrap *Bootstrap) Serve() {
	bootstrap.Application.Logger().SetLevel(bootstrap.LoggerLevel)
	bootstrap.Application.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hello World in root")
	})

	bootstrap.LoadRoutes()

	djangoEngine := iris.Django(".././template", ".html")

	if config.ServerConfig.Env == "development" {
		////测试环境下动态加载模板
		djangoEngine.Reload(true)
	}
	// 为引擎模板增加自定义函数
	djangoEngine.AddFunc("timestampToDate", TimestampToDate)

	bootstrap.Application.RegisterView(djangoEngine)
	bootstrap.Application.Run(
		iris.Addr(fmt.Sprintf("127.0.0.1:%d", bootstrap.Port)),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithoutBodyConsumptionOnUnmarshal,
	)
}
