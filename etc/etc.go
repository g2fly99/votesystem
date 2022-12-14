package etc

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	LevelEmergency = iota
	LevelAlert
	LevelCritical
	LevelError
	LevelWarning
	LevelNotice
	LevelInformational
	LevelDebug
)

type (
	baseConfT struct {
		LogLevel string `yaml:"loglevel"`
		Logpath  string `yaml:"logPath"`
	}

	ConfigT struct {
		Mysql mySqlConfT `yaml:"mysql"`
		Redis redisConfT `yaml:"redis"`
		Base  baseConfT  `yaml:"common"`
	}
)

var gConf ConfigT

func ConfDetail() ConfigT {
	return gConf
}

func (this ConfigT) LogPath() string {
	return this.Base.Logpath
}

func (this ConfigT) LogLevel() int {

	switch this.Base.LogLevel {
	case "DEBUG":
		return LevelDebug
	case "INFO":
		return LevelInformational
	case "Warning":
		return LevelWarning
	case "ERROR":
		return LevelError
	case "NOTICE":
		return LevelNotice
	default:
		return LevelInformational
	}
}

func LogPath() string {
	return gConf.Base.Logpath
}

func LogLevel() int {
	return gConf.LogLevel()
}

func RedisAddr() string {

	return gConf.Redis.ServerAddr()
}

func RedisPasswd() string {
	return gConf.Redis.Password()
}

// 初始化配置文件
func InitConfig(file string) error {

	bs, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(bs, &gConf)
}

// 从nacos获取数据，初始化配置
func InitConfigFromData(data []byte) error {

	return yaml.Unmarshal(data, &gConf)
}
