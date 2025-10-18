package entities

import "time"

type Todo struct {
	Id          int64     `db:"id" json:"id"`
	UserID      int64     `db:"user_id" json:"user_id"`
	Title       string    `db:"title" json:"title"`
	Description *string   `db:"description" json:"description,omitempty"`
	Status      string    `db:"status" json:"status"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

func (Todo) TableName() string {
	return "todos"
}