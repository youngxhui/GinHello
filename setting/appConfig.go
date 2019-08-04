package setting

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type appConfig struct {
	Database database `yaml:"database"`
	Jwt      jwt      `yaml:"jwt"`
}

type database struct {
	Type string `yaml:"type"`
	Url  string `yaml:"url"`
}

type jwt struct {
	Secret  string `yaml:"secret"`
	Expires int64  `yaml:"expires"`
}

var (
	AppConfig appConfig
)

func init() {
	data, e := ioutil.ReadFile("./config/app.yml")
	if e != nil {
		log.Panicln("读取文件发生错误", e.Error())
	}
	if err := yaml.Unmarshal(data, &AppConfig); err != nil {
		log.Panicln("数据解析错误，无法解析数据", err.Error())
	}
}
