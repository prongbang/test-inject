package gdbc

import "database/sql"

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:l1ackme@/Testing")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	return db
}
