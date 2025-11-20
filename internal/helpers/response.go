package helpers

import "github.com/gofiber/fiber/v2"

type Response struct{}

func NewResponse() *Response {
	return &Response{}
}

type ResponseBody struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
	Message string      `json:"message"`
}

func (r *Response) Send(c *fiber.Ctx, statusCode int, data interface{}, message string, err interface{}) error {
	status := "failed"
	if statusCode >= 200 && statusCode < 300 {
		status = "success"
	}

	body := ResponseBody{
		Code:    statusCode,
		Status:  status,
		Data:    data,
		Error:   err,
		Message: message,
	}

	return c.Status(statusCode).JSON(body)
}
