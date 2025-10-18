package todos

import (
	"github.com/jmoiron/sqlx"
	"github.com/rafli024/mytodo-app/internal/contract"
	"github.com/rafli024/mytodo-app/internal/datasources"
	"github.com/rafli024/mytodo-app/internal/entities"
)

type Statement struct {
	insert       *sqlx.NamedStmt
	findById     *sqlx.Stmt
	findByUserId *sqlx.Stmt
	update       *sqlx.NamedStmt
	delete       *sqlx.Stmt
}

type Repository struct {
	app  *contract.App
	stmt Statement
}

func NewTodoRepository(app *contract.App) contract.TodoRepository {
	stmts := Statement{
		insert:       datasources.PrepareNamed(app.Ds.WriterDB, insert),
		findById:     datasources.Prepare(app.Ds.ReaderDb, findById),
		findByUserId: datasources.Prepare(app.Ds.ReaderDb, findByUserId),
		update:       datasources.PrepareNamed(app.Ds.WriterDB, update),
		delete:       datasources.Prepare(app.Ds.WriterDB, delete),
	}

	r := Repository{
		app:  app,
		stmt: stmts,
	}

	return &r
}

func (r *Repository) Insert(todo entities.Todo) (entities.Todo, error) {
	result, err := r.stmt.insert.Exec(todo)
	if err != nil {
		r.app.Logger.Error().Stack().
			Str("Context", "Insert New Todo").
			Err(err).Msg("")
		return entities.Todo{}, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		r.app.Logger.Error().Stack().
			Str("Context", "Getting Last Insert ID").
			Err(err).Msg("")
		return entities.Todo{}, err
	}

	return r.FindById(lastInsertID)
}

func (r *Repository) FindById(id int64) (todo entities.Todo, err error) {
	err = r.stmt.findById.Get(&todo, id)
	if err != nil {
		r.app.Logger.Error().Stack().
			Str("Context", "Find todo by id").
			Err(err).Msg("")
	}
	return
}

func (r *Repository) FindByUserId(userId int64) (todos []entities.Todo, err error) {
	err = r.stmt.findByUserId.Select(&todos, userId)
	if err != nil {
		r.app.Logger.Error().Stack().
			Str("Context", "Find todos by user id").
			Err(err).Msg("")
	}
	return
}

func (r *Repository) Update(todo entities.Todo) (entities.Todo, error) {
	_, err := r.stmt.update.Exec(todo)
	if err != nil {
		r.app.Logger.Error().Stack().
			Str("Context", "Update Todo").
			Err(err).Msg("")
		return entities.Todo{}, err
	}
	// After an update, you might want to return the full updated object.
	// A FindById call here would do that. For now, returning the input.
	return todo, nil
}

func (r *Repository) Delete(id int64) (err error) {
	_, err = r.stmt.delete.Exec(id)
	if err != nil {
		r.app.Logger.Error().Stack().
			Str("Context", "Delete Todo").
			Err(err).Msg("")
	}
	return
}
