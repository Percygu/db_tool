package db_tool

import (
	"flag"
	"runtime"

	"db_tool/config"
)

// Init 初始化配置
func Init() {
	config.InitConfig() // 初始化配置文件
}

func main() {
	flag.Parse()
	Init()
	runtime.GOMAXPROCS(config.GetGlobalConf().LogSvrConfig.CpuNum) // 设置程序运行核数
}
