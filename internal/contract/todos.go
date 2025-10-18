package contract

import "github.com/rafli024/mytodo-app/internal/entities"

type TodoRepository interface {
	Insert(todo entities.Todo) (entities.Todo, error)
	FindById(id int64) (entities.Todo, error)
	FindByUserId(userId int64) ([]entities.Todo, error)
	Update(todo entities.Todo) (entities.Todo, error)
	Delete(id int64) error
}

type TodoService interface {
	Create(todo entities.Todo) (entities.Todo, error)
	GetById(id int64) (entities.Todo, error)
	GetByUserId(userId int64) ([]entities.Todo, error)
	Update(id int64, todo entities.Todo) (entities.Todo, error)
	Delete(id int64) error
}
