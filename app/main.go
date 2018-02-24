package main

import (
	"database/sql"
	"goframework/app/util"
	"goframework/app/persistent"
	"goframework/app/cmm"
	"goframework/app/handler"
)

var (
	db *sql.DB	// db var
	s *cmm.Server	// server var
)

func main() {

	// 로그 세팅
	util.SetLogger()
	// db 접속
	db := persistent.GetConnection(util.DRIVER_NAME, util.DB_ADDRESS, util.DB_NAME, util.DB_USER, util.DB_PASSWORD)

	// 서버 생성
	s := cmm.NewServer()

	// main Page
	handler.HandleMainPage(s)

	////////////////////////////////// 전체조회 //////////////////////////////////

	// 개발자 전체조회 + 프로덕트
	handler.HandleGetDevelopers(s, db)
	// 프로덕트 전체조회 + 개발자
	handler.HandleGetProducts(s, db)

	/////////////////////////////////////// 개별정보 조회 /////////////////////////

	// product 개별 조회
	handler.HandleGetProductByProductName(s, db)
	// 개발자 개안정보 조회 + Projects, Products
	handler.HandleGetDeveloperByDeveloperName(s, db)

	///////////////////////////////////// PROJECT /////////////////////////////

	// 개발자별 프로젝트 전체 조회
	handler.HandleGetProjects(s, db)


	///////////////////////////////////// 삽입 /////////////////////////////////

	// 프로덕트 삽입(프로덕트 + detail)
	handler.HandleAddProductAndProductDetail(s, db)

	///////////////////////////////////////////////////////////////////////////

	// 웹서버 구동
	s.Run(":38001")
}
