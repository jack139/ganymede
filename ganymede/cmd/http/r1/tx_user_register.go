package r1

import (
	cmdclient "github.com/jack139/ganymede/ganymede/cmd/client"
	"github.com/jack139/ganymede/ganymede/cmd/http/helper"
	persontypes "github.com/jack139/ganymede/ganymede/x/ganymede/types"
	ganymedecli "github.com/jack139/ganymede/ganymede/cmd/client"

	"log"
	"time"
	"bytes"
	"strings"
	"encoding/json"
	"encoding/base64"
	"github.com/valyala/fasthttp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/tjfoc/gmsm/sm4"
)

/* 用户注册 */

func TxUserRegister(ctx *fasthttp.RequestCtx) {

	// POST 的数据
	content := ctx.PostBody()

	// 验签
	reqData, err := helper.CheckSign(content)
	if err != nil {
		helper.RespError(ctx, 9000, err.Error())
		return
	}

	// 检查参数
	keyName, ok := (*reqData)["key_name"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need key_name")
		return
	}
	userType, ok := (*reqData)["user_type"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need user_type")
		return
	}
	acc_name, _ := (*reqData)["name"].(string)
	acc_no, _ := (*reqData)["acc_no"].(string)
	contact_address, _ := (*reqData)["address"].(string)
	phone, _ := (*reqData)["phone"].(string)
	refer, _ := (*reqData)["ref"].(string)

	// 检查参数
	keyName = strings.TrimSpace(keyName)
	userType = strings.TrimSpace(userType)

	if !helper.IsAlphanum(keyName) {
		helper.RespError(ctx, 9002, "wrong key_name format")
		return		
	}

	if len(keyName) > helper.Settings.Api.MAXSIZE_KEY {
		helper.RespError(ctx, 9003, "key_name is too long")
		return				
	}
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

	// TODO： 检查 字段大小

	// 初始化用户状态
	userStatus := "WAIT"
	if strings.HasPrefix(userType, "USR") {
		userStatus = "ACTIVE"
	}


	// 生成新用户密钥
	address, mnemonic, serializedPubkey, err := cmdclient.AddUserAccount(helper.HttpCmd, keyName)
	if err != nil {
		helper.RespError(ctx, 9021, err.Error())
		return
	}


	// 构建userInfo
	userInfoMap := map[string]interface{}{
		"acc_name": acc_name,
		"acc_no": acc_no,
		"address": contact_address,
		"phone": phone,
		"mnemonic": mnemonic, // 保存 用于 user verify
	}

	userInfo, err := json.Marshal(userInfoMap)
	if err != nil {
		helper.RespError(ctx, 9022, err.Error())
		return
	}

	// 获取加密密钥
	flagHome, err := helper.HttpCmd.Flags().GetString(flags.FlagHome)
	if err != nil {
		helper.RespError(ctx, 9023, err.Error())
		return
	}

	smUser, err := ganymedecli.GetSM2Key(flagHome, address)
	if err != nil {
		helper.RespError(ctx, 9004, "invaild user: " + err.Error())
		return
	}

	// 加密 user info
	encryptKey := (*(*smUser).CryptoPair.PrivKey)[:16]
	encrypted, err := sm4.Sm4CFB(encryptKey, userInfo, true)
	if err != nil {
		helper.RespError(ctx, 9024, "sm4 encrypt error: " + err.Error())
		return
	}
	userInfoStr := "@@SM4:" + base64.StdEncoding.EncodeToString(encrypted)


	// 设置 caller_addr
	originFlagFrom, err := helper.HttpCmd.Flags().GetString(flags.FlagFrom) // 保存 --from 设置
	if err != nil {
		helper.RespError(ctx, 9026, err.Error())
		return
	}
	//helper.HttpCmd.Flags().Set(flags.FlagFrom, callerAddr)  // 设置 --from 地址
	//defer helper.HttpCmd.Flags().Set(flags.FlagFrom, originFlagFrom)  // 结束时恢复 --from 设置

	/* 信号量 */
	helper.AcquireSem(originFlagFrom)
	defer helper.ReleaseSem(originFlagFrom)

	// 获取 ctx 上下文
	clientCtx, err := client.GetClientTxContext(helper.HttpCmd)
	if err != nil {
		helper.RespError(ctx, 9027, err.Error())
		return
	}

	// 创建者地址，如果在生成新用户后，会变成faucet的地址
	//creatorAddr := clientCtx.GetFromAddress().String()

	// 数据上链
	msg := persontypes.NewMsgCreateUsers(
		originFlagFrom, // creator string,
		address, // chainAddr string,
		keyName, // keyName string,
		userType, // userType string,
		userInfoStr, // name string,
		serializedPubkey, // address string, 借用： 传输 serializedPubkey
		"", // phone string,
		"", // accountNo string,
		refer, // ref string,
		time.Now().Format("2006-01-02 15:04:05"), // regDate string,
		userStatus, // status string,
		time.Now().Format("2006-01-02 15:04:05"), // lastDate string,
		"", // linkStatus string,
	)
	if err := msg.ValidateBasic(); err != nil {
		helper.RespError(ctx, 9028, err.Error())
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
			helper.RespError(ctx, 9029, err.Error())
		}
		return
	}

	// 结果输出
	respBytes := []byte(buf.String())

	//log.Println("output: ", buf.String())

	// 转换成map, 生成返回数据
	var respData map[string]interface{}

	if err := json.Unmarshal(respBytes, &respData); err != nil {
		helper.RespError(ctx, 9030, err.Error())
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
		"chain_addr" : address,  // 用户地址
		"mystery" : mnemonic, // 机密串
	}

	log.Println(resp)

	helper.RespJson(ctx, &resp)
}
