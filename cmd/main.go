package main

import (
	"github.com/gofiber/fiber/v2"
	//_ "github.com/golang-jwt/jwt/v5"
	conf "guying-go/configs"
	"guying-go/internal/interceptor"
	"guying-go/internal/route"
	"strconv"
)

func main() {
	app := fiber.New()
	app.Use(interceptor.Log())
	app.Use(interceptor.Cors())
	route.InitRouter(app)
	var port = ":" + strconv.FormatInt(conf.Conf.App.Port, 10)
	app.Listen(port)

}
