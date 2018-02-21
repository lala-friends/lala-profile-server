package cmm

import (
	"time"
	"log"
	"net/http"
	"encoding/json"
	"strings"
	"path"
)

type Middleware func(next HandlerFunc) HandlerFunc

func logHandler(next HandlerFunc) HandlerFunc {
	return func(c *Context) {
		// next(c) fㅡㄹ 실항하기 전에 현재 시간을 기록
		t := time.Now()

		// 다음 핸들러 수행
		next(c)

		// 웹 요청 정보와 전체 소요시간을 로그로 남김
		log.Printf("[%s] %q %v\n", c.Request.Method, c.Request.URL.String(), time.Now().Sub(t))
	}
}

func recoverHandler(next HandlerFunc) HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(c.ResponseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next(c)
	}
}

func parseFormHandler(next HandlerFunc) HandlerFunc {
	return func(c *Context) {
		c.Request.ParseForm()
		for k, v := range c.Request.PostForm {
			if len(v) > 0 {
				c.Params[k] = v[0]
			}
		}
		next(c)
	}
}

func parseJsonBodyHandler(next HandlerFunc) HandlerFunc {
	return func(c *Context) {
		var m map[string]interface{}
		if json.NewDecoder(c.Request.Body).Decode(&m); len(m) > 0 {
			for k, v := range m {
				c.Params[k] = v
			}
		}
		next(c)
	}
}

func staticHandler(next HandlerFunc) HandlerFunc {
	var (
		dir = http.Dir(".")
		indexFile = "index.html"
	)
	return func(c *Context) {
		// http 메서드가 GET HEAD 가 아니면 다음핸들러 수행
		if c.Request.Method != "GET" && c.Request.Method != "HEAD" {
			next(c)
			return
		}

		file := c.Request.URL.Path
		// URL 경로에 해당하는 파일 열기 시도
		f, err := dir.Open(file)
		if err != nil {
			// 파일 열기에 실패하면 다음 핸들러 수행
			next(c)
			return
		}
		defer f.Close()

		fi, err := f.Stat()
		if err != nil {
			// 파일 상태가 정상이 아니면 바로 다음 핸들러 수행
			next(c)
			return
		}

		// URL 경로가 디렉터리면 indexFile을 사용
		if fi.IsDir() {
			// 디렉터리 경로를 URL 로 사용하면 경로끝에 "/" 붙여야 함
			if !strings.HasSuffix(c.Request.URL.Path, "/") {
				http.Redirect(c.ResponseWriter, c.Request, c.Request.URL.Path+"/", http.StatusFound)
				return
			}

			// 디레겉리를 가리키는 URL 경로에 indexFile 이름을 붙여서 전체 파일 경로 생성
			file := path.Join(file, indexFile)

			// indexFile 열기 시도
			f, err := dir.Open(file)
			if err != nil {
				next(c)
				return
			}
			defer f.Close()

			fi, err = f.Stat()
			if err != nil || fi.IsDir() {
				// indexFile 상태가 정상이 아니면 다음 핸들러 수행
				next(c)
				return
			}
		}

		// file 의 내용 전달(next 핸들러로 제어권을 안넘기고 처리종료)
		http.ServeContent(c.ResponseWriter, c.Request, file, fi.ModTime(), f)
	}

}