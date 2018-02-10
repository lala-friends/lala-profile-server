package util

import (
	//"log"
	"database/sql"
	"log"
)

const CERT_FILE_PATH_LOCAL = "/Users/ryan/go/src/goframework/app/server.pem"
const KEY_FILE_PATH_LOCAL = "/Users/ryan/go/src/goframework/app/server.key"
const CERT_FILE_PATH_SERVER = "/home/muzi/goprojects/conf/server.pem"
const KEY_FILE_PATH_SERVER = "/home/muzi/goprojects/conf/server.key"

func GetUserId(db *sql.DB, username string) int {
	id := 0
	err := db.QueryRow(SELECT_PERSON_ID_BY_PERSON_NAME, username).Scan(&id)
	if err != nil && err == sql.ErrNoRows {
		return 0
	} else {
		return id
	}
}

func GetProductId(db *sql.DB, productName string) int {
	id := 0
	err := db.QueryRow(SELECT_PRODUCT_BY_PRODUCT_NAME, productName).Scan(&id)
	if err != nil && err == sql.ErrNoRows {
		return 0
	} else {
		return id
	}
}

func HandleSqlErr(err error) {
	if err != nil && err == sql.ErrNoRows {
		return
	} else {
		log.Fatal(err)
	}
}
