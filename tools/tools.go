package tools

import (
	"fmt"
	"sync"

	"db_tool/config"
	tcaplus "github.com/tencentyun/tcaplusdb-go-sdk/pb"
)

var (
	once sync.Once
)

func InitPBSyncClient(tcaplusConfig *config.TcaplusConfig, tableName string) (*tcaplus.PBClient, error) {
	var err error

	pbClient := tcaplus.NewPBClient()
	zoneList := []uint32{tcaplusConfig.ZoneID}
	zoneTable := make(map[uint32][]string)
	//构造Map对象存储对应表格组下所有的表
	zoneTable[tcaplusConfig.ZoneID] = []string{tableName}

	err = pbClient.Dial(tcaplusConfig.APPID, zoneList, tcaplusConfig.DirUrl, tcaplusConfig.Signature, 30, zoneTable)
	if err != nil {
		fmt.Printf("excepted dial success, %s", err.Error())
		return nil, err
	}
	return pbClient, nil
}
