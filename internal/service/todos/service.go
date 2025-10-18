package todos

import (
	"github.com/rafli024/mytodo-app/internal/contract"
	"github.com/rafli024/mytodo-app/internal/entities"
)

type Service struct {
	app  *contract.App
	repo contract.TodoRepository
}

func InitTodoService(a *contract.App) contract.TodoService {
	r := NewTodoRepository(a)

	svc := &Service{
		app:  a,
		repo: r,
	}

	return svc
}

func (s *Service) Create(todo entities.Todo) (entities.Todo, error) {
	// Set a default status if it's not provided.
	todo.Status = "pending"
	return s.repo.Insert(todo)
}

func (s *Service) GetById(id int64) (entities.Todo, error) {
	return s.repo.FindById(id)
}

func (s *Service) GetByUserId(userId int64) ([]entities.Todo, error) {
	return s.repo.FindByUserId(userId)
}

func (s *Service) Update(id int64, todo entities.Todo) (entities.Todo, error) {
	todo.Id = id
	return s.repo.Update(todo)
}

func (s *Service) Delete(id int64) error {
	return s.repo.Delete(id)
}
