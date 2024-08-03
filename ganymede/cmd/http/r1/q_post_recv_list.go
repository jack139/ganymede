package r1

import (
	"github.com/jack139/ganymede/ganymede/cmd/http/helper"
	posttypes "github.com/jack139/ganymede/ganymede/x/postoffice/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/query"

	//"log"
	"bytes"
	"context"
	"strings"
	"encoding/json"
	"github.com/valyala/fasthttp"
)



/* 查询 跨链 recv 清单 */
func QueryPostRecvList(ctx *fasthttp.RequestCtx) {

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
	senderAddr, ok := (*reqData)["sender_addr"].(string)
	targetAddr, ok := (*reqData)["target_addr"].(string)
	askUuid, ok := (*reqData)["uuid"].(string)

	if page < 1 || limit < 1 {
		helper.RespError(ctx, 9002, "page and limit need begin from 1")
		return		
	}

	if len(senderAddr) > 0 && len(targetAddr) > 0 {
		helper.RespError(ctx, 9002, "only need one addr: sender OR replier")
		return				
	}

	if len(senderAddr) == 0 && len(targetAddr) == 0 && len(askUuid) > 0 {
		helper.RespError(ctx, 9002, "uuid should exist with sender or replier")
		return				
	}

	// 查询链上数据
	respData2, err := queryPostRecvListPage(uint64(page), uint64(limit), senderAddr, targetAddr, askUuid)
	if err!=nil{
		helper.RespError(ctx, 9021, err.Error())
		return
	}
	dataList := *respData2

	// 构建返回结构
	respData := make([]map[string]interface{}, 0) 

	for _, item0 := range dataList {
		item := item0.(map[string]interface{})

		newItem := map[string]interface{} {
			//"sender_info" : item["senderInfo"],
			"post_id"      : item["id"],
			"post_type"    : "POST",
			"post_channel" : item["fromChain"],
			//"payload"     : item["payload"],
			"recv_date"    : item["sentDate"],
		}

		// 解析 post 类型
		if strings.HasPrefix(item["title"].(string), "@EXCH:ASK:") {
			// 解析 payload json
			var askData map[string]string
			if err := json.Unmarshal([]byte(item["payload"].(string)), &askData); err != nil {  // 解析 json
				helper.RespError(ctx, 9022, err.Error())
				return
			}
			newItem["asker_addr"]   = item["sender"]
			newItem["replier_addr"] = item["receiver"]
			newItem["ask_post_id"]  = item["id"]
			newItem["post_type"]    = "EXCH:ASK"
			newItem["uuid"]         = item["title"].(string)[10:]
			newItem["payload"]      = askData["text"]

		} else if strings.HasPrefix(item["title"].(string), "@EXCH:RPLY:") {
			// 解析 payload json
			var replyData map[string]string
			if err := json.Unmarshal([]byte(item["payload"].(string)), &replyData); err != nil {  // 解析 json
				helper.RespError(ctx, 9023, err.Error())
				return
			}
			newItem["asker_addr"]   = item["receiver"]
			newItem["replier_addr"] = item["sender"]
			//newItem["reply_post_id"] = item["id"]
			newItem["post_type"]    = "EXCH:RPLY"
			newItem["reply"]        = replyData["reply"]=="PASS"

			// title 分解 出 uuid 和 ask_id
			params := strings.Split(item["title"].(string)[11:], "|")
			if len(params) > 1 { 
				newItem["uuid"] = params[0]
				newItem["ask_post_id"] = params[1]
			} else {
				newItem["uuid"] = params[0]
				newItem["ask_post_id"] = ""
			}

		} else {
			newItem["sender_addr"] = item["sender"]
			newItem["target_addr"] = item["receiver"]
			newItem["uuid"]        = item["title"]
		}

		respData = append(respData, newItem)
	}

	resp := map[string] interface{} {
		"recv_list" : respData,
	}

	helper.RespJson(ctx, &resp)
}


// 查询链上数据, 返回 map
func queryPostRecvListPage(page uint64, limit uint64, 
	senderAddr string, targetAddr string, askUuid string) (*[]interface{}, error) {
	// 获取 ctx 上下文
	clientCtx := client.GetClientContextFromCmd(helper.HttpCmd)

	// 准备查询
	queryClient := posttypes.NewQueryClient(clientCtx)

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

	if len(senderAddr)==0 && len(targetAddr)==0 { // 查所有的
		pageReq := query.PageRequest{
			Key:        []byte(""),
			Offset:     (page - 1) * limit,
			Limit:      limit,
			CountTotal: true,
		}

		params := &posttypes.QueryAllPostRequest{
			Pagination: &pageReq,
		}

		res, err := queryClient.PostAll(context.Background(), params)
		if err != nil {
			return nil, err
		}

		// 转换输出
		clientCtx.PrintProto(res)

		jsonName = "Post"

	} else if len(senderAddr) > 0 { // 查指定 sender 的
		params := &posttypes.QueryListPostBySenderRequest{
			Page: page,
			Limit: limit,
			Sender: senderAddr + askUuid,
		}		

		res, err := queryClient.ListPostBySender(context.Background(), params)
		if err != nil {
			return nil, err
		}

		// 转换输出
		clientCtx.PrintProto(res)

		jsonName = "post"

	} else { // 查指定 replier 的
		params := &posttypes.QueryListPostByReceiverRequest{
			Page: page,
			Limit: limit,
			Receiver: targetAddr + askUuid,
		}		

		res, err := queryClient.ListPostByReceiver(context.Background(), params)
		if err != nil {
			return nil, err
		}

		// 转换输出
		clientCtx.PrintProto(res)

		jsonName = "post"
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
