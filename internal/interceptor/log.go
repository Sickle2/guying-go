package interceptor

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"guying-go/utils"
	"strconv"
	"time"
)

func Log() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		method := c.Method()
		uri := c.OriginalURL()
		utils.Logger.Info("------------------request begin------------------")
		utils.Logger.Info("------------------[" + method + "] " + uri + "------------------")
		auth := c.GetReqHeaders()["Authorization"]
		utils.Logger.Info("------------------auth: " + auth + "------------------")
		if c.Method() == fasthttp.MethodGet {
			args := c.Request().PostArgs().String()
			utils.Logger.Info("------------------args: " + args + "------------------")
		}
		c.Next()

		end := time.Now()
		duration := end.Sub(start)
		code := c.Response().Header.StatusCode()
		utils.Logger.Info("[" + method + "] " + uri + " code:" + strconv.Itoa(code) + " duration:" + duration.String())
		utils.Logger.Info("------------------request end------------------")
		return nil
	}
}
