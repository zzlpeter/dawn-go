package libs

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"sync"
)


type tomlConfig struct {
	Mysql	map[string]mysql 	`toml:"mysql"`
	Redis 	map[string]redis 	`toml:"redis"`
	Basic   map[string]interface{}	`toml:"basic"`
}

var (
	cfg tomlConfig
	once sync.Once
)

type mysql struct {
	Host 		string 		`json:"host" map:"host"`
	Port 		int			`json:"port" map:"port"`
	Username  	string		`json:"username" map:"username"`
	Password 	string		`json:"password" map:"password"`
	Db 			string		`json:"db" map:"db"`
}

type redis struct {
	Host 		string		`json:"host" map:"host"`
	Port		int			`json:"port" map:"port"`
	Db 			int			`json:"db" map:"db"`
	Password 	string		`json:"password" map:"password"`
}

func getFilePath() string {
	return ""
}

func readConfig() {
	_, err := toml.DecodeFile("/Users/zhangzhiliang/tiger/zzl/go/src/dawn-go/conf/conf.toml", &cfg)
	if err != nil {
		log.Fatalf("read conf.toml fails with error: %v", err.Error())
	}
	fmt.Println("read config init...")
}

func getConfig() *tomlConfig {
	once.Do(func() {
		readConfig()
	})
	return &cfg
}

type Config struct {
}

func (c Config) MysqlConf(alias string) map[string]interface{} {
	conf := getConfig()
	cS := conf.Mysql[alias]
	m := Struct2Map(cS)
	return m
}

func (c Config) MysqlConfS() map[string]map[string]interface{} {
	conf := getConfig()
	mysqlMap := make(map[string]map[string]interface{})
	for k, v := range conf.Mysql {
		m := Struct2Map(v)
		mysqlMap[k] = m
	}
	return mysqlMap
}

func (c Config) RedisConf(alias string) map[string]interface{} {
	conf := getConfig()
	cS := conf.Redis[alias]
	m := Struct2Map(cS)
	return m
}

func (c Config) RedisConfS() map[string]map[string]interface{} {
	conf := getConfig()
	redisMap := make(map[string]map[string]interface{})
	for k, v := range conf.Redis {
		m := Struct2Map(v)
		redisMap[k] = m
	}
	return redisMap
}

func (c Config) BasicConf() map[string]interface{} {
	conf := getConfig()
	return conf.Basic
}

func init() {
	getConfig()
}