package response

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func NewErrorResponse(status int, msg string) error {
	return &errorResponse{
		Status:  status,
		Message: msg,
	}
}

type errorResponse struct {
	Status  int
	Message string `json:"message"`
}

func (e *errorResponse) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

func InternalServerError(c *fiber.Ctx) error {
	return c.Status(http.StatusInternalServerError).JSON(map[string]string{
		"message": "internal server error",
	})
}

func ErrorResponse(c *fiber.Ctx, err error, hook func(error)) error {
	var response *errorResponse
	if errors.As(err, &response) {
		return c.Status(response.Status).JSON(map[string]string{
			"message": response.Message,
		})
	}

	if hook != nil {
		hook(err)
	}
	return InternalServerError(c)
}
