package handler

import (
	"goframework/app/cmm"
	"database/sql"
	"goframework/app/util"
	"goframework/app/domain"
	"strings"
)

func HandleGetProducts(s *cmm.Server, db *sql.DB) {
	s.HandleFunc("GET", "/products", func(c *cmm.Context) {
		c.SetDefaultHeader()

		productRows, err := db.Query(util.SELECT_PRODUCT_ALL)
		util.HandleSqlErr(err)
		defer productRows.Close()

		productPersons := make([]domain.ProductPerson, 0)
		for productRows.Next() {
			var productId int
			var productName string
			var productIntroduce, productTech, productImageUrl sql.NullString
			err := productRows.Scan(&productId, &productName, &productIntroduce, &productTech, &productImageUrl)
			util.HandleSqlErr(err)
			personRows, err := db.Query(util.SELECT_PERSON_BY_PRODUCT, productId)
			util.HandleSqlErr(err)
			persons := make([]domain.Person, 0)
			for personRows.Next() {
				var personId int
				var personName string
				var personEmail, personIntroduce, personImageUrl, personRepColor, personBlog, personGitHub, personFacebook sql.NullString
				personErr := personRows.Scan(&personId, &personName, &personEmail, &personIntroduce, &personImageUrl, &personRepColor, &personBlog, &personGitHub, &personFacebook)
				util.HandleSqlErr(personErr)
				person := domain.Person{personId, personName, personEmail.String, personIntroduce.String, personImageUrl.String, personRepColor.String, personBlog.String, personGitHub.String, personFacebook.String}
				persons = append(persons, person)
			}
			techs := strings.Split(productTech.String, "\n")
			productPerson := domain.ProductPerson{productId, productName, productIntroduce.String, techs, productImageUrl.String, persons}
			productPersons = append(productPersons, productPerson)
		}
		c.RenderJson(productPersons)
	})
}

func HandleGetProductByProductName(s *cmm.Server, db *sql.DB) {
	s.HandleFunc("GET", "/product/:productName", func(c *cmm.Context) {
		c.SetDefaultHeader()
		productId := util.GetProductId(db, c.Params["productName"].(string))
		var productName string
		var productIntroduce, productTech, productRepColor, productImageUrl sql.NullString
		err := db.QueryRow(util.SELECT_PRODUCT_BY_PRODUCT_ID, productId).Scan(
			&productId, &productName, &productIntroduce, &productTech, &productRepColor, &productImageUrl)
		util.HandleSqlErr(err)
		// detail append
		productDetailRows, productDetailRowsErr := db.Query(util.SELECT_PRODUCT_DETAIL_BY_PRODUCT_ID, productId)
		util.HandleSqlErr(productDetailRowsErr)
		productDetails := make([]domain.ProductDetail, 0)
		for productDetailRows.Next() {
			var productDetailId int
			var productDetailTitle string
			var productDetailDescription, productDetailImageUrl sql.NullString
			productDetailErr := productDetailRows.Scan(
				&productDetailId, &productDetailTitle, &productDetailDescription, &productDetailImageUrl, &productId)
			util.HandleSqlErr(productDetailErr)
			productDetail := domain.ProductDetail{productDetailId, productDetailTitle, productDetailDescription.String, productDetailImageUrl.String, productId}
			productDetails = append(productDetails, productDetail)
		}

		// comment append
		commentRows, commentRowsErr := db.Query(util.SELECT_COMMENT_BY_PRODUCT_ID, productId);
		util.HandleSqlErr(commentRowsErr)
		comments := make([]domain.Comment, 0)
		for commentRows.Next() {
			var commentId int
			var commentParentId sql.NullInt64
			var commentEmail, commentRegDt, commentModDt string
			var commentMessage sql.NullString
			commentErr := commentRows.Scan(
				&commentId, &commentEmail, &commentMessage, &commentParentId, &productId, &commentRegDt, &commentModDt)
			util.HandleSqlErr(commentErr)
			comment := domain.Comment{commentId, commentEmail, commentMessage.String, commentParentId.Int64, productId, commentRegDt, commentModDt}
			comments = append(comments, comment)
		}
		techs := strings.Split(productTech.String, "\n")
		productProductDetailComment := domain.ProductProductDetailComment{
			productId, productName, productIntroduce.String, techs, productRepColor.String, productImageUrl.String, productDetails, comments}
		c.RenderJson(productProductDetailComment)
	})
}
