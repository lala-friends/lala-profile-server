package util

import (
	"testing"
	"fmt"
	"goframework/app/util"
)

// 프로젝트의 절대경로를 가져옴
// assert 하지 않음
func Test_WhenMakeProjectAbsPath_ThenPrintAbsPath(t *testing.T)  {
	//given

	// when
	projectPath := util.GetProjectAbsPath()

	// then
	fmt.Println(projectPath)
}
