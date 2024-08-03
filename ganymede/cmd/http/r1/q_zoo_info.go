package r1

import (
	"github.com/jack139/ganymede/ganymede/cmd/http/helper"
	invtypes "github.com/jack139/ganymede/ganymede/x/zoo/types"
	ganymedecli "github.com/jack139/ganymede/ganymede/cmd/client"
	"github.com/jack139/ganymede/ganymede/cmd/ipfs"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"strings"
	"bytes"
	"context"
	//"log"
	"encoding/json"
	"encoding/base64"
	"github.com/valyala/fasthttp"
	"github.com/tjfoc/gmsm/sm4"
)



/* 查询KV数据信息 */
func QueryZooInfo(ctx *fasthttp.RequestCtx) {

	// POST 的数据
	content := ctx.PostBody()

	// 验签
	reqData, err := helper.CheckSign(content)
	if err != nil {
		helper.RespError(ctx, 9000, err.Error())
		return
	}

	// 检查参数
	ownerAddr, ok := (*reqData)["owner_addr"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need owner_addr")
		return
	}
	zooKey, ok := (*reqData)["key"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need key")
		return
	}


	// 查询链上数据
	respData2, err := queryZooInfoByKey(ownerAddr, zooKey)
	if err!=nil{
		helper.RespError(ctx, 9021, err.Error())
		return
	}	
	itemMap := *respData2

	// 检查 ownerAddr 合法性
	flagHome, err := helper.HttpCmd.Flags().GetString(flags.FlagHome)
	if err != nil {
		helper.RespError(ctx, 9022, err.Error())
		return
	}

	smUser, err := ganymedecli.GetSM2Key(flagHome, ownerAddr)
	if err != nil {
		helper.RespError(ctx, 9004, "invaild owner: " + err.Error())
		return
	}

	// 处理数据解密, 先 base64
	zooValue := itemMap["zooValue"].(string)
	if strings.HasPrefix(zooValue, "@@SM4:") {
		decryptKey := (*(*smUser).CryptoPair.PrivKey)[:16]
		zooValue = zooValue[6:] // 如果解码出错，则会直接返回
		// 先解码base64， 再解密
		encrypted, err := base64.StdEncoding.DecodeString(zooValue)
		if err == nil {
			decrypted, err := sm4.Sm4CFB(decryptKey, encrypted, false)
			if err == nil {
				zooValue = string(decrypted)
			}
		}
	} else if strings.HasPrefix(zooValue, "PLAIN:") {
		zooValue = zooValue[6:]
	}

	// 是否需要从IPFS取回
	zooValue, err = ipfs.GetFromIpfsIfPossible(zooValue, false)
	if err!=nil {
		helper.RespError(ctx, 9032, "IPFS fail:" + err.Error())
		return
	}

	// 构建返回结构
	respData := map[string]interface{} {
		"owner_addr" : itemMap["owner"],
		"key"       : itemMap["zooKey"],
		"value"     : zooValue,
		"last_date" : itemMap["lastDate"],
	}

	resp := map[string] interface{} {
		"kv" : respData,
	}

	helper.RespJson(ctx, &resp)
}


// 查询链上数据, 返回 map
func queryZooInfoByKey(ownerAddr string, zooKey string) (*map[string]interface{}, error) {
	// 获取 ctx 上下文
	clientCtx := client.GetClientContextFromCmd(helper.HttpCmd)

	// 准备查询
	queryClient := invtypes.NewQueryClient(clientCtx)

	params := &invtypes.QueryGetKvzooRequest{
		Owner: ownerAddr,
		ZooKey: zooKey,
	}

	res, err := queryClient.Kvzoo(context.Background(), params)
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

	itemMap := respData["kvzoo"].(map[string]interface{})

	return &(itemMap), nil
}
