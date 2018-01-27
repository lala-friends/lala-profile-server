package main

import (
	"fmt"
	"net/http"

)

func main() {


	r := &router{make(map[string]map[string]http.HandlerFunc)}

	r.HandleFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "welcom!")
	})

	r.HandleFunc("GET", "/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "about")
	})

	r.HandleFunc("GET", "/users/:id", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "retrieve user")
	})

	r.HandleFunc("GET", "/users/:user_id/addresses/:address_id", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "retrieve user`s address")
	})

	r.HandleFunc("POST", "/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "create user")
	})

	r.HandleFunc("POST", "/users/:user_id/address", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "create user`s address")
	})

	// 3000 포트로 웹서버 구동
	http.ListenAndServe(":38001", r)
}
