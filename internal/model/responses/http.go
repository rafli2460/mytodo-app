package responses

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type ErrorResponse struct {
	Response
	Debug string `json:"_debug,omitempty"`
}

func BadRequest(err error) *ErrorResponse {
	return &ErrorResponse{
		Response: Response{
			Status:  fiber.ErrBadRequest.Code,
			Data:    nil,
			Message: fiber.ErrBadRequest.Message,
		},
		Debug: err.Error(),
	}
}

func InternalServerError(err error) *ErrorResponse {
	return &ErrorResponse{
		Response: Response{
			Status:  fiber.ErrInternalServerError.Code,
			Data:    nil,
			Message: fiber.ErrInternalServerError.Message,
		},
		Debug: err.Error(),
	}
}

func NotFound(err error) *ErrorResponse {
	return &ErrorResponse{
		Response: Response{
			Status:  fiber.ErrNotFound.Code,
			Data:    nil,
			Message: fiber.ErrNotFound.Message,
		},
		Debug: err.Error(),
	}
}

func UnAuthorized(err error) *ErrorResponse {
	return &ErrorResponse{
		Response: Response{
			Status:  fiber.ErrUnauthorized.Code,
			Data:    nil,
			Message: fiber.ErrUnauthorized.Message,
		},
		Debug: err.Error(),
	}
}

func (e *ErrorResponse) Error() string {
	return e.Debug
}
