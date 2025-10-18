package middlewares

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rafli024/mytodo-app/internal/constant"
	"github.com/rafli024/mytodo-app/internal/model/responses"
)

func JwtAuth() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(constant.JWTSecret)},
		SuccessHandler: jwtSuccess,
		ErrorHandler: jwtError,
	})
}

func jwtSuccess(c *fiber.Ctx) error {
	// Get the user from the context and extract the user_id
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := int64(claims["user_id"].(float64))

	// Set the user_id in the context for the next handlers
	c.Locals("user_id", userId)

	// Continue to the next middleware or handler
	return c.Next()
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or Malformed JWT" {
		return responses.BadRequest(err)
	} else {
		return responses.UnAuthorized(err)
	}
}
