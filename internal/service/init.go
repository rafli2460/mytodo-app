package service

import (
	"github.com/rafli024/mytodo-app/internal/contract"
	"github.com/rafli024/mytodo-app/internal/service/todos"
	"github.com/rafli024/mytodo-app/internal/service/user"
)

func Init(app *contract.App) *contract.Service {
	srv := &contract.Service{
		User:  user.InitUserService(app),
		Todos: todos.InitTodoService(app),
	}

	app.Logger.Log().Msg("Initializing: pass")

	return srv
}
