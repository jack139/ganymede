package exchange

import (
	"github.com/jack139/ganymede/ganymede/cmd/http/helper"
	ganymedecli "github.com/jack139/ganymede/ganymede/cmd/client"

	//"log"
	"encoding/json"
	"encoding/base64"
	"crypto/rand"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/tjfoc/gmsm/sm2"
	uuid "github.com/satori/go.uuid"
)

// 生成请求数据
func GenerateAskPayload(senderAddr string, payloadStr string) (string, string, error) {

	// 检查 senderAddr 合法性
	flagHome, err := helper.HttpCmd.Flags().GetString(flags.FlagHome)
	if err != nil {
		return "", "", err
	}

	smUser, err := ganymedecli.GetSM2Key(flagHome, senderAddr)
	if err != nil {
		return "", "", err
	}

	// 生成 密钥交换的数据

	// 生成 rB
	rb, _ := sm2.GenerateKey(rand.Reader) // 生成密钥对
	rbPubBytes := sm2.Compress(&rb.PublicKey) // 33 bytes

	// 用sm2加密rb私钥
	encrypted, err := sm2.EncryptAsn1(&((*smUser).SignKey.PublicKey), rb.D.Bytes(), rand.Reader)
	if err != nil {
		return "", "", err
	}

	// data格式： rb.pub长度(byte) + rb.pub(33bytes?) + sm2加密的rb.priv
	cryptData := append([]byte{byte(len(rbPubBytes))}, rbPubBytes...)
	cryptData = append(cryptData, encrypted...)

	// 生成 payload
	var payloadBytes = make(map[string]string)

	payloadBytes["text"] = payloadStr
	payloadBytes["crypt"] = base64.StdEncoding.EncodeToString(cryptData)
	payloadBytes["pubkey"] = base64.StdEncoding.EncodeToString(*(*smUser).CryptoPair.PubKey)
	payloadBytes["uuid"] = uuid.NewV4().String()

	payload, err := json.Marshal(payloadBytes)
	if err != nil {
		return "", "", err
	}

	//log.Println(string(payload))

	return string(payload), payloadBytes["uuid"], nil
}
