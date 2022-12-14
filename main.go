package main

import (
	"db_tool/sync"
	"flag"
	"fmt"

	"db_tool/config"
)

var (
	nameSpaceFrom string
	nameSpaceTo   string
	tableName     string
)

func init() {
	flag.StringVar(&nameSpaceFrom, "name_space_from", "mainline_dev", "name_space_from")
	flag.StringVar(&nameSpaceTo, "name_space_to", "mainline_dev", "name_space_to")
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

	sync.Count(nameSpaceFrom, tableName)
	sync.CopyRecord(nameSpaceFrom, nameSpaceTo, tableName, tf.ZoneID)
}
