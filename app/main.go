package main

import (
	"fmt"
	"time"
)

type User struct {
	Id string
	AddressId string
}

func main() {
	// 서버 생성
	s := NewServer()

	s.HandleFunc("GET", "/", func(c *Context) {
		//fmt.Fprintln(c.ResponseWriter, "welcom!")
		c.RenderTemplate("/public/index.html", map[string]interface{}{"time" : time.Now()})
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
		u :=User{Id: c.Params["id"].(string)}
		c.RenderXml(u)
	})

	s.HandleFunc("GET", "/users/:user_id/addresses/:address_id", func(c *Context) {
		u := User{c.Params["user_id"].(string), c.Params["address_id"].(string)}
		c.RenderJson(u)
	})

	// 웹서버 구동
	s.Run(":38001")
}
