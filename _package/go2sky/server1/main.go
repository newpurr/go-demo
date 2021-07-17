package main

import (
	"fmt"
	"time"

	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	"github.com/gin-gonic/gin"

	v3 "github.com/SkyAPM/go2sky-plugins/gin/v3"
)

const (
	serverName = "demo-server2"
	serverPort = 8082
)

var skyAddr = "192.168.123.160:11800"

type Params struct {
	Name string
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	r := gin.Default()
	// skyAddr 是skywaling的grpc地址，默认是localhost:11800， 默认心跳检测时间是1s
	rp, err := reporter.NewGRPCReporter(skyAddr, reporter.WithCheckInterval(5*time.Second))
	panicErr(err)
	// 初始化一个 tracer，一个服务只需要一个tracer，其含义是这个服务名称
	tracer, err := go2sky.NewTracer(serverName, go2sky.WithReporter(rp))
	panicErr(err)
	// gin 使用 sky自带的middleware
	r.Use(v3.Middleware(r, tracer))

	// 自定义一个接口
	r.POST("/user/info", func(context *gin.Context) {
		// LocalSpan可以理解为本地日志的tracer，一般用户当前应用
		span, ctx, err := tracer.CreateLocalSpan(context.Request.Context())
		panicErr(err)
		// 每一个span都有一个名字去标实操作的名称！
		span.SetOperationName("UserInfo")
		// 记住重新设置一个ctx，再其次这个ctx不是gin的ctx，而是http request的ctx
		context.Request = context.Request.WithContext(ctx)

		params := new(Params)
		err = context.BindJSON(params)
		panicErr(err)
		// 记录日志信息
		span.Log(time.Now(), "[UserInfo]", fmt.Sprintf(serverName+" satrt, req : %+v", params))
		local := gin.H{
			"msg": fmt.Sprintf(serverName+" time : %s", time.Now().Format("15:04:05.9999")),
		}
		context.JSON(200, local)
		span.Log(time.Now(), "[UserInfo]", fmt.Sprintf(serverName+" end, resp : %s", local))
		// 切记最后要设置span - end，不然就是一个非闭环的
		span.End()
	})

	r.Run(fmt.Sprintf(":%d", serverPort))
}
