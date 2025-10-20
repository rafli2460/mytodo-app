package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rafli024/mytodo-app/internal/contract"
	"github.com/rafli024/mytodo-app/internal/entities"
	"github.com/rafli024/mytodo-app/internal/model/requests"
	"github.com/rafli024/mytodo-app/internal/model/responses"
)

type TodoHandler struct {
	app *contract.App
}

func NewTodoHandler(app *contract.App) *TodoHandler {
	return &TodoHandler{app: app}
}

// GetTodos is a function to get all todos for the authenticated user
// @Summary Get all todos
// @Description Get all todos for the authenticated user
// @Tags todos
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} responses.Response
// @Failure 401 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /v1/todos [get]
func (h *TodoHandler) GetTodos(c *fiber.Ctx) error {
	// Assuming a middleware has placed the user's ID in locals.
	userID, ok := c.Locals("user_id").(int64)
	if !ok {
		h.app.Logger.Error().Msg("Failed to get user_id from context")
		return HttpError(c, responses.UnAuthorized(fmt.Errorf("unauthorized")))
	}

	todos, err := h.app.Services.Todos.GetByUserId(userID)
	if err != nil {
		// The service layer already logs the detailed error.
		// We just need to return a user-friendly error.
		return HttpError(c, responses.InternalServerError(fmt.Errorf("failed to fetch todos")))
	}

	return HttpSuccess(c, "Todos fetched successfully", todos)
}

// CreateTodo is a function to create a new todo
// @Summary Create a new todo
// @Description Create a new todo
// @Tags todos
// @Accept json
// @Produce json
// @Security Bearer
// @Param todo body requests.Todo true "Todo"
// @Success 201 {object} responses.Response
// @Failure 400 {object} responses.ErrorResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /v1/todos [post]
func (h *TodoHandler) CreateTodo(c *fiber.Ctx) error {
	// Get user_id from the JWT middleware
	userID, ok := c.Locals("user_id").(int64)
	if !ok {
		h.app.Logger.Error().Msg("Failed to get user_id from context")
		return HttpError(c, responses.UnAuthorized(fmt.Errorf("unauthorized")))
	}

	var req requests.Todo
	if err := c.BodyParser(&req); err != nil {
		h.app.Logger.Error().Err(err).Msg("Failed to parse request body")
		return HttpError(c, responses.BadRequest(err))
	}

	// Map request to entity
	todoEntity := entities.Todo{
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
		Status:      *req.Status,
	}

	createdTodo, err := h.app.Services.Todos.Create(todoEntity)
	if err != nil {
		return HttpError(c, responses.InternalServerError(fmt.Errorf("failed to create todo")))
	}

	return HttpSuccess(c, "Todo created successfully", createdTodo)
}

// UpdateTodo is a function to update a todo
// @Summary Update a todo
// @Description Update a todo
// @Tags todos
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Todo ID"
// @Param todo body requests.Todo true "Todo"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /v1/todos/{id} [put]
func (h *TodoHandler) UpdateTodo(c *fiber.Ctx) error {
	// Get user_id from the JWT middleware
	userID, ok := c.Locals("user_id").(int64)
	if !ok {
		h.app.Logger.Error().Msg("Failed to get user_id from context")
		return HttpError(c, responses.UnAuthorized(fmt.Errorf("unauthorized")))
	}

	// Get todo ID from URL parameter
	id, err := c.ParamsInt("id")
	if err != nil {
		return HttpError(c, responses.BadRequest(fmt.Errorf("invalid todo ID")))
	}

	var req requests.Todo
	if err := c.BodyParser(&req); err != nil {
		h.app.Logger.Error().Err(err).Msg("Failed to parse request body")
		return HttpError(c, responses.BadRequest(err))
	}

	// Map request to entity
	todoEntity := entities.Todo{
		Title:       req.Title,
		Description: req.Description,
		// The UserId is used for authorization in the service/repo layer
		UserID: userID,
	}

	if req.Status != nil {
		todoEntity.Status = *req.Status
	}

	updatedTodo, err := h.app.Services.Todos.Update(int64(id), todoEntity)
	if err != nil {
		return HttpError(c, responses.InternalServerError(fmt.Errorf("failed to update todo")))
	}

	return HttpSuccess(c, "Todo updated successfully", updatedTodo)
}

// DeleteTodo is a function to delete a todo
// @Summary Delete a todo
// @Description Delete a todo
// @Tags todos
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Todo ID"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /v1/todos/{id} [delete]
func (h *TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	// Get todo ID from URL parameter
	id, err := c.ParamsInt("id")
	if err != nil {
		return HttpError(c, responses.BadRequest(fmt.Errorf("invalid todo ID")))
	}

	if err := h.app.Services.Todos.Delete(int64(id)); err != nil {
		return HttpError(c, responses.InternalServerError(fmt.Errorf("failed to delete todo")))
	}

	return HttpSuccess(c, "Todo deleted successfully", nil)
}
