package main

import (
	"fmt"
	"time"
	"strings"
	"net/http"
	"crypto/hmac"
	"crypto/sha256"
	"io"
	"encoding/hex"
)

type User struct {
	Id        string
	AddressId string
}

func main() {
	// 서버 생성
	s := NewServer()

	s.HandleFunc("GET", "/login", func(c *Context) {
		// "login.html" 렌더링
		c.RenderTemplate("/public/login.html", map[string]interface{}{"message": "로그인이 필요합니다"})
	})

	s.HandleFunc("POST", "/login", func(c *Context) {
		// 로그인 정보를 확인하여 쿠키에 인증 토큰 값 기록
		if CheckLogin(c.Params["username"].(string), c.Params["password"].(string)) {
			http.SetCookie(c.ResponseWriter, &http.Cookie{
				Name:  "X_AUTH",
				Value: Sign(VerifyMessage),
				Path:  "/",
			})
			c.Redirect("/")
		}
		// id 와 password 가 맞지 않으면 다시 "/login" 페이지 렌더링
		c.RenderTemplate("/public/login.html", map[string]interface{}{"message": "id 또는 password 가 일치하지 않습니다"})
	})

	s.Use(AuthHandler)

	s.HandleFunc("GET", "/", func(c *Context) {
		//fmt.Fprintln(c.ResponseWriter, "welcom!")
		c.RenderTemplate("/public/index.html", map[string]interface{}{"time": time.Now()})
	})

	s.HandleFunc("GET", "/about", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "about")
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

	s.HandleFunc("GET", "/users/:user_id/addresses/:address_id", func(c *Context) {
		u := User{c.Params["user_id"].(string), c.Params["address_id"].(string)}
		c.RenderJson(u)
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
