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
<<<<<<< HEAD
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
=======
	rows, err := db.Query(SELECT_PERSON_ID_BY_PERSON_NAME, username)
	HandleSqlErr(err)
	defer rows.Close()
	return GetIdFromRows(rows)
}

func GetProductId(db *sql.DB, productName string) int {
	rows, err := db.Query(SELECT_PRODUCT_BY_PRODUCT_NAME, productName)
	HandleSqlErr(err)
	defer rows.Close()
	return GetIdFromRows(rows)
}

func GetIdFromRows(rows *sql.Rows) int {
	var id int
	if rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
	}
	return id
>>>>>>> parent of 9aa03a4... modify bug
}

func HandleSqlErr(err error) {
	if err != nil && err == sql.ErrNoRows {
		return
	} else {
		log.Fatal(err)
	}
}
