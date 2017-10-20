package model

import "database/sql"

type Connecttion struct {
	Db *sql.DB
}
