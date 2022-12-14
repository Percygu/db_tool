package sync

import (
	"db_tool/config"
	"db_tool/tools"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"time"

	"db_tool/table/tcaplusservice"
	tcaplus "github.com/tencentyun/tcaplusdb-go-sdk/pb"
)

func TableMap() map[string]proto.Message {
	return map[string]proto.Message{
		"tb_player":                     &tcaplusservice.TbPlayer{},
		"tb_account_online":             &tcaplusservice.TbAccountOnline{},
		"tb_activity":                   &tcaplusservice.TbActivity{},
		"tb_battle_record":              &tcaplusservice.TbBattleRecord{},
		"tb_battle_record_brief":        &tcaplusservice.TbBattleRecordBrief{},
		"tb_battle_round_record":        &tcaplusservice.TbBattleRoundRecord{},
		"tb_chat_list":                  &tcaplusservice.TbChatList{},
		"tb_credit_exp_record":          &tcaplusservice.TbCreditExpRecord{},
		"tb_credit_record":              &tcaplusservice.TbCreditRecord{},
		"tb_excellent_moment_list":      &tcaplusservice.TbExcellentMomentList{},
		"tb_friend_list":                &tcaplusservice.TbFriendList{},
		"tb_friend_share":               &tcaplusservice.TbFriendShare{},
		"tb_game_info":                  &tcaplusservice.TbGameInfo{},
		"tb_gamename":                   &tcaplusservice.TbGamename{},
		"tb_global_id":                  &tcaplusservice.TbGlobalId{},
		"tb_group":                      &tcaplusservice.TbGroup{},
		"tb_group_short_info":           &tcaplusservice.TbGroupShortInfo{},
		"tb_guild":                      &tcaplusservice.TbGuild{},
		"tb_guild_apply_list":           &tcaplusservice.TbGuildApplyList{},
		"tb_guildname":                  &tcaplusservice.TbGuildname{},
		"tb_mail_compensation":          &tcaplusservice.TbMailCompensation{},
		"tb_mail_list":                  &tcaplusservice.TbMailList{},
		"tb_negative_compensate_record": &tcaplusservice.TbNegativeCompensateRecord{},
		"tb_openid":                     &tcaplusservice.TbOpenid{},
		"tb_player_guild_invite_list":   &tcaplusservice.TbPlayerGuildInviteList{},
		"tb_player_guild_refuse_list":   &tcaplusservice.TbPlayerGuildRefuseList{},
		"tb_player_order":               &tcaplusservice.TbPlayerOrder{},
		"tb_key_value":                  &tcaplusservice.TbKeyValue{},
		"tb_last_game_result":           &tcaplusservice.TbLastGameResult{},
		"tb_player_statistic":           &tcaplusservice.TbPlayerStatistic{},
		"tb_player_tag":                 &tcaplusservice.TbPlayerTag{},
		"tb_player_tag_result":          &tcaplusservice.TbPlayerTagResult{},
		"tb_reg_cnt":                    &tcaplusservice.TbRegCnt{},
		"tb_report_credit":              &tcaplusservice.TbReportCredit{},
		"tb_role_brief_info":            &tcaplusservice.TbRoleBriefInfo{},
		"tb_role_ext":                   &tcaplusservice.TbRoleExt{},
		"tb_role_online":                &tcaplusservice.TbRoleOnline{},
		"tb_role_online_cnt":            &tcaplusservice.TbRoleOnlineCnt{},
		"tb_room":                       &tcaplusservice.TbRoom{},
		"tb_white_list":                 &tcaplusservice.TbWhiteList{},
	}
}

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
	//record := &tcaplusservice.TbPlayer{}
	record := TableMap()[tableName]
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

// Count 获取表记录总数
func Count(nameSpace, tableName string) error {
	client, err := GetPbClient(nameSpace, tableName)
	if err != nil {
		log.Errorf("traverse err:%v", err)
		return err
	}
	defer client.Close()
	count, err := client.GetTableCount(tableName)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Case Count:")
	fmt.Printf("Count:%d\n", count)
	log.Infof("Count: %d", count)
	return nil
}

func CopyRecord(nameSpaceFrom, nameSpaceTo, tableName string) error {
	getClient, err := GetPbClient(nameSpaceFrom, tableName)
	if err != nil {
		log.Errorf("traverse err:%v", err)
		return err
	}
	defer getClient.Close()
	return nil
}
