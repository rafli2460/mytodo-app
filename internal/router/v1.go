package router

import (
	"github.com/rafli024/mytodo-app/internal/contract"
	"github.com/rafli024/mytodo-app/internal/handler"
	"github.com/rafli024/mytodo-app/internal/middlewares"
)

func InitRoutes(app *contract.App) {
	// Initialize handlers
	todoHandler := handler.NewTodoHandler(app)
	userHandler := handler.NewUserHandler(app)

	// --- API Routes ---
	v1 := app.Fiber.Group("/v1")

	// --- Public Routes ---
	// Group for routes that don't require authentication
	auth := v1.Group("/auth")
	auth.Post("/register", userHandler.Register)
	auth.Post("/login", userHandler.Login)

	// --- Protected Routes ---
	// Group for routes that require authentication.
	// The JwtAuth middleware will validate the token and extract the user_id.
	todos := v1.Group("/todos", middlewares.JwtAuth())
	todos.Get("/", todoHandler.GetTodos)
	todos.Post("/", todoHandler.CreateTodo)
	todos.Put("/:id", todoHandler.UpdateTodo)
	todos.Delete("/:id", todoHandler.DeleteTodo)
}
