package r1

import (
	"github.com/jack139/ganymede/ganymede/cmd/http/helper"

	"github.com/cosmos/cosmos-sdk/client"
	//"github.com/cosmos/cosmos-sdk/codec/legacy"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"

	//"context"
	"encoding/json"
	"github.com/valyala/fasthttp"
	//"log"
	"fmt"
	"bytes"
	//"strconv"
)


/* 获取交易区块数据, 指定过滤条件 */
func getBlockTxs(clientCtx client.Context, tmEvents []string, page, limit int) ([]byte, error) {
	if len(tmEvents) == 0 {
		return nil, fmt.Errorf("need event filter")
	}

	// If hash is given, then query the tx by hash.
	output, err := authtx.QueryTxsByEvents(clientCtx, tmEvents, page, limit, "")
	if err != nil {
		return nil, err
	}

	// 设置 接收输出
	buf := new(bytes.Buffer)
	clientCtx.Output = buf
	clientCtx.OutputFormat = "json"

	// 转换输出
	clientCtx.PrintProto(output)

	return []byte(buf.String()), nil
}


/* 指定区块查询交易, 指定过滤条件 */
func QueryBlockTxs(ctx *fasthttp.RequestCtx) {

	// POST 的数据
	content := ctx.PostBody()

	// 验签
	reqData, err := helper.CheckSign(content)
	if err != nil {
		helper.RespError(ctx, 9000, err.Error())
		return
	}

	// 检查参数
	var page, limit float64

	// 检查参数
	page, ok := (*reqData)["page"].(float64)
	if !ok {
		page = 1.0
	}
	limit, ok = (*reqData)["limit"].(float64)
	if !ok {
		limit = 50.0
	}

	creatorAddr, ok := (*reqData)["creator_addr"].(string)
	messageAction, ok := (*reqData)["tx_action"].(string)

	if len(creatorAddr)==0 && len(messageAction)==0 {
		helper.RespError(ctx, 9001, "need one of creator_addr or tx_action")
		return
	}

	// 获取 ctx 上下文
	clientCtx, err := client.GetClientTxContext(helper.HttpCmd)
	if err != nil {
		helper.RespError(ctx, 9021, err.Error())
		return
	}

	// 准备查询
	var tmEvents []string

	if len(creatorAddr) > 0 {
		tmEvents = append(tmEvents, fmt.Sprintf("tx.fee_payer='%s'", creatorAddr))
	}

	if len(messageAction) > 0 {
		var action string

		switch messageAction {
		case "user/new":
			action = "/cosmostest.ganymede.ganymede.MsgCreateUsers"
		case "user/update":
			action = "/cosmostest.ganymede.ganymede.MsgUpdateUsers"
		case "kv/new":
			action = "/cosmostest.ganymede.zoo.MsgCreateKvzoo"
		case "kv/update":
			action = "/cosmostest.ganymede.zoo.MsgUpdateKvzoo"
		case "kv/delete":
			action = "/cosmostest.ganymede.zoo.DeleteKvzoo"
		case "exchange/ask":
			action = "/cosmostest.ganymede.exchange.MsgNewAsk"
		case "exchange/reply":
			action = "/cosmostest.ganymede.exchange.MsgNewReply"
		case "post/send":
			action = "/cosmostest.ganymede.postoffice.MsgSendIbcPost"
		default:
			helper.RespError(ctx, 9002, "unknown tx_action")
			return
		}

		tmEvents = append(tmEvents, fmt.Sprintf("message.action='%s'", action))
	}

	respBytes, err := getBlockTxs(clientCtx, tmEvents, int(page), int(limit))
	if err != nil {
		helper.RespError(ctx, 9022, err.Error())
		return
	}

	//log.Printf("%v\n", string(respBytes))

	// 转换成map, 生成返回数据
	var respData map[string]interface{}
	if len(respBytes) > 0 {
		if err := json.Unmarshal(respBytes, &respData); err != nil {
			helper.RespError(ctx, 9023, err.Error())
			return
		}
	}
	resp := map[string]interface{}{
		"txs": respData,
	}

	helper.RespJson(ctx, &resp)
}

