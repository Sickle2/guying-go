package e

import (
	"github.com/gofiber/fiber/v2"
)

/**
 * @program: guying-go
 * @description:
 * @author: sickle
 * @create: 2023-07-28 14:29
 **/
type Meta struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message,omitempty"`
}

type Response struct {
	Meta *Meta       `json:"meta,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func Success(ctx *fiber.Ctx) {
	meta := &Meta{
		Code:    Ok,
		Message: GetMessage(Ok),
	}
	resp := &Response{
		Meta: meta,
	}
	ctx.JSON(resp)
}

func SuccessWitchData(ctx *fiber.Ctx, data interface{}) {
	meta := &Meta{
		Code:    Ok,
		Message: GetMessage(Ok),
	}
	resp := &Response{
		Meta: meta,
		Data: data,
	}
	ctx.JSON(resp)
}

func Error(ctx *fiber.Ctx, code int) {
	meta := &Meta{
		Code:    code,
		Message: GetMessage(code),
	}
	resp := &Response{
		Meta: meta,
	}
	ctx.JSON(resp)
}

func ErrorWithDate(ctx *fiber.Ctx, code int, data interface{}) {
	meta := &Meta{
		Code:    code,
		Message: GetMessage(code),
	}
	resp := &Response{
		Meta: meta,
		Data: data,
	}
	ctx.JSON(resp)
}
