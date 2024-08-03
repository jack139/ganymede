package r1

import (
	"github.com/jack139/ganymede/ganymede/cmd/http/helper"
	"github.com/jack139/ganymede/ganymede/cmd/exchange"
	posttypes "github.com/jack139/ganymede/ganymede/x/postoffice/types"
	"github.com/jack139/ganymede/ganymede/cmd/ipfs"

	"log"
	"time"
	"bytes"
	"strings"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	channelutils "github.com/cosmos/ibc-go/v7/modules/core/04-channel/client/utils"
)

/* 新建 跨链 数据交换请求响应 */

func TxPostReply(ctx *fasthttp.RequestCtx) {

	// POST 的数据
	content := ctx.PostBody()

	// 验签
	reqData, err := helper.CheckSign(content)
	if err != nil {
		helper.RespError(ctx, 9000, err.Error())
		return
	}

	// 检查参数
	replierAddr, ok := (*reqData)["asker_addr"].(string) // asker 是 post 收件人
	if !ok {
		helper.RespError(ctx, 9001, "need asker_addr")
		return
	}
	senderAddr, ok := (*reqData)["replier_addr"].(string) // replier 是 post 发件人
	if !ok {
		helper.RespError(ctx, 9001, "need replier_addr")
		return
	}
	postChannel, ok := (*reqData)["post_channel"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need post_channel")
		return
	}
	askId, ok := (*reqData)["ask_post_id"].(float64)
	if !ok {
		helper.RespError(ctx, 9001, "need ask_post_id")
		return
	}
	replyBool, ok := (*reqData)["reply"].(bool)
	if !ok {
		helper.RespError(ctx, 9001, "need reply")
		return
	}
	payloadStr, ok := (*reqData)["payload"].(string)
	if !ok && replyBool { // replyBool==true 时必须有 payload
		helper.RespError(ctx, 9001, "need payload")
		return
	}

	// 检查数据尺寸
	if ok:=ipfs.CheckSize(payloadStr); !ok {
		helper.RespError(ctx, 9003, "payloadStr is too long")
		return
	}


	// 检查 senderAddr 合法性
	userMap, err := queryUserInfoByChainAddr(senderAddr, nil)
	if err!=nil {
		helper.RespError(ctx, 9004, err.Error())
		return		
	}
	if (*userMap)["status"] != "ACTIVE" {
		helper.RespError(ctx, 9004, "the user is not ACTIVE")
		return				
	}


	// 查询 ask_id 的 payload
	askData, err := queryPostRecvInfoById(uint64(askId))
	if err!=nil{
		helper.RespError(ctx, 9021, "get Ask data fail: " + err.Error())
		return
	}

	if !strings.HasPrefix((*askData)["title"].(string), "@EXCH:ASK:") {
		helper.RespError(ctx, 9022, "not a EXCHANGE post id")
		return
	}

	// 检查 replierAddr 合法性 -- 一致性
	if replierAddr != (*askData)["sender"].(string) {
		helper.RespError(ctx, 9004, "wrong asker_addr")
		return
	}

	if senderAddr != (*askData)["receiver"].(string) {
		helper.RespError(ctx, 9004, "wrong replier_addr")
		return
	}

	// 是否需要存储到IPFS
	payloadStr, err = ipfs.SaveToIpfsIfPossible(payloadStr, true) // 'true' for use IBC ipfs
	if err!=nil {
		helper.RespError(ctx, 9031, "IPFS fail: " + err.Error())
		return
	}

	// 生成 payload 数据
	payload, askUuid, err := exchange.GenerateReplyPayload(replyBool, senderAddr, (*askData)["payload"].(string), payloadStr)
	if err!=nil{
		helper.RespError(ctx, 9023, "make payload fail: " + err.Error())
		return
	}

	// 生成 post 的 title:  "@EXCH:RPLY:" + "<ask_uuid>|<ask_id>"
	title := "@EXCH:RPLY:" + askUuid + "|" + (*askData)["id"].(string)

	// 生成 post 的 content
	// content 的 json 格式： { "sender" : "", "receiver" : "", "payload" : "", "sender_info" : "" }
	var contentData = make(map[string]interface{})
	contentData["sender"] = senderAddr
	contentData["receiver"] = replierAddr
	contentData["sender_info"] = ""
	contentData["payload"] = payload

	contentJson, err := json.Marshal(contentData)
	if err != nil {
		helper.RespError(ctx, 9024, err.Error())
		return
	}

	/* 信号量 */
	helper.AcquireSem(senderAddr)
	defer helper.ReleaseSem(senderAddr)


	// 设置 caller_addr
	originFlagFrom, err := helper.HttpCmd.Flags().GetString(flags.FlagFrom) // 保存 --from 设置
	if err != nil {
		helper.RespError(ctx, 9025, err.Error())
		return
	}
	helper.HttpCmd.Flags().Set(flags.FlagFrom, senderAddr)  // 设置 --from 地址
	defer helper.HttpCmd.Flags().Set(flags.FlagFrom, originFlagFrom)  // 结束时恢复 --from 设置


	// 获取 ctx 上下文
	clientCtx, err := client.GetClientTxContext(helper.HttpCmd)
	if err != nil {
		helper.RespError(ctx, 9026, err.Error())
		return
	}

	// Get the relative timeout timestamp
	consensusState, _, _, err := channelutils.QueryLatestConsensusState(clientCtx, "postoffice", postChannel)
	if err != nil {
		helper.RespError(ctx, 9027, err.Error())
		return
	}
	timeoutTimestamp := consensusState.GetTimestamp() + DefaultRelativePacketTimeoutTimestamp


	// 数据上链
	msg := posttypes.NewMsgSendIbcPost(
		senderAddr, // creator string,
		"postoffice", // port string,
		postChannel, // channelID string,
		timeoutTimestamp, // timeoutTimestamp uint64,
		title, // title string,
		string(contentJson), // content string,
		time.Now().Format("2006-01-02 15:04:05"), // sentDate
	)
	if err := msg.ValidateBasic(); err != nil {
		helper.RespError(ctx, 9028, err.Error())
		return
	}

	// 设置 接收输出
	buf := new(bytes.Buffer)
	clientCtx.Output = buf
	clientCtx.OutputFormat = "json"
	clientCtx.BroadcastMode = flags.BroadcastSync // 默认是 flags.BroadcastSync

	err = helper.BroadcastTxWithRetry(clientCtx, msg)
	if err != nil {
		if strings.Contains(err.Error(), "account sequence mismatch") {
			helper.RespError(ctx, 9090, "Same user's TX busy! please try again later.")
		} else {
			helper.RespError(ctx, 9029, err.Error())
		}
		return		
	}

	// 结果输出
	respBytes := []byte(buf.String())

	//log.Println("output: ", buf.String())

	// 转换成map, 生成返回数据
	var respData map[string]interface{}

	if err := json.Unmarshal(respBytes, &respData); err != nil {
		helper.RespError(ctx, 9030, err.Error())
		return
	}

	// code==0 提交成功
	if respData["code"].(float64)!=0 { 
		helper.RespError(ctx, 9099, buf.String())  ///  提交失败
		return
	}

	// 返回交易hash
	resp := map[string]interface{}{
		"uuid"   : askUuid,
		"txhash" : respData["txhash"].(string),  // 交易hash
	}

	log.Println(resp)

	helper.RespJson(ctx, &resp)
}
