package main

import (
	"net/http"
	"encoding/json"
	"encoding/xml"
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