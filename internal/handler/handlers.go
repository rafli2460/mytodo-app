package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/rafli024/mytodo-app/internal/constant"
	"github.com/rafli024/mytodo-app/internal/model/responses"
)

func HttpError(c *fiber.Ctx, err error) error {
	var errResponse *responses.ErrorResponse
	if errors.As(err, &errResponse) {
		c.Status(errResponse.Status)
	}

	if errResponse == nil {
		errResponse = &responses.ErrorResponse{}
	}

	if app.Config[constant.ServerEnv] == constant.EnvDevelopment {
		errResponse.Debug = errResponse.Error()
	}

	c.Append("Access-Control-Allow-Origin", "*")

	//app.Logger.Error().Stack().
	//	Str("Method", c.Method()).
	//	Str("Path", c.Path()).
	//	Int("Status", errResponse.Status).
	//	Err(err).Msg(errResponse.Message)

	_ = c.JSON(errResponse)

	return nil
}

func HttpSuccess(c *fiber.Ctx, message string, data interface{}) (err error) {
	response := responses.Response{}
	response.Status = fiber.StatusOK
	response.Message = message
	response.Data = data

	c.Append("Access-Control-Allow-Origin", "*")

	//app.Logger.Log().
	//	Str("Method", c.Method()).
	//	Str("Path", c.Path()).
	//	Str("Status", response.Status).
	//	Msg(response.Message)

	_ = c.JSON(response)

	return nil
}
