package interceptor

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	conf "guying-go/configs"
	"strings"
)

/**
 * @program: guying-go
 * @description: 登录
 * @author: sickle
 * @create: 2023-07-27 10:17
 **/

func Auth() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(conf.Conf.App.Secret)},
		Filter: func(ctx *fiber.Ctx) bool {
			for _, k := range conf.Conf.App.UnAuthUrl {
				if strings.Contains(ctx.Path(), k) {
					return true
				}
			}
			return false
		},
	})
}
