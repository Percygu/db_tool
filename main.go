package main

import (
	"db_tool/sync"
	"flag"
	"fmt"

	"db_tool/config"
)

var (
	nameSpace string
	tableName string
)

func init() {
	flag.StringVar(&nameSpace, "name_space", "mainline_dev", "name_space")
	flag.StringVar(&tableName, "table_name", "tb_player", "table_name")
}

// Init 初始化配置
func Init() {
	config.InitConfig() // 初始化配置文件
}

func main() {
	flag.Parse()
	Init()
	tcaplusConf := config.GetGlobalConf().TcaplusConf
	tf := tcaplusConf["mainline_dev"]
	fmt.Printf("tf: %v", tf)

	sync.Count(nameSpace, tableName)
	sync.Traverse(nameSpace, tableName)
}
