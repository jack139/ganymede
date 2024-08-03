package exchange

import (
	"github.com/jack139/ganymede/ganymede/cmd/http/helper"
	ganymedecli "github.com/jack139/ganymede/ganymede/cmd/client"

	//"log"
	"encoding/json"
	"encoding/base64"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm4"
)

// 解析响应数据，返回明文
func QueryReplyPayload(askerAddr string, askPayload string, replyPayload string) (bool, string, string, error) {

	// 检查 senderAddr 合法性
	flagHome, err := helper.HttpCmd.Flags().GetString(flags.FlagHome)
	if err != nil {
		return false, "", "", err
	}

	smUser, err := ganymedecli.GetSM2Key(flagHome, askerAddr)
	if err != nil {
		return false, "", "", err
	}

	// 解析 ask数据
	var askData map[string]string

	if err := json.Unmarshal([]byte(askPayload), &askData); err != nil {  // 解析 json
		return false, "", "", err
	}

	cryptAskData, err := base64.StdEncoding.DecodeString(askData["crypt"]) // 解码base64
	if err != nil {
		return false, "", "", err
	}

	// 解析 reply 数据
	var replyData map[string]string

	if err = json.Unmarshal([]byte(replyPayload), &replyData); err != nil {  // 解析 json
		return false, "", "", err
	}

	cryptReplyData, err := base64.StdEncoding.DecodeString(replyData["crypt"]) // 解码base64
	if err != nil {
		return false, "", "", err
	}

	replyPubkeyData, err := base64.StdEncoding.DecodeString(replyData["pubkey"]) // 解码base64
	if err != nil {
		return false, "", "", err
	}


	// 解密出rb, 加密密钥使用私钥前16字节（128bit）
	rbBytesLen := int(cryptAskData[0])
	// 加密数据, rb私钥sm2解密
	rbPrivBytes, err := sm2.DecryptAsn1(&((*smUser).SignKey), cryptAskData[rbBytesLen+1:])
	if err != nil {
		return false, "", "", err
	}
	rbPriv := ganymedecli.RestoreKey(&rbPrivBytes)

	dbBytes := *(*smUser).CryptoPair.PubKey

	// 私钥
	dbPriv := (*smUser).SignKey

	// auth.Data里取得密钥协商数据: ra.pub da.pub data
	raBytesLen := int(cryptReplyData[0])
	daBytes := replyPubkeyData
	raBytes := cryptReplyData[1:raBytesLen+1]

	daPub := ganymedecli.RestorePublicKey(daBytes)
	raPub := ganymedecli.RestorePublicKey(raBytes)

	// 生成 解密密钥
	decryptKey, _, _, err := sm2.KeyExchangeB(16, 
		daBytes, dbBytes, &dbPriv, daPub, rbPriv, raPub)
	if err != nil {
		return false, "", "", err
	}

	// 解密
	decrypted, err := sm4.Sm4CFB(decryptKey, cryptReplyData[raBytesLen+1:], false)
	if err!=nil {
		return false, "", "", err
	} 

	//log.Println(string(decrypted))

	return replyData["reply"]=="PASS", string(decrypted), replyData["uuid"], nil
}
