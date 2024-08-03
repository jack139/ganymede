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


/* 获取交易区块数据 */
func getBlockTx(clientCtx client.Context, txhash string) ([]byte, error) {
	if txhash == "" {
		return nil, fmt.Errorf("argument should be a tx hash")
	}

	// If hash is given, then query the tx by hash.
	output, err := authtx.QueryTx(clientCtx, txhash)
	if err != nil {
		return nil, err
	}

	if output.Empty() {
		return nil, fmt.Errorf("no transaction found with hash %s", txhash)
	}

	// 设置 接收输出
	buf := new(bytes.Buffer)
	clientCtx.Output = buf
	clientCtx.OutputFormat = "json"

	// 转换输出
	clientCtx.PrintProto(output)

	return []byte(buf.String()), nil
}


/* 指定区块查询交易 */
func QueryBlockTx(ctx *fasthttp.RequestCtx) {

	// POST 的数据
	content := ctx.PostBody()

	// 验签
	reqData, err := helper.CheckSign(content)
	if err != nil {
		helper.RespError(ctx, 9000, err.Error())
		return
	}

	// 检查参数
	txhash, ok := (*reqData)["txhash"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need height")
		return
	}


	// 获取 ctx 上下文
	clientCtx, err := client.GetClientTxContext(helper.HttpCmd)
	if err != nil {
		helper.RespError(ctx, 9021, err.Error())
		return
	}

	// 准备查询
	respBytes, err := getBlockTx(clientCtx, txhash)
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
		"blcok": respData,
	}

	helper.RespJson(ctx, &resp)
}

