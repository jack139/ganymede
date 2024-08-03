package exchange

import (
	"github.com/jack139/ganymede/ganymede/cmd/http/helper"
	ganymedecli "github.com/jack139/ganymede/ganymede/cmd/client"

	"log"
	"encoding/json"
	"encoding/base64"
	"crypto/rand"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm4"
)

// 生成响应数据数据
func GenerateReplyPayload(passReply bool, replierAddr string, askPayload string, 
	payloadStr string) (string, string, error) {

	// 检查 senderAddr 合法性
	flagHome, err := helper.HttpCmd.Flags().GetString(flags.FlagHome)
	if err != nil {
		return "", "", err
	}

	smUser, err := ganymedecli.GetSM2Key(flagHome, replierAddr)
	if err != nil {
		return "", "", err
	}

	// 解析 ask数据
	var askData map[string]string

	if err := json.Unmarshal([]byte(askPayload), &askData); err != nil {  // 解析 json
		return "", "", err
	}

	cryptAskData, err := base64.StdEncoding.DecodeString(askData["crypt"]) // 解码base64
	if err != nil {
		return "", "", err
	}

	askPubkeyData, err := base64.StdEncoding.DecodeString(askData["pubkey"]) // 解码base64
	if err != nil {
		return "", "", err
	}


	// 从 ask请求 里，获取密钥交换的数据
	rbBytesLen := int(cryptAskData[0])
	dbBytes := askPubkeyData // sender 的 pubkey
	rbBytes := cryptAskData[1:rbBytesLen+1]

	dbPub := ganymedecli.RestorePublicKey(dbBytes)
	rbPub := ganymedecli.RestorePublicKey(rbBytes)

	// da 就是自己的密钥, replier
	daBytes := *(*smUser).CryptoPair.PubKey
	daPriv := (*smUser).SignKey
	// 生成 ra
	raPriv, _ := sm2.GenerateKey(rand.Reader) // 生成密钥对


	//  生成 key
	encryptKey, _, _, err := sm2.KeyExchangeA(16, 
		daBytes, dbBytes, &daPriv, dbPub, raPriv, rbPub)
	if err != nil {
		return "", "", err
	}
	log.Println("exchange key:", encryptKey)

	// 用协商密钥加密
	encrypted, err := sm4.Sm4CFB(encryptKey, []byte(payloadStr), true)
	if err != nil {
		return "", "", err
	}
	//log.Printf("%d %v\n", len(encrypted), encrypted)


	raPubBytes := sm2.Compress(&raPriv.PublicKey) // 33 bytes

	// data格式： ra.pub长度(byte) + ra.pub(33bytes?) + K加密的数据
	cryptData := append([]byte{byte(len(raPubBytes))}, raPubBytes...)
	cryptData = append(cryptData, encrypted...)


	// 生成 payload
	var payloadBytes = make(map[string]string)

	payloadBytes["crypt"] = base64.StdEncoding.EncodeToString(cryptData)
	payloadBytes["pubkey"] = base64.StdEncoding.EncodeToString(*(*smUser).CryptoPair.PubKey)
	payloadBytes["uuid"] = askData["uuid"]
	if passReply { // 授权/拒绝 标记
		payloadBytes["reply"] = "PASS"
	} else {
		payloadBytes["reply"] = "DENY"
	}

	payload, err := json.Marshal(payloadBytes)
	if err != nil {
		return "", "", err
	}

	//log.Println(string(payload))

	return string(payload), payloadBytes["uuid"], nil
}
