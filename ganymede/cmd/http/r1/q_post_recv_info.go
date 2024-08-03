package r1

import (
	"github.com/jack139/ganymede/ganymede/cmd/http/helper"
	"github.com/jack139/ganymede/ganymede/cmd/exchange"
	posttypes "github.com/jack139/ganymede/ganymede/x/postoffice/types"
	"github.com/jack139/ganymede/ganymede/cmd/ipfs"

	"github.com/cosmos/cosmos-sdk/client"

	"bytes"
	"context"
	"log"
	"strings"
	"strconv"
	"encoding/json"
	"github.com/valyala/fasthttp"
)



/* 查询 跨链 recv 数据信息 */
func QueryPostRecvInfo(ctx *fasthttp.RequestCtx) {

	// POST 的数据
	content := ctx.PostBody()

	// 验签
	reqData, err := helper.CheckSign(content)
	if err != nil {
		helper.RespError(ctx, 9000, err.Error())
		return
	}

	// 检查参数
	targetAddr, ok := (*reqData)["target_addr"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need target_addr")
		return
	}
	postId, ok := (*reqData)["post_id"].(float64)
	if !ok {
		helper.RespError(ctx, 9001, "need post_id")
		return
	}
	doDecrypt, ok := (*reqData)["decrypt"].(bool)
	if !ok {
		doDecrypt = false
	}

	// 检查 receiver 合法性 -- 是否是正常链用户
	_, err = queryUserInfoByChainAddr(targetAddr, nil)
	if err!=nil {
		helper.RespError(ctx, 9004, err.Error())
		return		
	}


	// 查询链上数据
	respData2, err := queryPostRecvInfoById(uint64(postId))
	if err!=nil{
		helper.RespError(ctx, 9021, err.Error())
		return
	}
	itemMap := *respData2

	// 检查 receiver 合法性 -- 一致性
	if targetAddr != itemMap["receiver"].(string) {
		helper.RespError(ctx, 9004, "wrong target_addr")
		return		
	}

	// 构建返回结构
	respData := map[string]interface{} {
		//"sender_info" : itemMap["senderInfo"],
		"post_id"     : itemMap["id"],
		"post_type"   : "POST",
		"post_channel": itemMap["fromChain"],
		"payload"     : itemMap["payload"],
		"recv_date"   : itemMap["sentDate"],
	}

	// 解析 post 类型
	if strings.HasPrefix(itemMap["title"].(string), "@EXCH:ASK:") { // ask 数据
		// 解析 payload json
		var askData map[string]string
		if err := json.Unmarshal([]byte(itemMap["payload"].(string)), &askData); err != nil {  // 解析 json
			helper.RespError(ctx, 9022, err.Error())
			return
		}
		respData["asker_addr"]   = itemMap["sender"]
		respData["replier_addr"] = itemMap["receiver"]
		respData["ask_post_id"]  = itemMap["id"]
		respData["post_type"]    = "EXCH:ASK"
		respData["uuid"]         = itemMap["title"].(string)[10:]
		respData["payload"]      = askData["text"]

	} else if strings.HasPrefix(itemMap["title"].(string), "@EXCH:RPLY:") { // reply 数据
		// 解析 payload json
		var replyData map[string]string
		if err := json.Unmarshal([]byte(itemMap["payload"].(string)), &replyData); err != nil {  // 解析 json
			helper.RespError(ctx, 9023, err.Error())
			return
		}
		respData["asker_addr"]   = itemMap["receiver"]
		respData["replier_addr"] = itemMap["sender"]
		//respData["reply_post_id"] = itemMap["id"]
		respData["post_type"]    = "EXCH:RPLY"
		respData["reply"]        = replyData["reply"]=="PASS"

		// title 分解 出 uuid 和 ask_id
		params := strings.Split(itemMap["title"].(string)[11:], "|")
		if len(params) > 1 { 
			respData["uuid"] = params[0]
			respData["ask_post_id"] = params[1]
		} else {
			respData["uuid"] = params[0]
			respData["ask_post_id"] = ""
		}

		if doDecrypt { // 解密 
			// 通过 uuid 取得 sent_id
			askData2, err := queryPostSentListPage(1, 50000,
				respData["asker_addr"].(string), "", respData["uuid"].(string))
			if err!=nil{
				helper.RespError(ctx, 9027, err.Error())
				return
			}

			if len(*askData2) == 0 { // 返回是是一个 列表， 应该只有一个元素，因为在ask里 uuid 唯一
				helper.RespError(ctx, 9028, "get Ask data fail: no data")
				return
			}

			sentId := (*askData2)[0].(map[string]interface{})["id"].(string)
			log.Print("sent_id: ", sentId)

			// 转换 sent_id 为 uint64
			sentIdUint64, err := strconv.ParseUint(sentId, 10, 64)
			if err!=nil{
				helper.RespError(ctx, 9024, err.Error())
				return
			}

			// 查询 ask_id 的 payload
			askData, err := queryPostSentInfoById(sentIdUint64)
			if err!=nil{
				helper.RespError(ctx, 9025, "get Ask data fail: " + err.Error())
				return
			}

			// 解析 reply 数据
			_, plainPayload, _, err := exchange.QueryReplyPayload(targetAddr,
						(*askData)["payload"].(string), itemMap["payload"].(string))
			if err!=nil{
				helper.RespError(ctx, 9026, "query plain payload fail: " + err.Error())
				return
			}

			// 是否需要从IPFS取回
			plainPayload, err = ipfs.GetFromIpfsIfPossible(plainPayload, true) // 'true' for use IBC ipfs
			if err!=nil {
				helper.RespError(ctx, 9032, "IPFS fail:" + err.Error())
				return
			}

			respData["payload"] = plainPayload // 明文
		}

	} else { // 普通 post 数据
		respData["sender_addr"] = itemMap["sender"]
		respData["target_addr"] = itemMap["receiver"]
		respData["uuid"]        = itemMap["title"]
	}


	resp := map[string] interface{} {
		"recv" : respData,
	}

	helper.RespJson(ctx, &resp)
}


// 查询链上数据, 返回 map
func queryPostRecvInfoById(postId uint64) (*map[string]interface{}, error) {
	// 获取 ctx 上下文
	clientCtx := client.GetClientContextFromCmd(helper.HttpCmd)

	// 准备查询
	queryClient := posttypes.NewQueryClient(clientCtx)

	params := &posttypes.QueryGetPostRequest{
		Id: postId,
	}

	res, err := queryClient.Post(context.Background(), params)
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

	itemMap := respData["Post"].(map[string]interface{})

	return &(itemMap), nil
}
