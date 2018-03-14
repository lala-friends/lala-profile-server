package util

import (
	"log"
	"database/sql"
	"os"
	"strings"
	"time"
)

var GOPATH = ""
var LOG_PATH = ""
var LOG_FILE = ""

const APPLICATION_NAME = "lala-profile-server"
const CONST_GOPATH = "GOPATH"
const DRIVER_NAME = "mysql"
const DB_USER = "tiffany"
const DB_PASSWORD = "xlvksl"
const DB_ADDRESS = "13.125.241.114:3306"
const DB_NAME = "lala_profile"

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
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
}

func SetLogger() {
	getEnviron := os.Environ()
	for _, getEnv := range getEnviron {
		getTmp := strings.Split(getEnv, "=")
		if strings.EqualFold(CONST_GOPATH, getTmp[0]) {
			GOPATH = getTmp[1]
		}
	}
	LOG_PATH = GOPATH + "/log/"
	logDate := time.Now().Local().Format("2006-01-02")
	LOG_FILE = LOG_PATH + APPLICATION_NAME + "_" + logDate
	fpLog, err := os.OpenFile(LOG_FILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	//defer fpLog.Close()

	// 표준로거를 파일로그로 변경
	log.SetOutput(fpLog)

	// 어플리케이션 시작 로그를 남김 확인
	log.Println("lala-profile-server start!! " + time.Now().Local().Format("2006-01-02"))
}