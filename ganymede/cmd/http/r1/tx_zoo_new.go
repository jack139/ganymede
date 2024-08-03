package r1

import (
	"github.com/jack139/ganymede/ganymede/cmd/http/helper"
	invtypes "github.com/jack139/ganymede/ganymede/x/zoo/types"
	ganymedecli "github.com/jack139/ganymede/ganymede/cmd/client"
	"github.com/jack139/ganymede/ganymede/cmd/ipfs"

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

/* 新建KV数据 */

func TxZooNew(ctx *fasthttp.RequestCtx) {

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
	zooValue, ok := (*reqData)["value"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need value")
		return
	}
	crypto, ok := (*reqData)["crypto"].(bool)
	if !ok {
		crypto = false
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
	// 检查数据尺寸
	if ok:=ipfs.CheckSize(zooValue); !ok {
		helper.RespError(ctx, 9003, "zooValue is too long")
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
	flagHome, err := helper.HttpCmd.Flags().GetString(flags.FlagHome)
	if err != nil {
		helper.RespError(ctx, 9021, err.Error())
		return
	}

	smUser, err := ganymedecli.GetSM2Key(flagHome, ownerAddr)
	if err != nil {
		helper.RespError(ctx, 9004, "invaild owner: " + err.Error())
		return
	}


	// 获取当前数据， 主要获取 creator
	_, err = queryZooInfoByKey(ownerAddr, zooKey)
	if err==nil{
		helper.RespError(ctx, 9005, "the key already existed!")
		return
	}

	// 是否需要存储到IPFS
	zooValue, err = ipfs.SaveToIpfsIfPossible(zooValue, false)
	if err!=nil {
		helper.RespError(ctx, 9031, "IPFS fail: " + err.Error())
		return
	}

	// 处理数据加密 + base64
	if crypto {
		encryptKey := (*(*smUser).CryptoPair.PrivKey)[:16]
		encrypted, err := sm4.Sm4CFB(encryptKey, []byte(zooValue), true)
		if err != nil {
			helper.RespError(ctx, 9022, "sm4 encrypt error: " + err.Error())
			return
		}
		zooValue = "@@SM4:" + base64.StdEncoding.EncodeToString(encrypted)
	} else {
		zooValue = "PLAIN:" + zooValue
	}

	/* 信号量 */
	helper.AcquireSem(ownerAddr)
	defer helper.ReleaseSem(ownerAddr)

	// 设置 caller_addr
	originFlagFrom, err := helper.HttpCmd.Flags().GetString(flags.FlagFrom) // 保存 --from 设置
	if err != nil {
		helper.RespError(ctx, 9023, err.Error())
		return
	}
	helper.HttpCmd.Flags().Set(flags.FlagFrom, ownerAddr)  // 设置 --from 地址
	defer helper.HttpCmd.Flags().Set(flags.FlagFrom, originFlagFrom)  // 结束时恢复 --from 设置


	// 获取 ctx 上下文
	clientCtx, err := client.GetClientTxContext(helper.HttpCmd)
	if err != nil {
		helper.RespError(ctx, 9024, err.Error())
		return
	}


	// 数据上链
	msg := invtypes.NewMsgCreateKvzoo(
		ownerAddr, //creator string,
		ownerAddr, //owner string,
		zooKey, //zooKey string,
		zooValue, //zooValue string,
		time.Now().Format("2006-01-02 15:04:05"), //lastDate string,
		"", //linkCreator string,
	)
	if err := msg.ValidateBasic(); err != nil {
		helper.RespError(ctx, 9025, err.Error())
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
			helper.RespError(ctx, 9026, err.Error())
		}
		return		
	}

	// 结果输出
	respBytes := []byte(buf.String())

	//log.Println("output: ", buf.String())

	// 转换成map, 生成返回数据
	var respData map[string]interface{}

	if err := json.Unmarshal(respBytes, &respData); err != nil {
		helper.RespError(ctx, 9027, err.Error())
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
