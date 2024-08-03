package r1

import (
	"github.com/jack139/ganymede/ganymede/cmd/http/helper"
	invtypes "github.com/jack139/ganymede/ganymede/x/zoo/types"
	ganymedecli "github.com/jack139/ganymede/ganymede/cmd/client"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/client/flags"

	//"log"
	"bytes"
	"context"
	//"strings"
	"encoding/json"
	//"encoding/base64"
	"github.com/valyala/fasthttp"
	//"github.com/tjfoc/gmsm/sm4"
)



/* 查询KV数据清单 */
func QueryZooList(ctx *fasthttp.RequestCtx) {

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
	ownerAddr, ok := (*reqData)["owner_addr"].(string)

	if page < 1 || limit < 1 {
		helper.RespError(ctx, 9002, "page and limit need begin from 1")
		return		
	}

	// 查询链上数据
	respData2, err := queryZooListPage(uint64(page), uint64(limit), ownerAddr)
	if err!=nil{
		helper.RespError(ctx, 9021, err.Error())
		return
	}
	dataList := *respData2


	// 获取 home 路径
	flagHome, err := helper.HttpCmd.Flags().GetString(flags.FlagHome)
	if err != nil {
		helper.RespError(ctx, 9022, err.Error())
		return
	}

	// 密钥缓存
	decryptKeyMap := make(map[string][]byte, 0) 

	// 构建返回结构
	respData := make([]map[string]interface{}, 0) 

	for _, item0 := range dataList {
		item := item0.(map[string]interface{})

		owner := item["owner"].(string)
		decryptKey, ok := decryptKeyMap[owner]
		if !ok {
			// 检查 ownerAddr 合法性
			smUser, err := ganymedecli.GetSM2Key(flagHome, owner)
			if err != nil {
				helper.RespError(ctx, 9004, "invaild owner: " + err.Error())
				return
			}
			decryptKey = (*(*smUser).CryptoPair.PrivKey)[:16]
			decryptKeyMap[owner] = decryptKey
		} else {
			//log.Println("bingo!", owner)
		}

		/* 在 list 里不展示 value 2023-05-19
		// 处理数据解密, 先 base64
		zooValue := item["zooValue"].(string)
		//log.Println(item["zooKey"], zooValue)
		if strings.HasPrefix(zooValue, "@@SM4:") {
			zooValue = zooValue[6:] // 如果解码出错，则会直接返回
			if len(zooValue)>0 {
				// 先解码base64， 再解密
				encrypted, err := base64.StdEncoding.DecodeString(zooValue)
				if err == nil {
					decrypted, err := sm4.Sm4CFB(decryptKey, encrypted, false)
					if err == nil {
						zooValue = string(decrypted)
					}
				}
			}
		} else if strings.HasPrefix(zooValue, "PLAIN:") {
			zooValue = zooValue[6:]
		}
		*/

		newItem := map[string]interface{} {
			"owner_addr" : item["owner"],
			"key"        : item["zooKey"],
			//"value"      : zooValue,
			"last_date"  : item["lastDate"],
		}
		respData = append(respData, newItem)
	}

	resp := map[string] interface{} {
		"kv_list" : respData,
	}

	helper.RespJson(ctx, &resp)
}


// 查询链上数据, 返回 map
func queryZooListPage(page uint64, limit uint64, ownerAddr string) (*[]interface{}, error) {
	// 获取 ctx 上下文
	clientCtx := client.GetClientContextFromCmd(helper.HttpCmd)

	// 准备查询
	queryClient := invtypes.NewQueryClient(clientCtx)


	// 设置 接收输出
	buf := new(bytes.Buffer)
	clientCtx.Output = buf
	clientCtx.OutputFormat = "json"

	if len(ownerAddr)==0 { // 查所有的
		pageReq := query.PageRequest{
			Key:        []byte(""),
			Offset:     (page - 1) * limit,
			Limit:      limit,
			CountTotal: true,
		}

		params := &invtypes.QueryAllKvzooRequest{
			Pagination: &pageReq,
		}

		res, err := queryClient.KvzooAll(context.Background(), params)
		if err != nil {
			return nil, err
		}

		// 转换输出
		clientCtx.PrintProto(res)

	} else { // 查指定owner_addr的
		params := &invtypes.QueryListByOwnerRequest{
			Page: page,
			Limit: limit,
			Owner: ownerAddr,
		}		

		res, err := queryClient.ListByOwner(context.Background(), params)
		if err != nil {
			return nil, err
		}

		// 转换输出
		clientCtx.PrintProto(res)
	}

	//log.Printf("%T\n", res)

	// 输出的字节流
	respBytes := []byte(buf.String())

	//log.Println("output: ", buf.String())

	// 转换成map, 生成返回数据
	var respData map[string]interface{}

	if err := json.Unmarshal(respBytes, &respData); err != nil {
		return nil, err
	}

	itemMapList := respData["kvzoo"].([]interface{})

	return &(itemMapList), nil
}


