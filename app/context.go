package main

import (
	"net/http"
	"encoding/json"
	"encoding/xml"
	"html/template"
	"path/filepath"
)

type Context struct {
	Params map[string]interface{}

	ResponseWriter http.ResponseWriter
	Request *http.Request
}

type HandlerFunc func(*Context)

func (c *Context) RenderJson(v interface{})  {
	// http status 를 status ok 로 지정
	c.ResponseWriter.WriteHeader(http.StatusOK)
	// Content-type 을 application/json 으로 지정
	c.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	c.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	c.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")

	// v 값을 json 으로 출력
	if err := json.NewEncoder(c.ResponseWriter).Encode(v); err != nil {
		// 에러 발생시 renderErr 메서드 호출
		c.RenderErr(http.StatusInternalServerError, err)
	}
}

func (c *Context) RenderXml(v interface{}) {
	// http status 를 ok 로 지정
	c.ResponseWriter.WriteHeader(http.StatusOK)
	// content-type 을 application/xml 으로 지정
	c.ResponseWriter.Header().Set("Content-Type", "application/xml; charset=utf-8")

	// v 값을 xml 으로 출력
	if err := xml.NewEncoder(c.ResponseWriter).Encode(v); err != nil {
		// 에러 발생시 renderErr 메서드 호출
		c.RenderErr(http.StatusInternalServerError, err)
	}
}

func (c *Context) RenderErr(code int, err error) {
	if err != nil {
		if code > 0 {
			// 정상적인 code를 전달하면 http status 를 해당 code 로 지정
			http.Error(c.ResponseWriter, http.StatusText(code), code)
		} else {
			// 정상적인 code 가 아니면 http status 를 StatusInternalServerError 로 지정
			defaultErr := http.StatusInternalServerError
			http.Error(c.ResponseWriter, http.StatusText(defaultErr), defaultErr)
		}
	}
}

// templates: 템플릿 객체를 보관하기 위한 map
var templates = map[string]*template.Template{}

func (c *Context) RenderTemplate(path string, v interface{}) {
	// path 에 해당하는 템플릿이 있는지 확인
	t, ok := templates[path]
	if !ok {
		// path 에 해당하는 템플릿 객체가 없으면 템플릿 객체 생성
		t = template.Must(template.ParseFiles(filepath.Join(".", path)))
		templates[path] = t
	}

	// v 값을 템플릿 내부로 전달하여 만들어진 최종 결과를 c.ResponseWriter 에 출력
	t.Execute(c.ResponseWriter, v)
}

// 리다이렉트
func (c *Context) Redirect(url string)  {
	http.Redirect(c.ResponseWriter, c.Request, url, http.StatusMovedPermanently)
}