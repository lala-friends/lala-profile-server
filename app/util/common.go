package util

import (
	"log"
	"database/sql"
)

const CERT_FILE_PATH_LOCAL  = "/Users/ryan/go/src/goframework/app/server.pem"
const KEY_FILE_PATH_LOCAL = "/Users/ryan/go/src/goframework/app/server.key"
//const CERT_FILE_PATH_SERVER  = "/home/muzi/goprojects/conf/server.pem"
const CERT_FILE_PATH_SERVER  = "/home/muzi/.keystore/cert.pem"
const KEY_FILE_PATH_SERVER = "/home/muzi/.keystore/key.pem"

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
