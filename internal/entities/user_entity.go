package entities

type User struct {
	Id       int64  `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"-"`
}

func (User) TableName() string {
	return "users"
}
