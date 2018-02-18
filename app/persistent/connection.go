package persistent

import (
	"database/sql"
	"log"
)

func GetConnection(driverName, address, dbName, id, password string) *sql.DB {
	db, err := sql.Open(driverName, id+":"+password+"@tcp("+address+")/"+dbName)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("[DB] Connection Success!!")
	}
	//defer db.Close()

	return db
}
