package cmm

import (
	"net/http"
	"strings"
)

type router struct {
	// 키 : http 메서드
	// 값 : URL 패턴별로 실행할 핸들러

	handlers map[string]map[string]HandlerFunc
}

func (r *router) HandleFunc(method, pattern string, h HandlerFunc) {

	// http 메서드로 등록된 맵이 있는지 확인
	m, ok := r.handlers[method]
	if !ok {
		// 동록된 맵이 없으면 새 맵을 생성
		m = make(map[string]HandlerFunc)
		r.handlers[method] = m
	}
	// http 메서드로 등록된 맵에 URL 패턴과 핸들러 함수 등록
	m[pattern] = h
}

func match(pattern, path string) (bool, map[string]string) {
	// 패턴과 패스가 정확히 일치하면 바로 true 리턴
	if pattern == path {
		return true, nil
	}

	// 패턴과 패스를 "/" 단위로 구분
	patterns := strings.Split(pattern, "/")
	paths := strings.Split(path, "/")

	// 부분 문자열 집합의 개수가 다르면 false
	if len(patterns) != len(paths) {
		return false, nil
	}

	// 패턴에 일치하는 URL 매개변수를 담기 위한 params 맵 생성
	params := make(map[string]string)

	// "/" 로 구분된 패턴/패스의 각 문자열을 하나씩 비교
	for i := 0; i < len(patterns); i++ {
		switch {
		case patterns[i] == paths[i]:
			// 패턴과 패스의 부분 문자열이 일치하면 다음루프 수행
		case len(patterns[i]) > 0 && patterns[i][0] == ':':
			params[patterns[i][1:]] = paths[i]
		default:
			// 일치하는 경우가 없으면 false 를 반환
			return false, nil
		}
	}

	// true 와 params 를 리턴
	return true, params
}

func (r *router) handler() HandlerFunc {
	return func(c *Context) {
		// http ㅁㅔ서드에 맞는 모든 handlers 를 반복하며 요청 URL 에 해당하는 handler 를 찾음
		for pattern, handler := range r.handlers[c.Request.Method] {
			if ok, params := match(pattern, c.Request.URL.Path); ok {
				for k, v := range params{
					c.Params[k] = v
				}
				handler(c)
				return
			}
		}
		// 요청 URL 에 해당하는 handler 를 찾지 못하면 NotFound 에러처리
		http.NotFound(c.ResponseWriter, c.Request)
		return
	}
}
