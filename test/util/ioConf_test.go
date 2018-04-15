package util

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
	"goframework/app/util"
)

// 로칼 conf를 읽어오면 Error가 발생하지 않는다
func Test_WhenReadLocalConf_ThenNoErr(t *testing.T)  {
	// given

	// when
	projectPath := util.GetProjectAbsPath()
	contPath := projectPath + "/conf/lala_profile_server_local.conf"
	localConf, err := os.OpenFile(contPath, os.O_RDONLY, 0777)
	// then
	assert.Nil(t, err)
	assert.NotNil(t, localConf, "localConf must not be nil")
}
