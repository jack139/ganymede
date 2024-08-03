package r1

import (
	"github.com/jack139/ganymede/ganymede/cmd/http/helper"
	persontypes "github.com/jack139/ganymede/ganymede/x/ganymede/types"

	"log"
	"bytes"
	"time"
	"strings"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

/* 审核用户（修改用户状态） */

func TxUserAudit(ctx *fasthttp.RequestCtx) {

	// POST 的数据
	content := ctx.PostBody()

	// 验签
	reqData, err := helper.CheckSign(content)
	if err != nil {
		helper.RespError(ctx, 9000, err.Error())
		return
	}

	// 检查参数
	chainAddr, ok := (*reqData)["chain_addr"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need chain_addr")
		return
	}
	status, ok := (*reqData)["status"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need status")
		return
	}

	// 获取当前链上数据
	userMap, err := queryUserInfoByChainAddr(chainAddr, nil)
	if err!=nil {
		helper.RespError(ctx, 9021, err.Error())
		return		
	}


	/* 信号量 */
	helper.AcquireSem((*userMap)["creator"].(string))
	defer helper.ReleaseSem((*userMap)["creator"].(string))


	// 设置 caller_addr
	originFlagFrom, err := helper.HttpCmd.Flags().GetString(flags.FlagFrom) // 保存 --from 设置
	if err != nil {
		helper.RespError(ctx, 9022, err.Error())
		return
	}
	helper.HttpCmd.Flags().Set(flags.FlagFrom, (*userMap)["creator"].(string))  // 设置 --from 地址
	defer helper.HttpCmd.Flags().Set(flags.FlagFrom, originFlagFrom)  // 结束时恢复 --from 设置

	// 获取 ctx 上下文
	clientCtx, err := client.GetClientTxContext(helper.HttpCmd)
	if err != nil {
		helper.RespError(ctx, 9023, err.Error())
		return
	}

	// 数据上链
	msg := persontypes.NewMsgUpdateUsers(
		(*userMap)["creator"].(string), //creator string,
		(*userMap)["chainAddr"].(string), //chainAddr string,
		(*userMap)["keyName"].(string), //keyName string,
		(*userMap)["userType"].(string), //userType string,
		(*userMap)["name"].(string), //name string,
		(*userMap)["address"].(string), //address string,
		(*userMap)["phone"].(string), //phone string,
		(*userMap)["accountNo"].(string), //accountNo string,
		(*userMap)["ref"].(string), //ref string,
		(*userMap)["regDate"].(string), //regDate string,
		status, //status string, // 其实 只修改了这里
		time.Now().Format("2006-01-02 15:04:05"), //lastDate string,
		(*userMap)["linkStatus"].(string), //linkStatus string,
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

	// 返回区块id
	resp := map[string]interface{}{
		"txhash" : respData["txhash"].(string),  // 交易hash
	}

	log.Println(resp)

	helper.RespJson(ctx, &resp)
}
