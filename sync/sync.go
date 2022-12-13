package sync

import (
	"db_tool/config"
	"db_tool/tools"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"

	"db_tool/table/tcaplusservice"
	tcaplus "github.com/tencentyun/tcaplusdb-go-sdk/pb"
)

//声明一个TcaplusDB连接客户端
var client *tcaplus.PBClient

func GetPbClient(nameSpace, tableName string) (*tcaplus.PBClient, error) {
	if _, ok := config.GetGlobalConf().TcaplusConf[nameSpace]; !ok {
		log.Errorf("name_space not invalid")
		return nil, fmt.Errorf("name_space not invalid")
	}
	tcaplusConf := config.GetGlobalConf().TcaplusConf[nameSpace]
	client, err := tools.InitPBSyncClient(tcaplusConf, tableName)
	if err != nil {
		log.Errorf("GetPbClient err: %v", err)
		return nil, err
	}
	return client, nil
}

// 遍历记录
func Traverse(nameSpace, tableName string) error {
	client, err := GetPbClient(nameSpace, tableName)
	if err != nil {
		log.Errorf("traverse err:%v", err)
		return err
	}
	defer client.Close()
	record := &tcaplusservice.TbPlayer{}
	// 遍历时间可能比较长超时时间设长一些
	client.SetDefaultTimeOut(30 * time.Second)
	msgs, err := client.Traverse(record)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Case Traverse:")
	fmt.Printf("message:%+v\n", msgs)
	return nil
}

// 获取表记录总数
func Count(nameSpace, tableName string) error {
	client, err := GetPbClient(nameSpace, tableName)
	if err != nil {
		log.Errorf("traverse err:%v", err)
		return err
	}
	defer client.Close()
	count, err := client.GetTableCount("game_players")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Case Count:")
	fmt.Printf("Count:%d\n", count)
	log.Infof("Count: %d", count)
	return nil
}
