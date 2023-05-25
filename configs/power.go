package configs

import (
	"os"

	"gopkg.in/yaml.v2"
)

type ConfigType struct {
	Port      int
	Prefix    string
	SqlSecret string
}

var Config ConfigType

const template = `
prefix: "/permissions"  # 路由前缀
port: 8080  # 启动端口
sqlSecret: "root:password@tcp(0.0.0.0:3306)/permissions"  # sql 密匙
`

func init() {
	configFile := "./config.yml"
	data, err := os.ReadFile(configFile)
	if err != nil {
		os.Create(configFile)
		os.WriteFile(configFile, []byte(template), 0777)
		data, _ = os.ReadFile(configFile)
	}

	if err := yaml.Unmarshal([]byte(data), &Config); err != nil {
		panic(err)
	}
}
