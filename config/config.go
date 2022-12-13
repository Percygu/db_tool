package config

import (
	"github.com/spf13/viper"
	"sync"
)

var (
	globalConfig GlobalConfig // 全局配置文件
	configFile   string
	once         sync.Once
)

// GlobalConfig 业务配置结构体
type GlobalConfig struct {
	LogConf     *LogConfig                `yaml:"log_conf" mapstructure:"log_conf"`         // 本地日志配置
	TcaplusConf map[string]*TcaplusConfig `yaml:"tcaplus_conf" mapstructure:"tcaplus_conf"` // tcaplus配置
}

// SelfLogConfig 日志配置
type LogConfig struct {
	LogPattern string `yaml:"log_pattern" mapstructure:"log_pattern"` // 日志输出标准，终端输出/文件输出
	LogPath    string `yaml:"log_path" mapstructure:"log_path"`       // 日志路径
	SaveDays   uint   `yaml:"save_days" mapstructure:"save_days"`     // 日志保存天数
	Level      string `yaml:"level" mapstructure:"level"`             // 日志级别
}

type TcaplusConfig struct {
	DirUrl    string `yaml:"dir_url" mapstructure:"dir_url"`
	APPID     uint64 `yaml:"app_id" mapstructure:"app_id"`
	Signature string `yaml:"signature" mapstructure:"signature"`
	ZoneID    uint32 `yaml:"zone_id" mapstructure:"zone_id"`
}

// GetGlobalConf 获取全局配置文件
func GetGlobalConf() *GlobalConfig {
	once.Do(readConf)
	return &globalConfig
}

func readConf() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	//fileNameAll := path.Base(configFile)
	//log.Debugf("fileNameAll ==== %s\n", fileNameAll)
	//filePrefix := configFile[0 : len(configFile)-len(fileNameAll)]
	//log.Debugf("filePrefix ==== %s\n", filePrefix)
	//viper.AddConfigPath(filePrefix)
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic("read log config file err:" + err.Error())
	}
	err = viper.Unmarshal(&globalConfig)
	if err != nil {
		panic("log config file unmarshal err:" + err.Error())
	}
}
