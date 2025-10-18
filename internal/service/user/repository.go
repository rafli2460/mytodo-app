package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/rafli024/mytodo-app/internal/contract"
	"github.com/rafli024/mytodo-app/internal/datasources"
	"github.com/rafli024/mytodo-app/internal/entities"
)

type Statement struct {
	insert        *sqlx.NamedStmt
	findById      *sqlx.Stmt
	getByUsername *sqlx.Stmt
}

type Repository struct {
	app  *contract.App
	stmt Statement
}

func NewUserRepository(app *contract.App) contract.UserRepository {
	stmts := Statement{
		insert:        datasources.PrepareNamed(app.Ds.WriterDB, "INSERT INTO users (username, password) VALUES (:username, :password)"),
		findById:      datasources.Prepare(app.Ds.ReaderDb, "SELECT * FROM users WHERE id = ?"),
		getByUsername: datasources.Prepare(app.Ds.ReaderDb, "SELECT * FROM users WHERE username = ?"),
	}

	r := Repository{
		app:  app,
		stmt: stmts,
	}

	return &r
}

func (r *Repository) Insert(user entities.User) (err error) {
	_, err = r.stmt.insert.Exec(user)
	if err != nil {
		r.app.Logger.Error().Stack().
			Str("Context", "Insert New User").
			Err(err).Msg("")
	}

	return
}

func (r *Repository) FindById(id string) (user entities.User, err error) {
	err = r.stmt.findById.Get(&user, id)
	if err != nil {
		r.app.Logger.Error().Stack().
			Str("Context", "Find user by id").
			Err(err).Msg("")
	}

	return
}

// FindByUSername implements contract.UserRepository.
func (r *Repository) FindByUsername(username string) (user entities.User, err error) {
	err = r.stmt.getByUsername.Get(&user, username)
	if err != nil {
		r.app.Logger.Error().Stack().
			Str("Context", "Find user by Name").
			Err(err).Msg("")
	}

	return
}
