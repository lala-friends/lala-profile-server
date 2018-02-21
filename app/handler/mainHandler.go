package handler

import (
	"goframework/app/cmm"
	"fmt"
)

func HandleMainPage(s *cmm.Server) {
	s.HandleFunc("GET", "/", func(c *cmm.Context) {
		fmt.Fprintln(c.ResponseWriter, "welcom to lala-profile!")
		//c.RenderTemplate("/public/index.html", map[string]interface{}{"time": time.Now()})
	})
}