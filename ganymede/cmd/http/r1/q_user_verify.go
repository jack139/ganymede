package r1

import (
	"github.com/jack139/ganymede/ganymede/cmd/http/helper"
	ganymedecli "github.com/jack139/ganymede/ganymede/cmd/client"

	"github.com/cosmos/cosmos-sdk/client/flags"

	//"log"
	"strconv"
	"strings"
	"github.com/valyala/fasthttp"
)

/* 用户验证 */

func QueryUserVerify(ctx *fasthttp.RequestCtx) {

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
	mnemonic, ok := (*reqData)["mystery"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need mystery")
		return
	}

	positions, ok := (*reqData)["positions"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need positions")
		return
	}

	// 检查数据合理性
	mnemonic2 := strings.Split(mnemonic, " ")
	pos2 := strings.Split(positions, " ")

	if len(mnemonic2) != len(pos2) {
		helper.RespError(ctx, 9002, "mystery counts NOT equal to position counts")
		return
	}

	if len(pos2) < 3 {
		helper.RespError(ctx, 9002, "the minimum mystery length is 3")
		return		
	}


	/*
	// 验证用户
	verified, err := cmdclient.VerifyUserAccount(helper.HttpCmd, chainAddr, mnemonic)
	if err != nil {
		helper.RespError(ctx, 9009, err.Error())
		return
	}
	*/

	// 获取加密密钥
	flagHome, err := helper.HttpCmd.Flags().GetString(flags.FlagHome)
	if err != nil {
		helper.RespError(ctx, 9021, err.Error())
		return
	}

	smUser, err := ganymedecli.GetSM2Key(flagHome, chainAddr)
	if err != nil {
		helper.RespError(ctx, 9004, "invaild owner: " + err.Error())
		return
	}
	decryptKey := (*(*smUser).CryptoPair.PrivKey)[:16] // 加密密钥


	// 查询链上数据
	userMap, err := queryUserInfoByChainAddr(chainAddr, decryptKey)
	if err!=nil{
		helper.RespError(ctx, 9022, err.Error())
		return
	}

	mnemonic1 := strings.Split((*userMap)["mnemonic"].(string), " ")

	// 转换 position
	var pos2i []int
	for _, v := range pos2 {
		i, err := strconv.Atoi(v)
		if err != nil {
			helper.RespError(ctx, 9002, "position should be integer number")
			return
		}
		if i > len(mnemonic1) || i < 1 {
			helper.RespError(ctx, 9003, "invaild position")
			return			
		}
		pos2i = append(pos2i, i - 1)
	}

	verified := true

	//log.Println(mnemonic1)
	//log.Println(mnemonic2)
	//log.Println(pos2i)

	// 验证用户, 比对指定位置密码串
	for n, p := range pos2i {
		if mnemonic2[n] != mnemonic1[p] {
			verified = false
			break
		}
	}

	// 返回区块id
	resp := map[string]interface{}{
		"verified" : verified,  // 是否验证通过
	}

	helper.RespJson(ctx, &resp)
}
