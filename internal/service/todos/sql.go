package todos

const (
	insert = `INSERT INTO todos (user_id, title, description, status) 
			VALUES (:user_id, :title, :description, :status)`

	findById = `SELECT * FROM todos WHERE id = ?`

	findByUserId = `SELECT * FROM todos WHERE user_id = ?`

	update = `UPDATE todos SET title=:title, description=:description, status=:status 
			WHERE id=:id`

	delete = `DELETE FROM todos WHERE id = ?`
)
