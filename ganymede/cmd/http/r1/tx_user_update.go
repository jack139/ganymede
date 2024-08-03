package r1

import (
	"github.com/jack139/ganymede/ganymede/cmd/http/helper"
	persontypes "github.com/jack139/ganymede/ganymede/x/ganymede/types"
	ganymedecli "github.com/jack139/ganymede/ganymede/cmd/client"

	"log"
	"bytes"
	"time"
	"strings"
	"encoding/json"
	"encoding/base64"
	"github.com/valyala/fasthttp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/tjfoc/gmsm/sm4"
)

/* 修改用户信息 */

func TxUserUpdate(ctx *fasthttp.RequestCtx) {

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
	acc_name, _ := (*reqData)["name"].(string)
	acc_no, _ := (*reqData)["acc_no"].(string)
	contact_address, _ := (*reqData)["address"].(string)
	phone, _ := (*reqData)["phone"].(string)
	refer, _ := (*reqData)["ref"].(string)

	// 检查参数 长度
	if len(acc_name) > helper.Settings.Api.MAXSIZE_KEY {
		helper.RespError(ctx, 9003, "acc_name is too long")
		return				
	}
	if len(acc_no) > helper.Settings.Api.MAXSIZE_KEY {
		helper.RespError(ctx, 9003, "acc_no is too long")
		return				
	}
	if len(contact_address) > 512 {
		helper.RespError(ctx, 9003, "contact_address is too long")
		return				
	}
	if len(phone) > helper.Settings.Api.MAXSIZE_KEY {
		helper.RespError(ctx, 9003, "phone is too long")
		return				
	}
	if len(refer) > helper.Settings.Api.MAXSIZE_KEY {
		helper.RespError(ctx, 9003, "refer is too long")
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
	encryptKey := (*(*smUser).CryptoPair.PrivKey)[:16] // 加密密钥


	// 获取当前链上数据
	userMap, err := queryUserInfoByChainAddr(chainAddr, encryptKey)
	if err!=nil {
		helper.RespError(ctx, 9022, err.Error())
		return		
	}

	// 用户状态: 除USR外，需要审核
	userStatus := "WAIT"
	if strings.HasPrefix((*userMap)["userType"].(string), "USR") {
		userStatus = "ACTIVE"
	}


	// 是否要修改？
	if len(acc_name)>0 {
		(*userMap)["name"] = acc_name
	}
	if len(acc_no)>0 {
		(*userMap)["accountNo"] = acc_no
	}
	if len(contact_address)>0 {
		(*userMap)["address"] = contact_address
	}
	if len(phone)>0 {
		(*userMap)["phone"] = phone
	}
	if len(refer)>0 {
		(*userMap)["ref"] = refer
	}


	// 构建userInfo
	userInfoMap := map[string]interface{}{
		"acc_name": (*userMap)["name"].(string),
		"acc_no": (*userMap)["accountNo"].(string),
		"address": (*userMap)["address"].(string),
		"phone": (*userMap)["phone"].(string),
		"mnemonic": (*userMap)["mnemonic"].(string),
	}

	userInfo, err := json.Marshal(userInfoMap)
	if err != nil {
		helper.RespError(ctx, 9023, err.Error())
		return
	}

	// 加密 user info
	encrypted, err := sm4.Sm4CFB(encryptKey, userInfo, true)
	if err != nil {
		helper.RespError(ctx, 9024, "sm4 encrypt error: " + err.Error())
		return
	}
	userInfoStr := "@@SM4:" + base64.StdEncoding.EncodeToString(encrypted)


	/* 信号量 */
	helper.AcquireSem((*userMap)["creator"].(string))
	defer helper.ReleaseSem((*userMap)["creator"].(string))


	// 设置 caller_addr
	originFlagFrom, err := helper.HttpCmd.Flags().GetString(flags.FlagFrom) // 保存 --from 设置
	if err != nil {
		helper.RespError(ctx, 9025, err.Error())
		return
	}
	helper.HttpCmd.Flags().Set(flags.FlagFrom, (*userMap)["creator"].(string))  // 设置 --from 地址
	defer helper.HttpCmd.Flags().Set(flags.FlagFrom, originFlagFrom)  // 结束时恢复 --from 设置

	// 获取 ctx 上下文
	clientCtx, err := client.GetClientTxContext(helper.HttpCmd)
	if err != nil {
		helper.RespError(ctx, 9026, err.Error())
		return
	}

	// 数据上链
	msg := persontypes.NewMsgUpdateUsers(
		(*userMap)["creator"].(string), //creator string,
		(*userMap)["chainAddr"].(string), //chainAddr string,
		(*userMap)["keyName"].(string), //keyName string,
		(*userMap)["userType"].(string), //userType string,
		userInfoStr, //name string,
		"", //address string,
		"", //phone string,
		"", //accountNo string,
		(*userMap)["ref"].(string), //ref string,
		(*userMap)["regDate"].(string), //regDate string,
		userStatus, //status string,
		time.Now().Format("2006-01-02 15:04:05"), //lastDate string,
		(*userMap)["linkStatus"].(string), //linkStatus string,
	)
	if err := msg.ValidateBasic(); err != nil {
		helper.RespError(ctx, 9027, err.Error())
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
			helper.RespError(ctx, 9028, err.Error())
		}
		return		
	}

	// 结果输出
	respBytes := []byte(buf.String())

	//log.Println("output: ", buf.String())

	// 转换成map, 生成返回数据
	var respData map[string]interface{}

	if err := json.Unmarshal(respBytes, &respData); err != nil {
		helper.RespError(ctx, 9029, err.Error())
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
