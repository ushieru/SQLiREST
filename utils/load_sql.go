package utils

import (
	"database/sql"
	"os"
)

func LoadSQL(file string, db *sql.DB) {
	if file == "" {
		return
	}
	data, err := os.ReadFile(file)
	if err != nil {
		return
	}
	db.Exec(string(data))
}
