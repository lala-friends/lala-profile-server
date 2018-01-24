package main

import (
	"fmt"
	"net/http"
)

func main()  {
	// "/" 경로로 접속했을 때 처리할 핸들러 함수 지정
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// "welcom!" 문자열을 화면에 출력
		fmt.Fprintln(writer, "welcom!")
	})

	// 3000 포트로 웹서버 구동
	http.ListenAndServe(":3000", nil)
}
