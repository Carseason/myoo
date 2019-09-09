package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

/**********************************
			配置文件读取
**********************************/
var config *Config

type Config struct {
	Port         string
	Runmode      string
	ServerRender bool
	Db           struct {
		Uname string
		Pass  string
		Table string
		Addr  string
	}
	Redis struct {
		Addr string
		Pass string
	}
	Key string
}

func GetConfig() *Config {
	return config
}

func NewConfig() *Config {
	config = new(Config)
	config.openFile()
	return config
}

func (*Config) openFile() {
	pathFile, err := filepath.Abs(filepath.Dir("./"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}
	buf, err := ioutil.ReadFile(pathFile + "/conf/config.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}
	if err := json.Unmarshal(buf, &config); err != nil {
		fmt.Fprintf(os.Stderr, "Config is json an: %s\n", err)
	}
}
func (*Config) GetPort() string {
	return config.Port
}
func (*Config) GetRunmode() string {
	return config.Runmode
}
func (*Config) GetDbAddr() string {
	var addr string
	if config.Db.Addr != "" {
		addr = "tcp(" + config.Db.Addr + ")"
	}
	return config.Db.Uname + ":" + config.Db.Pass + "@" + addr + "/" + config.Db.Table + "?charset=utf8&parseTime=true&loc=Local"
}
func (*Config) GetRedisAddr() string {
	return config.Redis.Addr
}
func (*Config) GetRedisPass() string {
	return config.Redis.Pass
}
func (*Config) GetKey() string {
	return config.Key
}

func (*Config) GetServerRender() bool {
	return config.ServerRender
}
