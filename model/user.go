package model

type User struct {
	ID        int    `db:"id" json:"id"`
	Username  string `db:"username" json:"username"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
	Password  string `db:"password" json:"password"`
	Types     string `db:"types" json:"types"`
}
