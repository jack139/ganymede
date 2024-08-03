package r1

import (
	"github.com/jack139/ganymede/ganymede/cmd/http/helper"
	"github.com/jack139/ganymede/ganymede/cmd/exchange"
	exchangetypes "github.com/jack139/ganymede/ganymede/x/exchange/types"
	"github.com/jack139/ganymede/ganymede/cmd/ipfs"

	"log"
	"bytes"
	"time"
	"strings"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

/* 新建数据交换请求响应 */

func TxExchangeReply(ctx *fasthttp.RequestCtx) {

	// POST 的数据
	content := ctx.PostBody()

	// 验签
	reqData, err := helper.CheckSign(content)
	if err != nil {
		helper.RespError(ctx, 9000, err.Error())
		return
	}

	// 检查参数
	senderAddr, ok := (*reqData)["asker_addr"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need asker_addr")
		return
	}
	replierAddr, ok := (*reqData)["replier_addr"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need replier_addr")
		return
	}

	askId, ok := (*reqData)["ask_id"].(float64)
	if !ok {
		helper.RespError(ctx, 9001, "need ask_id")
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


	// 检查 replierAddr 合法性
	userMap, err := queryUserInfoByChainAddr(replierAddr, nil)
	if err!=nil {
		helper.RespError(ctx, 9004, err.Error())
		return		
	}
	if (*userMap)["status"] != "ACTIVE" {
		helper.RespError(ctx, 9004, "the user is not ACTIVE")
		return				
	}


	// 查询 ask_id 的 payload
	askData, err := queryExchangeAskInfoById(uint64(askId))
	if err!=nil{
		helper.RespError(ctx, 9021, "get Ask data fail: " + err.Error())
		return
	}

	// 检查 replierAddr 合法性 -- 一致性
	if replierAddr != (*askData)["replier"].(string) {
		helper.RespError(ctx, 9004, "wrong replier_addr")
		return
	}

	if senderAddr != (*askData)["sender"].(string) {
		helper.RespError(ctx, 9004, "wrong asker_addr")
		return
	}

	// 是否需要存储到IPFS
	payloadStr, err = ipfs.SaveToIpfsIfPossible(payloadStr, false)
	if err!=nil {
		helper.RespError(ctx, 9031, "IPFS fail: " + err.Error())
		return
	}

	// 生成 payload 数据
	payload, askUuid, err := exchange.GenerateReplyPayload(replyBool, replierAddr, (*askData)["payload"].(string), payloadStr)
	if err!=nil{
		helper.RespError(ctx, 9022, "make payload fail: " + err.Error())
		return
	}


	/* 信号量 */
	helper.AcquireSem(replierAddr)
	defer helper.ReleaseSem(replierAddr)


	// 设置 caller_addr
	originFlagFrom, err := helper.HttpCmd.Flags().GetString(flags.FlagFrom) // 保存 --from 设置
	if err != nil {
		helper.RespError(ctx, 9023, err.Error())
		return
	}
	helper.HttpCmd.Flags().Set(flags.FlagFrom, replierAddr)  // 设置 --from 地址
	defer helper.HttpCmd.Flags().Set(flags.FlagFrom, originFlagFrom)  // 结束时恢复 --from 设置


	// 获取 ctx 上下文
	clientCtx, err := client.GetClientTxContext(helper.HttpCmd)
	if err != nil {
		helper.RespError(ctx, 9024, err.Error())
		return
	}


	// 数据上链
	msg := exchangetypes.NewMsgNewReply(
		replierAddr, // creator string, 
		(*askData)["id"].(string), // askId string, 
		senderAddr, // sender string, 
		replierAddr, // replier string, 
		payload, // payload string
		time.Now().Format("2006-01-02 15:04:05"), // sentDate
	)
	if err := msg.ValidateBasic(); err != nil {
		helper.RespError(ctx, 9025, err.Error())
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
			helper.RespError(ctx, 9026, err.Error())
		}
		return		
	}

	// 结果输出
	respBytes := []byte(buf.String())

	//log.Println("output: ", buf.String())

	// 转换成map, 生成返回数据
	var respData map[string]interface{}

	if err := json.Unmarshal(respBytes, &respData); err != nil {
		helper.RespError(ctx, 9027, err.Error())
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
