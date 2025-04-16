package conf

import (
	"github.com/bytedance/go-tagexpr/v2/validator"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"sync"
)

var (
	conf *Config
	once sync.Once
)

type Config struct {
	Server         Server         `yaml:"server"`         //服务器配置
	Authentication Authentication `yaml:"authentication"` //
	MySQL          MySQL          `yaml:"mysql"`          //MySQL配置
	Redis          Redis          `yaml:"redis"`          //Redis配置
}

type Server struct {
	LogLevel  string `yaml:"log_level"`
	LogPath   string `yaml:"log_path"`
	ErrorLog  string `yaml:"error_log"`
	AccessLog string `yaml:"access_log"`
	DebugLog  string `yaml:"debug_log"`
	Addr      string `yaml:"addr"`
}

type MySQL struct {
	DSN string `yaml:"dsn"` // MySQL dsn配置
}

type Redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
	MaxIdle  int    `yaml:"max_idle"`
}

type Authentication struct {
	MaxAge        int    `yaml:"max_age"`
	AuthSecret    string `yaml:"auth_secret"` // session key
	EnableSession bool   `yaml:"enable_session"`
}

// GetConf gets configuration instance
func GetConf() *Config {
	once.Do(initConf)
	return conf
}

func initConf() {
	prefix := "conf"
	confFileRelPath := filepath.Join(prefix, "config.yaml")
	content, err := os.ReadFile(confFileRelPath)
	if err != nil {
		panic(err)
	}

	conf = new(Config)
	err = yaml.Unmarshal(content, conf)
	if err != nil {
		hlog.Error("parse yaml error - %v", err)
		panic(err)
	}
	if err := validator.Validate(conf); err != nil {
		hlog.Error("validate config error - %v", err)
		panic(err)
	}
}
