package r1

import (
	"github.com/jack139/ganymede/ganymede/cmd/http/helper"
	invtypes "github.com/jack139/ganymede/ganymede/x/zoo/types"

	"log"
	"bytes"
	"strings"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

/* 删除KV数据 */

func TxZooDelete(ctx *fasthttp.RequestCtx) {

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

	// 检查参数
	zooKey = strings.TrimSpace(zooKey)
	if !helper.IsAlphanum(zooKey) {
		helper.RespError(ctx, 9002, "wrong key format")
		return		
	}

	if len(zooKey) > helper.Settings.Api.MAXSIZE_KEY {
		helper.RespError(ctx, 9003, "zooKey is too long")
		return
	}

	// 检查 ownerAddr 合法性
	userMap, err := queryUserInfoByChainAddr(ownerAddr, nil)
	if err!=nil {
		helper.RespError(ctx, 9004, err.Error())
		return		
	}
	if (*userMap)["status"] != "ACTIVE" {
		helper.RespError(ctx, 9004, "the user is not ACTIVE")
		return				
	}

	// 获取 sm2 密钥


	// 获取当前数据， 主要获取 creator
	zooData, err := queryZooInfoByKey(ownerAddr, zooKey)
	if err!=nil{
		helper.RespError(ctx, 9021, err.Error())
		return
	}


	/* 信号量 */
	helper.AcquireSem((*zooData)["creator"].(string))
	defer helper.ReleaseSem((*zooData)["creator"].(string))


	// 设置 caller_addr
	originFlagFrom, err := helper.HttpCmd.Flags().GetString(flags.FlagFrom) // 保存 --from 设置
	if err != nil {
		helper.RespError(ctx, 9022, err.Error())
		return
	}
	helper.HttpCmd.Flags().Set(flags.FlagFrom, (*zooData)["creator"].(string))  // 设置 --from 地址
	defer helper.HttpCmd.Flags().Set(flags.FlagFrom, originFlagFrom)  // 结束时恢复 --from 设置

	// 获取 ctx 上下文
	clientCtx, err := client.GetClientTxContext(helper.HttpCmd)
	if err != nil {
		helper.RespError(ctx, 9023, err.Error())
		return
	}


	// 数据上链
	msg := invtypes.NewMsgDeleteKvzoo(
		(*zooData)["creator"].(string),
		ownerAddr,
		zooKey,
	)
	if err := msg.ValidateBasic(); err != nil {
		helper.RespError(ctx, 9024, err.Error())
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
			helper.RespError(ctx, 9025, err.Error())
		}
		return		
	}

	// 结果输出
	respBytes := []byte(buf.String())

	//log.Println("output: ", buf.String())

	// 转换成map, 生成返回数据
	var respData map[string]interface{}

	if err := json.Unmarshal(respBytes, &respData); err != nil {
		helper.RespError(ctx, 9026, err.Error())
		return
	}

	// code==0 提交成功
	if respData["code"].(float64)!=0 { 
		helper.RespError(ctx, 9099, buf.String())  ///  提交失败
		return
	}

	// 返回交易hash
	resp := map[string]interface{}{
		"txhash" : respData["txhash"].(string),  // 交易hash
	}

	log.Println(resp)

	helper.RespJson(ctx, &resp)
}
