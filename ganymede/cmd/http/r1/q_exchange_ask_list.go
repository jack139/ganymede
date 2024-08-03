package r1

import (
	"github.com/jack139/ganymede/ganymede/cmd/http/helper"
	exchangetypes "github.com/jack139/ganymede/ganymede/x/exchange/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/query"

	//"log"
	"bytes"
	"context"
	"encoding/json"
	"github.com/valyala/fasthttp"
)



/* 查询 交换请求 ask 清单 */
func QueryExchangeAskList(ctx *fasthttp.RequestCtx) {

	// POST 的数据
	content := ctx.PostBody()

	// 验签
	reqData, err := helper.CheckSign(content)
	if err != nil {
		helper.RespError(ctx, 9000, err.Error())
		return
	}

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
	senderAddr, ok := (*reqData)["asker_addr"].(string)
	replierAddr, ok := (*reqData)["replier_addr"].(string)
	askUuid, ok := (*reqData)["uuid"].(string)

	if page < 1 || limit < 1 {
		helper.RespError(ctx, 9002, "page and limit need begin from 1")
		return		
	}

	if len(senderAddr) > 0 && len(replierAddr) > 0 {
		helper.RespError(ctx, 9002, "only need one addr: sender OR replier")
		return				
	}

	if len(senderAddr) == 0 && len(replierAddr) == 0 && len(askUuid) > 0 {
		helper.RespError(ctx, 9002, "uuid should exist with sender or replier")
		return				
	}

	// 查询链上数据
	respData2, err := queryExchangeAskListPage(uint64(page), uint64(limit), senderAddr, replierAddr, askUuid)
	if err!=nil{
		helper.RespError(ctx, 9021, err.Error())
		return
	}
	dataList := *respData2

	// 构建返回结构
	respData := make([]map[string]interface{}, 0) 

	for _, item0 := range dataList {
		item := item0.(map[string]interface{})

		// 解析 payload json
		var askData map[string]string
		if err := json.Unmarshal([]byte(item["payload"].(string)), &askData); err != nil {  // 解析 json
			helper.RespError(ctx, 9022, err.Error())
			return
		}

		newItem := map[string]interface{} {
			"ask_id"     : item["id"],
			"asker_addr"   : item["sender"],
			"replier_addr" : item["replier"],
			"payload"    : askData["text"],
			"uuid"       : askData["uuid"],
			"sent_date"  : item["sentDate"],
		}
		respData = append(respData, newItem)
	}

	resp := map[string] interface{} {
		"ask_list" : respData,
	}

	helper.RespJson(ctx, &resp)
}


// 查询链上数据, 返回 map
func queryExchangeAskListPage(page uint64, limit uint64, 
	senderAddr string, replierAddr string, askUuid string) (*[]interface{}, error) {
	// 获取 ctx 上下文
	clientCtx := client.GetClientContextFromCmd(helper.HttpCmd)

	// 准备查询
	queryClient := exchangetypes.NewQueryClient(clientCtx)

	if len(askUuid) == 0 {
		askUuid = "|"
	} else {
		askUuid = "|" + askUuid
	}

	// 设置 接收输出
	buf := new(bytes.Buffer)
	clientCtx.Output = buf
	clientCtx.OutputFormat = "json"

	var jsonName string

	if len(senderAddr)==0 && len(replierAddr)==0 { // 查所有的
		pageReq := query.PageRequest{
			Key:        []byte(""),
			Offset:     (page - 1) * limit,
			Limit:      limit,
			CountTotal: true,
		}

		params := &exchangetypes.QueryAllAskRequest{
			Pagination: &pageReq,
		}

		res, err := queryClient.AskAll(context.Background(), params)
		if err != nil {
			return nil, err
		}

		// 转换输出
		clientCtx.PrintProto(res)

		jsonName = "Ask"

	} else if len(senderAddr) > 0 { // 查指定 sender 的
		params := &exchangetypes.QueryListAskBySenderRequest{
			Page: page,
			Limit: limit,
			Sender: senderAddr + askUuid,
		}		

		res, err := queryClient.ListAskBySender(context.Background(), params)
		if err != nil {
			return nil, err
		}

		// 转换输出
		clientCtx.PrintProto(res)

		jsonName = "ask"

	} else { // 查指定 replier 的
		params := &exchangetypes.QueryListAskByReplierRequest{
			Page: page,
			Limit: limit,
			Replier: replierAddr + askUuid,
		}		

		res, err := queryClient.ListAskByReplier(context.Background(), params)
		if err != nil {
			return nil, err
		}

		// 转换输出
		clientCtx.PrintProto(res)

		jsonName = "ask"
	}

	// 输出的字节流
	respBytes := []byte(buf.String())

	//log.Println("output: ", buf.String())

	// 转换成map, 生成返回数据
	var respData map[string]interface{}

	if err := json.Unmarshal(respBytes, &respData); err != nil {
		return nil, err
	}

	itemMapList := respData[jsonName].([]interface{})

	return &(itemMapList), nil
}



// 查询链上数据, 返回 map -- 返回单个数据
func queryExchangeAskInfoById(askId uint64) (*map[string]interface{}, error) {
	// 获取 ctx 上下文
	clientCtx := client.GetClientContextFromCmd(helper.HttpCmd)

	// 准备查询
	queryClient := exchangetypes.NewQueryClient(clientCtx)

	params := &exchangetypes.QueryGetAskRequest{
		Id: askId,
	}

	res, err := queryClient.Ask(context.Background(), params)
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

	itemMap := respData["Ask"].(map[string]interface{})

	return &(itemMap), nil
}
