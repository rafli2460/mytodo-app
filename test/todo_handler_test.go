package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/rafli024/mytodo-app/internal/contract"
	"github.com/rafli024/mytodo-app/internal/entities"
	"github.com/rafli024/mytodo-app/internal/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockTodoService is a mock implementation of the TodoService interface
type MockTodoService struct {
	mock.Mock
}

func (m *MockTodoService) GetByUserId(userID int64) ([]entities.Todo, error) {
	args := m.Called(userID)
	return args.Get(0).([]entities.Todo), args.Error(1)
}

func (m *MockTodoService) Create(todo entities.Todo) (entities.Todo, error) {
	args := m.Called(todo)
	return args.Get(0).(entities.Todo), args.Error(1)
}

func (m *MockTodoService) Update(id int64, todo entities.Todo) (entities.Todo, error) {
	args := m.Called(id, todo)
	return args.Get(0).(entities.Todo), args.Error(1)
}

func (m *MockTodoService) Delete(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTodoService) GetById(id int64) (entities.Todo, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Todo), args.Error(1)
}

func TestCreateTodoHandler(t *testing.T) {
	// Setup
	app := fiber.New()
	mockService := new(MockTodoService)
	todoHandler := handler.NewTodoHandler(&contract.App{
		Services: &contract.Service{
			Todos: mockService,
		},
	})

	app.Post("/todos", func(c *fiber.Ctx) error {
		// Mock user_id middleware
		c.Locals("user_id", int64(1))
		return todoHandler.CreateTodo(c)
	})

	description := "Test Description"
	todoRequest := map[string]interface{}{
		"title":       "Test Todo",
		"description": &description,
	}

	createdTodo := entities.Todo{
		Id:          1,
		UserID:      1,
		Title:       "Test Todo",
		Description: &description,
	}

	mockService.On("Create", mock.AnythingOfType("entities.Todo")).Return(createdTodo, nil)

	// Create request
	body, _ := json.Marshal(todoRequest)
	req, _ := http.NewRequest("POST", "/todos", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Perform request
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	mockService.AssertExpectations(t)
}

func TestUpdateTodoHandler(t *testing.T) {
	// Setup
	app := fiber.New()
	mockService := new(MockTodoService)
	todoHandler := handler.NewTodoHandler(&contract.App{
		Services: &contract.Service{
			Todos: mockService,
		},
	})

	app.Put("/todos/:id", func(c *fiber.Ctx) error {
		// Mock user_id middleware
		c.Locals("user_id", int64(1))
		return todoHandler.UpdateTodo(c)
	})

	description := "Updated Description"
	todoRequest := map[string]interface{}{
		"title":       "Updated Todo",
		"description": &description,
	}

	updatedTodo := entities.Todo{
		Id:          1,
		UserID:      1,
		Title:       "Updated Todo",
		Description: &description,
	}

	mockService.On("Update", int64(1), mock.AnythingOfType("entities.Todo")).Return(updatedTodo, nil)

	// Create request
	body, _ := json.Marshal(todoRequest)
	req, _ := http.NewRequest("PUT", "/todos/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Perform request
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	mockService.AssertExpectations(t)
}
