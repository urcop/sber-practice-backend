package render

import (
	"github.com/gofiber/fiber/v2"
)

type ResponseError struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

type Message struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func SendError(ctx *fiber.Ctx, statusCode int, err error, message string) error {

	return ctx.Status(statusCode).JSON(ResponseError{
		Error:      err.Error(),
		Message:    message,
		StatusCode: statusCode,
	})

}

func SendMessage(ctx *fiber.Ctx, statusCode int, message string) error {
	return ctx.Status(statusCode).JSON(Message{
		StatusCode: statusCode,
		Message:    message,
	})
}
