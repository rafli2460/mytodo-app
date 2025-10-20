package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rafli024/mytodo-app/internal/constant"
	"github.com/rafli024/mytodo-app/internal/contract"
	"github.com/rafli024/mytodo-app/internal/entities"
	"github.com/rafli024/mytodo-app/internal/model/requests"
	"github.com/rafli024/mytodo-app/internal/model/responses"
)

type UserHandler struct {
	app *contract.App
}

func NewUserHandler(app *contract.App) *UserHandler {
	return &UserHandler{app: app}
}

// Register is a function to register a new user
// @Summary Register a new user
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body requests.User true "User"
// @Success 201 {object} responses.Response
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /v1/auth/register [post]
func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req requests.User
	if err := c.BodyParser(&req); err != nil {
		h.app.Logger.Error().Err(err).Msg("Failed to parse request body")
		return HttpError(c, responses.BadRequest(err))
	}

	// In a real app, you would add validation here for username and password strength.

	user := entities.User{
		Username: req.Username,
		Password: req.Password,
	}

	err := h.app.Services.User.Register(user)
	if err != nil {
		// This could be a duplicate username error or a server error.
		// More specific error handling could be added here.
		return HttpError(c, responses.InternalServerError(err))
	}

	return HttpSuccess(c, "User registered successfully", nil)
}

// Login is a function to login a user
// @Summary Login a user
// @Description Login a user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body requests.User true "User"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /v1/auth/login [post]
func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req requests.User
	if err := c.BodyParser(&req); err != nil {
		h.app.Logger.Error().Err(err).Msg("Failed to parse request body")
		return HttpError(c, responses.BadRequest(err))
	}

	// In a real app, you would add validation here.

	user, err := h.app.Services.User.Login(req.Username, req.Password)
	if err != nil {
		// The service returns a generic error for security, so we can pass it to the user.
		return HttpError(c, responses.UnAuthorized(err))
	}

	// --- JWT Generation ---
	// Create the claims for the token
	claims := jwt.MapClaims{
		"user_id": user.Id,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // Token expires in 3 days
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(constant.JWTSecret))
	if err != nil {
		h.app.Logger.Error().Err(err).Msg("Failed to sign JWT token")
		return HttpError(c, responses.InternalServerError(err))
	}

	// Return the token in the response
	return HttpSuccess(c, "Login successful", fiber.Map{
		"token": t,
	})
}
