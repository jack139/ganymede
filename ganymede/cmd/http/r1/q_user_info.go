package r1

import (
	"github.com/jack139/ganymede/ganymede/cmd/http/helper"
	persontypes "github.com/jack139/ganymede/ganymede/x/ganymede/types"
	ganymedecli "github.com/jack139/ganymede/ganymede/cmd/client"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"fmt"
	"bytes"
	//"log"
	"strings"
	"context"
	"encoding/json"
	"encoding/base64"
	"github.com/valyala/fasthttp"
	"github.com/tjfoc/gmsm/sm4"
)



/* 查询用户信息 */
func QueryUserInfo(ctx *fasthttp.RequestCtx) {

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
	respData2, err := queryUserInfoByChainAddr(chainAddr, decryptKey)
	if err!=nil{
		helper.RespError(ctx, 9022, err.Error())
		return
	}	

	userMap := *respData2

	// 构建返回结构
	respData := map[string] interface{} {
		"chain_addr"    : userMap["chainAddr"],
		"key_name"      : userMap["keyName"],
		"user_type"     : userMap["userType"],
		"name"          : userMap["name"],
		"acc_no"        : userMap["accountNo"],
		"address"       : userMap["address"],
		"phone"         : userMap["phone"],
		"ref"           : userMap["ref"],
		"reg_date"      : userMap["regDate"],
		"last_date"     : userMap["lastDate"],
		"status"        : userMap["status"],
	}

	resp := map[string] interface{} {
		"user" : respData,
	}

	helper.RespJson(ctx, &resp)
}


// 查询链上数据, 返回 User map
func queryUserInfoByChainAddr(chainAddr string, decryptKey []byte) (*map[string]interface{}, error) {
	// 获取 ctx 上下文
	clientCtx := client.GetClientContextFromCmd(helper.HttpCmd)

	// 检查 用户地址 是否存在
	_, err := helper.FetchKey(clientCtx.Keyring, chainAddr)
	if err != nil {
		return nil, fmt.Errorf("invalid chain_addr")
	}

	// 准备查询
	queryClient := persontypes.NewQueryClient(clientCtx)

	params := &persontypes.QueryGetUsersRequest{
		ChainAddr: chainAddr,
	}

	res, err := queryClient.Users(context.Background(), params)
	if err != nil {
		return nil, err
	}

	//log.Printf("%t\n", res)

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

	// 处理data字段
	respData2, err := unmarshalUser(&respData, decryptKey)
	if err!=nil{
		return nil, err
	}

	return respData2, nil
}


/* userInfo字段是已序列化的json串，反序列化一下，针对数据列表 */
func unmarshalUser(reqData *map[string]interface{}, decryptKey []byte) (*map[string]interface{}, error) {
	item := (*reqData)["users"].(map[string]interface{})

	if decryptKey==nil { // 不解码 user info
		return &item, nil
	}

	// 检查 userInfo 字段是否正常, 放在 name
	if _, ok := item["name"]; !ok {
		return nil, fmt.Errorf("userInfo empty") // 不应该发生
	}
	if !strings.HasPrefix(item["name"].(string), "@@SM4:") {
		return &item, nil  // 旧数据，没有做加密
	}

	// 先解码base64， 再解密
	encrypted, err := base64.StdEncoding.DecodeString((item["name"].(string))[6:])
	if err != nil {
		return nil, err
	}
	decrypted, err := sm4.Sm4CFB(decryptKey, encrypted, false)
	if err != nil {
		return nil, err
	}

	// 反序列化
	var data map[string]interface{}
	if err := json.Unmarshal(decrypted, &data); err != nil {
		return nil, err
	}

	//log.Println("user info: ", data)

	item["name"] = data["acc_name"].(string)
	item["accountNo"] = data["acc_no"].(string)
	item["address"] = data["address"].(string)
	item["phone"] = data["phone"].(string)
	_, ok:= data["mnemonic"]  // 旧数据没保存次字段
	if ok {
		item["mnemonic"] = data["mnemonic"].(string)
	} else {
		item["mnemonic"] = ""
	}

	return &item, nil
}

