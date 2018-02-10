package main

import (
	"fmt"
	"strings"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	//"time"
	"goframework/app/domain"
)

func main() {
	// db 접속
	db, err := sql.Open("mysql", "ryan:fkdldjs@tcp(52.79.98.34:3306)/lala_profile")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 서버 생성
	s := NewServer()

	s.HandleFunc("GET", "/", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "welcom to lala-profile!")
		//c.RenderTemplate("/public/index.html", map[string]interface{}{"time": time.Now()})
	})

	////////////////////////////////// 전체조회 //////////////////////////////////

	// 개발자 전체조회 + 프로덕트
	s.HandleFunc("GET", "/developers", func(c *Context) {
		c.SetDefaultHeader()

		rows, err := db.Query(util.SELECT_PERSON_ALL)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		persons := make([]domain.PersonProduct, 0)
		for rows.Next() {
			var personId int
			var name string
			var email, introduce, imageUrl, repColor, blog, github, facebook sql.NullString
			err := rows.Scan(&personId, &name, &email, &introduce, &imageUrl, &repColor, &blog, &github, &facebook)
			if err != nil {
				log.Fatal(err)
			}
			productRows, err := db.Query(util.SELECT_PRODUCT_BY_PERSON, personId)
			if err != nil {
				log.Fatal(err)
			}
			products := make([]domain.Product, 0)
			for productRows.Next() {
				var productId int
				var productName string
				var productIntroduce, tech, productImageUrl sql.NullString
				productErr := productRows.Scan(&productId, &productName, &productIntroduce, &tech, &productImageUrl)
				if productErr != nil {
					log.Fatal(productErr)
				}
				var techString string
				if tech.Valid {
					techString = tech.String
				}
				techs := strings.Split(techString, "\n")
				product := domain.Product{productId, productName, productIntroduce.String, techs, productImageUrl.String}
				products = append(products, product)
			}
			p := domain.PersonProduct{personId, name, email.String, introduce.String, imageUrl.String, repColor.String, blog.String, github.String, facebook.String, products}
			persons = append(persons, p)
		}
		c.RenderJson(persons)
	})

	// 프로덕트 전체조회 + 개발자
	s.HandleFunc("GET", "/products", func(c *Context) {
		c.SetDefaultHeader()

		productRows, err := db.Query(util.SELECT_PRODUCT_ALL)
		if err != nil {
			log.Fatal(err)
		}
		defer productRows.Close()

		productPersons := make([]domain.ProductPerson, 0)
		for productRows.Next() {
			var productId int
			var productName string
			var productIntroduce, productTech, productImageUrl sql.NullString
			err := productRows.Scan(&productId, &productName, &productIntroduce, &productTech, &productImageUrl)
			if err != nil {
				log.Fatal(err)
			}
			personRows, err := db.Query(util.SELECT_PERSON_BY_PRODUCT, productId)
			if err != nil {
				log.Fatal(err)
			}
			persons := make([]domain.Person, 0)
			for personRows.Next() {
				var personId int
				var personName string
				var personEmail, personIntroduce, personImageUrl, personRepColor, personBlog, personGitHub, personFacebook sql.NullString
				personoErr := personRows.Scan(&personId, &personName, &personEmail, &personIntroduce, &personImageUrl, &personRepColor, &personBlog, &personGitHub, &personFacebook)
				if personoErr != nil {
					log.Fatal(personoErr)
				}
				person := domain.Person{personId, personName, personEmail.String, personIntroduce.String, personImageUrl.String, personRepColor.String, personBlog.String, personGitHub.String, personFacebook.String}
				persons = append(persons, person)
			}
			techs := strings.Split(productTech.String, "\n")
			productPerson := domain.ProductPerson{productId, productName, productIntroduce.String, techs, productImageUrl.String, persons}
			productPersons = append(productPersons, productPerson)
		}
		c.RenderJson(productPersons)
	})

	/////////////////////////////////////// 개별정보 조회 ///////////////////////////////////////

	// product 개별 조회
	s.HandleFunc("GET", "/product/:productName", func(c *Context) {
		c.SetDefaultHeader()
		productId := util.GetProductId(db, c.Params["productName"].(string))
		var productName string
		var productIntroduce, productTech, productRepColor, productImageUrl sql.NullString
		err = db.QueryRow(util.SELECT_PRODUCT_BY_PRODUCT_ID, productId).Scan(
			&productId, &productName, &productIntroduce, &productTech, &productRepColor, &productImageUrl)
		println(err)
		// detail append
		productDetailRows, productDetailRowsErr := db.Query(util.SELECT_PRODUCT_DETAIL_BY_PRODUCT_ID, productId)
		println(productDetailRowsErr)
		productDetails := make([]domain.ProductDetail, 0)
		for productDetailRows.Next() {
			var productDetailId int
			var productDetailTitle string
			var productDetailDescription, productDetailImageUrl sql.NullString
			productDetailErr := productDetailRows.Scan(
				&productDetailId, &productDetailTitle, &productDetailDescription, &productDetailImageUrl, &productId)
			println(productDetailErr)
			productDetail := domain.ProductDetail{productDetailId, productDetailTitle, productDetailDescription.String, productDetailImageUrl.String, productId}
			productDetails = append(productDetails, productDetail)
		}

		// comment append
		commentRows, commentRowsErr := db.Query(util.SELECT_COMMENT_BY_PRODUCT_ID, productId);
		comments := make([]domain.Comment, 0)
		println(commentRows)
		println(comments)
		println(commentRowsErr)
		for commentRows.Next() {
			var commentId int
			var commentParentId sql.NullInt64
			var commentEmail, commentRegDt, commentModDt string
			var commentMessage sql.NullString
			commentErr := commentRows.Scan(
				&commentId, &commentEmail, &commentMessage, &commentParentId, &productId, &commentRegDt, &commentModDt)
			println(commentErr)
			comment := domain.Comment{commentId, commentEmail, commentMessage.String, commentParentId.Int64, productId, commentRegDt, commentModDt}
			comments = append(comments, comment)
		}
		techs := strings.Split(productTech.String, "\n")
		productProductDetailComment := domain.ProductProductDetailComment{
			productId, productName, productIntroduce.String, techs, productRepColor.String, productImageUrl.String, productDetails, comments}
		c.RenderJson(productProductDetailComment)
	})


	// 개발자 개안정보 조회
	s.HandleFunc("GET", "/developer/:username", func(c *Context) {
		c.SetDefaultHeader()

		id := util.GetUserId(db, c.Params["username"].(string))
		var name string
		var email, introduce, imageUrl, repColor, blog, github, facebook sql.NullString
		err = db.QueryRow(
			util.SELECT_PERSON, id).Scan(&name, &email, &introduce, &imageUrl, &repColor, &blog, &github, &facebook)
		util.HandleSqlErr(err)

		p := domain.Person{id, name, email.String, introduce.String, imageUrl.String, repColor.String, blog.String, github.String, facebook.String}
		c.RenderJson(p)
	})
	///////////////////////////////////// PROJECT ///////////////////////////////////////////

	s.HandleFunc("GET", "/developer/:username/projects", func(c *Context) {
		id := util.GetUserId(db, c.Params["username"].(string))
		rows, err := db.Query(util.SELECT_PROJECTS, id)
		util.HandleSqlErr(err)
		defer rows.Close()

		for rows.Next() {
			var projectName string
			var period, personalRole, mainOperator, projectSummary, responsibilities, usedTechnology, primaryRole, projectResult, linkedSite sql.NullString
			err := rows.Scan(&projectName, &period, &personalRole, &mainOperator, &projectSummary, &responsibilities, &usedTechnology, &primaryRole, &projectResult, &linkedSite)
			if err != nil {
				log.Fatal(err)
			}
			pjt := domain.Project{projectName, period.String, personalRole.String, mainOperator.String, projectSummary.String, responsibilities.String, usedTechnology.String, primaryRole.String, projectResult.String, linkedSite.String}
			c.RenderJson(pjt)
		}
	})

	s.HandleFunc("POST", "/users", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, c.Params)
	})

	s.HandleFunc("POST", "/users/:user_id/addresses", func(c *Context) {
		fmt.Println(c.ResponseWriter, c.Params)
	})

	// 웹서버 구동
	s.Run(":38001")
}