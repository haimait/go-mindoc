package db

import (
	"fmt"
	"github.com/haimait/go-mindoc/pkg/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestInstance_Load 加载yaml中的配置项到传一的结构体中
func TestInstance_LoadV2(t *testing.T) {
	err := core.LoadConfigurationFile("../../../user/rpc/etc/user.yaml")
	if err != nil {
		fmt.Printf("ioutil.ReadFile(path) failed err:(%v)", err.Error())
		return
	}
	//加载yaml中的配置项目到传一的结构体中，如DB / Redis
	core.LoadConfigOptionV1(&Instance{})
	source := Get().DataSource
	assert.Equal(t, source, "root:123456@tcp(127.0.0.1:3306)/go_mindoc?charset=utf8mb4&parseTime=true&loc=Asia/Shanghai&timeout=1000ms")
}

// TestInstance_Load 加载yaml中的配置项到传一的结构体中
func TestInstance_Load(t *testing.T) {
	//加载yaml中的配置项目到传一的结构体中，如DB / Redis
	core.LoadConfigOption("../../../user/rpc/etc/user.yaml", &Instance{})
	source := Get().DataSource
	assert.Equal(t, source, "root:123456@tcp(127.0.0.1:3306)/go_mindoc?charset=utf8mb4&parseTime=true&loc=Asia/Shanghai&timeout=1000ms")
}

//func TestInstance_Run(t *testing.T) {
//	langgo.Run(&Instance{Message: "hello"})
//	assert.Equal(t, Get().Message, "hello")
//}

//func Test_NewFromYaml(t *testing.T) {
//
//}
