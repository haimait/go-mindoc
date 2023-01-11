package component

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var componentConfiguration = make(map[string]interface{})

// LoadConfigurationFile 读取yaml中的内容 到componentConfiguration中
func LoadConfigurationFile(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &componentConfiguration)
	if err != nil {
		return err
	}
	return nil
}

// GetComponentConfiguration Unmarshal 到传入的 结构体中
func GetComponentConfiguration(name string, conf interface{}) error {
	if obj, ok := componentConfiguration[name]; ok {
		marshal, err := yaml.Marshal(obj)
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(marshal, conf)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("Component configuration not find")
	}
}

func LoadConfigOptionV1(instances ...Component) {
	//加载yaml中的配置项目到传一的结构体中，如DB / Redis
	LoadComponents(instances...)
}

func LoadConfigOption(configPath string, instances ...Component) {
	//读取yaml中的内容 到componentConfiguration中
	err := LoadConfigurationFile(configPath)
	if err != nil {
		fmt.Printf("ioutil.ReadFile(path) failed err:(%v)", err.Error())
		return
	}
	//加载yaml中的配置项目到传一的结构体中，如DB / Redis
	LoadComponents(instances...)
}
