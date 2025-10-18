package contract

import "github.com/rafli024/mytodo-app/internal/entities"

type UserRepository interface {
	Insert(user entities.User) (err error)
	FindById(id string) (user entities.User, err error)
	FindByUsername(username string) (user entities.User, err error)
}

type UserService interface {
	Register(user entities.User) (err error)
	GetById(id string) (user entities.User, err error)
	GetByUsername(username string) (user entities.User, err error)
	Login(username string, password string) (user entities.User, err error)
}
