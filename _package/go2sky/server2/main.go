package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	"github.com/gin-gonic/gin"

	v3 "github.com/SkyAPM/go2sky-plugins/gin/v3"
	agentv3 "skywalking.apache.org/repo/goapi/collect/language/agent/v3"
)

const (
	serverName       = "demo-server1"
	serverPort       = 8081
	remoteServerName = "demo-server2"
	remoteServerAddr = "localhost:8082"
	remotePath       = "/user/info"
)

var skyAddr = "192.168.123.160:11800"

func panicErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

type Params struct {
	Name string
}

var tracer *go2sky.Tracer

func skyMiddleware(r *gin.Engine) {
	var err error
	rp, err := reporter.NewGRPCReporter(skyAddr, reporter.WithCheckInterval(5*time.Second))
	panicErr(err)
	tracer, err = go2sky.NewTracer(serverName, go2sky.WithReporter(rp))
	panicErr(err)
	r.Use(v3.Middleware(r, tracer))
}

func trace(context *gin.Context) {
	span, ctx, err := tracer.CreateLocalSpan(context.Request.Context())
	panicErr(err)
	span.SetOperationName("Trace")

	context.Request = context.Request.WithContext(ctx)
	span.Log(time.Now(), "[Trace]", fmt.Sprintf(serverName+" satrt, params : %s", time.Now().Format("15:04:05.9999")))

	result := make([]map[string]interface{}, 0)

	//1、请求一次
	{
		url := fmt.Sprintf("http://%s%s", remoteServerAddr, remotePath)

		params := Params{
			Name: serverName + time.Now().Format("15:04:05.9999"),
		}
		buffer := &bytes.Buffer{}
		_ = json.NewEncoder(buffer).Encode(params)
		req, err := http.NewRequest(http.MethodPost, url, buffer)
		panicErr(err)

		// op_name 是每一个操作的名称
		reqSpan, err := tracer.CreateExitSpan(context.Request.Context(), "invoke - "+remoteServerName, "localhost:8082/user/info", func(headerKey, headerValue string) error {
			req.Header.Set(headerKey, headerValue)
			return nil
		})
		panicErr(err)
		reqSpan.SetComponent(2)
		reqSpan.SetSpanLayer(agentv3.SpanLayer_RPCFramework) // rpc 调用

		resp, err := http.DefaultClient.Do(req)
		panicErr(err)
		defer resp.Body.Close()

		reqSpan.Log(time.Now(), "[HttpRequest]", fmt.Sprintf("开始请求,请求服务:%s,请求地址:%s,请求参数:%+v", remoteServerName, url, params))
		body, err := ioutil.ReadAll(resp.Body)
		panicErr(err)
		fmt.Printf("接受到消息： %s\n", body)
		reqSpan.Tag(go2sky.TagHTTPMethod, http.MethodPost)
		reqSpan.Tag(go2sky.TagURL, url)
		reqSpan.Log(time.Now(), "[HttpRequest]", fmt.Sprintf("结束请求,响应结果: %s", body))
		reqSpan.End()
		res := map[string]interface{}{}
		err = json.Unmarshal(body, &res)
		panicErr(err)
		result = append(result, res)
	}

	//2 、再请求一次
	{
		url := fmt.Sprintf("http://%s%s", remoteServerAddr, remotePath)

		params := Params{
			Name: serverName + time.Now().Format("15:04:05.9999"),
		}
		buffer := &bytes.Buffer{}
		_ = json.NewEncoder(buffer).Encode(params)
		req, err := http.NewRequest(http.MethodPost, url, buffer)
		panicErr(err)

		// 出去必须用这个携带header
		reqSpan, err := tracer.CreateExitSpan(context.Request.Context(), "invoke - "+remoteServerName, "localhost:8082/user/info", func(headerKey, headerValue string) error {
			req.Header.Set(headerKey, headerValue)
			return nil
		})
		panicErr(err)
		reqSpan.SetComponent(2)
		reqSpan.SetSpanLayer(agentv3.SpanLayer_RPCFramework) // rpc 调用

		resp, err := http.DefaultClient.Do(req)
		panicErr(err)
		defer resp.Body.Close()

		reqSpan.Log(time.Now(), "[HttpRequest]", fmt.Sprintf("开始请求,请求服务:%s,请求地址:%s,请求参数:%+v", remoteServerName, url, params))
		body, err := ioutil.ReadAll(resp.Body)
		panicErr(err)
		fmt.Printf("接受到消息： %s\n", body)

		reqSpan.Tag(go2sky.TagHTTPMethod, http.MethodPost)
		reqSpan.Tag(go2sky.TagURL, url)
		reqSpan.Log(time.Now(), "[HttpRequest]", fmt.Sprintf("结束请求,响应结果: %s", body))
		reqSpan.End()
		res := map[string]interface{}{}
		err = json.Unmarshal(body, &res)
		panicErr(err)
		result = append(result, res)
	}

	// 设置响应结果
	local := gin.H{
		"msg": result,
	}
	context.JSON(200, local)
	span.Log(time.Now(), "[Trace]", fmt.Sprintf(serverName+" end, resp : %s", local))
	span.End()
	{
		span, ctx, err := tracer.CreateEntrySpan(context.Request.Context(), "Send", func(s string) (string, error) {
			return "", nil
		})
		context.Request = context.Request.WithContext(ctx)
		panicErr(err)
		span.SetOperationName("Send")
		span.Log(time.Now(), "[Info]", "send resp")
		span.End()
	}
}

func main() {

	// 这些都一样
	r := gin.Default()
	// 使用go2sky gin中间件
	skyMiddleware(r)

	// 调用接口
	r.GET("/trace", trace)

	r.Run(fmt.Sprintf(":%d", serverPort))
}
