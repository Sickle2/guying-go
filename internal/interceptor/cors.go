package interceptor

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"strings"
)

func Cors() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		//请求方法
		method := ctx.Method()
		//请求头部
		origin := ctx.GetReqHeaders()["Origin"]
		//声明请求头keys
		var headerKeys []string
		for k, _ := range ctx.GetReqHeaders() {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			//这是允许访问所有域
			ctx.Request().Header.Set("Access-Control-Allow-Origin", "*")
			//服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			ctx.Request().Header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//header的类型
			ctx.Request().Header.Set("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//允许跨域设置
			ctx.Request().Header.Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			//缓存请求信息 单位为秒
			ctx.Request().Header.Set("Access-Control-Max-Age", "172800")
			//跨域请求是否需要带cookie信息 默认设置为true
			ctx.Request().Header.Set("Access-Control-Allow-Credentials", "false")
			//设置返回格式是json
			ctx.Request().Header.Set("constant-type", "application/json")
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			ctx.SendStatus(fasthttp.StatusOK)
		}
		//处理请求
		return ctx.Next()
	}
}
