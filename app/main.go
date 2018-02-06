package main

import (
	"fmt"
	"strings"
	"net/http"
	"crypto/hmac"
	"crypto/sha256"
	"io"
	"encoding/hex"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	//"time"
	"goframework/app/domain"
	"goframework/app/util"
)

type User struct {
	Id        string
	AddressId string
}

func main() {
	// db 접속
	db, err := sql.Open("mysql", "ryan:fkdldjs@tcp(52.79.98.34:3306)/lala_profile")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 서버 생성
	s := NewServer()

	//s.HandleFunc("GET", "/login", func(c *Context) {
	//	// "login.html" 렌더링
	//	c.RenderTemplate("/public/login.html", map[string]interface{}{"message": "로그인이 필요합니다"})
	//})

	//s.HandleFunc("POST", "/login", func(c *Context) {
	//	// 로그인 정보를 확인하여 쿠키에 인증 토큰 값 기록
	//	if CheckLogin(c.Params["username"].(string), c.Params["password"].(string)) {
	//		http.SetCookie(c.ResponseWriter, &http.Cookie{
	//			Name:  "X_AUTH",
	//			Value: Sign(VerifyMessage),
	//			Path:  "/",
	//		})
	//		c.Redirect("/")
	//	}
	//	// id 와 password 가 맞지 않으면 다시 "/login" 페이지 렌더링
	//	c.RenderTemplate("/public/login.html", map[string]interface{}{"message": "id 또는 password 가 일치하지 않습니다"})
	//})

	//s.Use(AuthHandler)

	s.HandleFunc("GET", "/", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "welcom to lala-profile!")
		//c.RenderTemplate("/public/index.html", map[string]interface{}{"time": time.Now()})
	})

	// 개발자 전체조회 + 프로덕트
	s.HandleFunc("GET", "/developers", func(c *Context) {
		c.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")

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
				product := domain.Product{productId, productName, productIntroduce, techs, productImageUrl}
				products = append(products, product)
			}
			p := domain.PersonProduct{personId, name, email, introduce, imageUrl, repColor, blog, github, facebook, products}
			persons = append(persons, p)
		}
		c.RenderJson(persons)
	})


	// 개발자 개안정보 조회
	s.HandleFunc("GET", "/developer/:username", func(c *Context) {
		c.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")

		id := util.GetUserId(db, c.Params["username"].(string))
		var name string
		var email, introduce, imageUrl, repColor, blog, github, facebook sql.NullString
		err = db.QueryRow(
			util.SELECT_PERSON, id).Scan(&name, &email, &introduce, &imageUrl, &repColor, &blog, &github, &facebook)
		if err != nil {
			log.Fatal(err)
		}
		p := domain.Person{id, name, email, introduce, imageUrl, repColor, blog, github, facebook}
		c.RenderJson(p)
	})

	// 프로덕트 전체조회 + 개발자
	s.HandleFunc("GET", "/products", func(c *Context) {
		c.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	})

	///////////////////////////////////// PROJECT ///////////////////////////////////////////

	s.HandleFunc("GET", "/profile/:username/projects", func(c *Context) {
		id := util.GetUserId(db, c.Params["username"].(string))
		rows, err := db.Query(util.SELECT_PROJECTS, id)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var projectName string
			var period, personalRole, mainOperator, projectSummary, responsibilities, usedTechnology, primaryRole, projectResult, linkedSite sql.NullString
			err := rows.Scan(&projectName, &period, &personalRole, &mainOperator, &projectSummary, &responsibilities, &usedTechnology, &primaryRole, &projectResult, &linkedSite)
			if err != nil {
				log.Fatal(err)
			}
			pjt := domain.Project{projectName, period, personalRole, mainOperator, projectSummary, responsibilities, usedTechnology, primaryRole, projectResult, linkedSite}
			c.RenderJson(pjt)
		}
	})

	s.HandleFunc("GET", "/users/:user_id/addresses/:address_id", func(c *Context) {
		u := User{c.Params["user_id"].(string), c.Params["address_id"].(string)}
		c.RenderJson(u)
	})

	//s.HandleFunc("GET", "/users/:id", func(c *Context) {
	//	if c.Params["id"] == "0" {
	//		panic("id is zero")
	//	}
	//	fmt.Fprintf(c.ResponseWriter, "retrieve user %v\n", c.Params["id"])
	//})

	//s.HandleFunc("GET", "/users/:user_id/addresses/:address_id", func(c *Context) {
	//	fmt.Fprintf(c.ResponseWriter, "retrieve user %v's adddress %v\n", c.Params["user_id"], c.Params["address_id"])
	//})

	s.HandleFunc("POST", "/users", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, c.Params)
	})

	s.HandleFunc("POST", "/users/:user_id/addresses", func(c *Context) {
		fmt.Println(c.ResponseWriter, c.Params)
	})

	s.HandleFunc("GET", "/users/:id", func(c *Context) {
		u := User{Id: c.Params["id"].(string)}
		c.RenderXml(u)
	})

	// 웹서버 구동
	s.Run(":38001")
}

const VerifyMessage = "verified"

func AuthHandler(next HandlerFunc) HandlerFunc {
	ignore := []string{"/login", "public/index.html"}
	return func(c *Context) {
		// URL prifix 가 "/login", "/public/index.html" 이면 auth 를 체크하지 않음
		for _, s := range ignore {
			if strings.HasPrefix(c.Request.URL.Path, s) {
				next(c)
				return
			}
		}

		if v, err := c.Request.Cookie("X_AUTH"); err == http.ErrNoCookie {
			// "X_AUTH" 쿠키 값이 없으면 "/login" 으로 이동
			c.Redirect("/login")
			return
		} else if err != nil {
			// 예외 처리
			c.RenderErr(http.StatusInternalServerError, err)
			return
		} else if Verify(VerifyMessage, v.Value) {
			// 쿠키 값으로 인증이 확인되면 다음 핸들러로 넘어감
			next(c)
			return
		}

		// "/login" 으로 이동
		c.Redirect("/login")
	}
}

// 인증 토큰 확인
func Verify(message, sig string) bool {
	return hmac.Equal([]byte(sig), []byte(Sign(message)))
}

// 로그인 처리
func CheckLogin(username, password string) bool {
	const (
		USERNAME = "lala_dev"
		PASSWORD = "fkfkvmfpswm"
	)
	return username == USERNAME && password == PASSWORD
}

// 인증 토큰 생성
func Sign(message string) string {
	secretKey := []byte("lala-friends-secret-key2018")
	if len(secretKey) == 0 {
		return ""
	}
	mac := hmac.New(sha256.New, secretKey)
	io.WriteString(mac, message)
	return hex.EncodeToString(mac.Sum(nil))
}
