package r1

import (
	"github.com/jack139/ganymede/ganymede/cmd/http/helper"
	"github.com/jack139/ganymede/ganymede/cmd/exchange"
	exchangetypes "github.com/jack139/ganymede/ganymede/x/exchange/types"
	"github.com/jack139/ganymede/ganymede/cmd/ipfs"

	"github.com/cosmos/cosmos-sdk/client"

	"bytes"
	"context"
	//"log"
	"strconv"
	"encoding/json"
	"github.com/valyala/fasthttp"
)



/* 查询 请求交换 reply 数据信息 */
func QueryExchangeReplyInfo(ctx *fasthttp.RequestCtx) {

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
	replyId, ok := (*reqData)["reply_id"].(float64)
	if !ok {
		helper.RespError(ctx, 9001, "need reply_id")
		return
	}
	doDecrypt, ok := (*reqData)["decrypt"].(bool)
	if !ok {
		doDecrypt = false
	}

	// 查询链上数据
	replyData, err := queryExchangeReplyInfoById(uint64(replyId))
	if err!=nil{
		helper.RespError(ctx, 9021, "get Reply data fail: " + err.Error())
		return
	}

	// 坚持 sender 合法性
	if senderAddr != (*replyData)["sender"].(string) {
		helper.RespError(ctx, 9004, "wrong senderAddr")
		return		
	}

	askIdUint64, err := strconv.ParseUint((*replyData)["askId"].(string), 10, 64)
	if err!=nil{
		helper.RespError(ctx, 9022, err.Error())
		return
	}

	// 查询 ask_id 的 payload
	askData, err := queryExchangeAskInfoById(askIdUint64)
	if err!=nil{
		helper.RespError(ctx, 9023, "get Ask data fail: " + err.Error())
		return
	}

	// 解析 reply 数据
	replyBool, plainPayload, askUuid, err := exchange.QueryReplyPayload(senderAddr,
				(*askData)["payload"].(string), (*replyData)["payload"].(string))
	if err!=nil{
		helper.RespError(ctx, 9024, "query plain payload fail: " + err.Error())
		return
	}

	if !doDecrypt { // 不解密 显示原始文本
		plainPayload = (*replyData)["payload"].(string)
	} else {
		// 是否需要从IPFS取回
		plainPayload, err = ipfs.GetFromIpfsIfPossible(plainPayload, false)
		if err!=nil {
			helper.RespError(ctx, 9032, "IPFS fail:" + err.Error())
			return
		}
	}

	// 构建返回结构
	respData := map[string]interface{} {
		"reply_id"     : (*replyData)["id"],
		"ask_id"       : (*replyData)["askId"],
		"asker_addr"   : (*replyData)["sender"],
		"replier_addr" : (*replyData)["replier"],
		"uuid"         : askUuid,
		"reply"        : replyBool,
		"payload"      : plainPayload,
		"sent_date"    : (*replyData)["sentDate"],
	}

	resp := map[string] interface{} {
		"reply" : respData,
	}

	helper.RespJson(ctx, &resp)
}


// 查询链上数据, 返回 map
func queryExchangeReplyInfoById(replyId uint64) (*map[string]interface{}, error) {
	// 获取 ctx 上下文
	clientCtx := client.GetClientContextFromCmd(helper.HttpCmd)

	// 准备查询
	queryClient := exchangetypes.NewQueryClient(clientCtx)

	params := &exchangetypes.QueryGetReplyRequest{
		Id: replyId,
	}

	res, err := queryClient.Reply(context.Background(), params)
	if err != nil {
		return nil, err
	}

	//log.Printf("%T\n", res)

	// 设置 接收输出
	buf := new(bytes.Buffer)
	clientCtx.Output = buf
	clientCtx.OutputFormat = "json"

	// 转换输出
	clientCtx.PrintProto(res)

	// 输出的字节流
	respBytes := []byte(buf.String())

	//log.Println("output: ", buf.String())

	// 转换成map, 生成返回数据
	var respData map[string]interface{}

	if err := json.Unmarshal(respBytes, &respData); err != nil {
		return nil, err
	}

	itemMap := respData["Reply"].(map[string]interface{})

	return &(itemMap), nil
}
