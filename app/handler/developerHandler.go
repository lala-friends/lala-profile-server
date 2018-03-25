package handler

import (
	"goframework/app/cmm"
	"goframework/app/util"
	"goframework/app/domain"
	"database/sql"
	"strings"
	"github.com/go-sql-driver/mysql"
)

func HandleGetDevelopersLikeSearch(s *cmm.Server, db *sql.DB)  {
	s.HandleFunc("GET", "/developers/search/:searchKeyword", func(c *cmm.Context) {
		c.SetDefaultHeader()


		rows, err := db.Query(util.SELECT_DEVELOPERS_LIKE_SEARCH, "%"+c.Params["searchKeyword"].(string)+"%", "%"+c.Params["searchKeyword"].(string)+"%")
		util.HandleSqlErr(err)
		defer rows.Close()

		searchPerson := make([]domain.SearchPerson, 0)
		for rows.Next() {
			var id int
			var name string
			var email, introduce, imageUrl, color, blog, github, facebook, tags sql.NullString
			err := rows.Scan(&id, &name, &email, &introduce, &imageUrl, &color, &blog, &github, &facebook, &tags)
			util.HandleSqlErr(err)
			p := domain.SearchPerson{id, name, email.String, introduce.String, imageUrl.String, color.String, blog.String, github.String, facebook.String, tags.String}
			searchPerson = append(searchPerson, p)
		}
		c.RenderJson(searchPerson)
	})
}

func HandleGetDevelopers(s *cmm.Server, db *sql.DB) {
	s.HandleFunc("GET", "/developers", func(c *cmm.Context) {
		c.SetDefaultHeader()

		rows, err := db.Query(util.SELECT_PERSON_ALL)
		util.HandleSqlErr(err)
		defer rows.Close()

		persons := make([]domain.PersonProduct, 0)
		for rows.Next() {
			var personId int
			var name string
			var email, introduce, imageUrl, repColor, blog, github, facebook sql.NullString
			err := rows.Scan(&personId, &name, &email, &introduce, &imageUrl, &repColor, &blog, &github, &facebook)
			util.HandleSqlErr(err)
			productRows, err := db.Query(util.SELECT_PRODUCT_BY_PERSON, personId)
			util.HandleSqlErr(err)
			products := make([]domain.Product, 0)
			for productRows.Next() {
				var productId int
				var productName string
				var productIntroduce, tech, productImageUrl sql.NullString
				productErr := productRows.Scan(&productId, &productName, &productIntroduce, &tech, &productImageUrl)
				util.HandleSqlErr(productErr)
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
}

func HandleGetDeveloperByDeveloperName(s *cmm.Server, db *sql.DB) {
	s.HandleFunc("GET", "/developer/:username", func(c *cmm.Context) {
		c.SetDefaultHeader()

		personId := util.GetUserId(db, c.Params["username"].(string))
		var personName string
		var personEmail, personIntroduce, personImageUrl, personRepColor, personBlog, personGithub, personFacebook sql.NullString
		personErr := db.QueryRow(util.SELECT_PERSON, personId).Scan(&personName, &personEmail, &personIntroduce, &personImageUrl, &personRepColor, &personBlog, &personGithub, &personFacebook)
		util.HandleSqlErr(personErr)

		// Project 조회
		projectRows, err := db.Query(util.SELECT_PROJECTS, personId)
		util.HandleSqlErr(err)
		defer projectRows.Close()
		projects := make([]domain.Project, 0)
		for projectRows.Next() {
			var projectName string
			var periodFrom, periodTo mysql.NullTime
			var projectIntroduce, projectDescription, projectTechs, personalRole, projectLink sql.NullString
			err := projectRows.Scan(&projectName, &periodFrom, &periodTo, &projectIntroduce, &projectDescription, &projectTechs, &personalRole, &projectLink)
			util.HandleSqlErr(err)
			projectTechsArr := strings.Split(projectTechs.String, "\n")
			pjt := domain.Project{projectName, periodFrom.Time, periodTo.Time, projectIntroduce.String, projectDescription.String, projectTechsArr, personalRole.String, projectLink.String}
			projects = append(projects, pjt)
		}

		// Product 조회
		productRows, err := db.Query(util.SELECT_PRODUCT_BY_PERSON, personId)
		util.HandleSqlErr(err)
		defer projectRows.Close()
		products := make([]domain.Product, 0)
		for productRows.Next() {
			var productId int
			var productName string
			var productIntroduce, tech, productImageUrl sql.NullString
			productErr := productRows.Scan(&productId, &productName, &productIntroduce, &tech, &productImageUrl)
			util.HandleSqlErr(productErr)
			var techString string
			if tech.Valid {
				techString = tech.String
			}
			techs := strings.Split(techString, "\n")
			product := domain.Product{productId, productName, productIntroduce.String, techs, productImageUrl.String}
			products = append(products, product)
		}

		// 결과 조합
		person := domain.PersonProductProject{personId, personName, personEmail.String, personIntroduce.String, personImageUrl.String, personRepColor.String, personBlog.String, personGithub.String, personFacebook.String, projects, products}
		c.RenderJson(person)
	})
}
