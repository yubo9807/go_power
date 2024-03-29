package configs

import (
	"os"

	"gopkg.in/yaml.v2"
)

type ConfigType struct {
	Port           int
	Prefix         string
	LogReserveTime int    `yaml:"logReserveTime"`
	SqlSecret      string `yaml:"sqlSecret"`
	TablePrefix    string `yaml:"tablesPrefix"`

	OpenId string `yaml:"openId"`
}

var Config ConfigType

const template = `
prefix: "/permissions"  # 路由前缀
port: 8080  # 启动端口
logReserveTime: 30  # 日志保留时间(d)

sqlSecret: "user:password@tcp(0.0.0.0:3306)/database?timeout=5s"  # sql 密匙
tablesPrefix: "s_"  # 数据库表前缀，防止与业务表冲突

openId: "1hendj97f"  # 用户唯一标识
`

var (
	Table_Menu        = "menu"
	Table_Interface   = "interface"
	Table_Element     = "element"
	Table_Roles       = "roles"
	Table_Correlation = "correlation"
)

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

	Table_Menu = Config.TablePrefix + Table_Menu
	Table_Interface = Config.TablePrefix + Table_Interface
	Table_Element = Config.TablePrefix + Table_Element
	Table_Roles = Config.TablePrefix + Table_Roles
	Table_Correlation = Config.TablePrefix + Table_Correlation
}
