package todolist

type User struct {
	ID       int    `json:"-" db:"id"`
	Name     string `json:"name" db:"name"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password_hash"`
}
