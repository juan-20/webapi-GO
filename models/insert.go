package models

import (
	"github.com/juan-20/webapi-GO/db"
)

func Insert(todo Todo) (id int64, err error) {
	// abrir conexão
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	// o SQL
	sql := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`

	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	return
}
