package util

import (
	"log"
	"database/sql"
)

func GetUserId(db *sql.DB, username string) int {
	var id int
	rows, err := db.Query("SELECT ID FROM PERSON WHERE NAME = ?", username)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
	}
	return id
}
