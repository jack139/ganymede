package client

import (
	"math/big"
	"fmt"
	"log"
	"strings"
	"io/ioutil"
	cmn "github.com/cometbft/cometbft/libs/os"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm4"
)

var (
	cdc = AminoCdc
)

type cryptoPair struct {
	PrivKey *[]byte
	PubKey  *[]byte
}

type User struct {
	SignKey    sm2.PrivateKey `json:"sign_key"` // 节点私钥，用户签名
	CryptoPair cryptoPair     // 密钥协商使用
}

// 生成用户环境
func GetSM2Key(path string, addr string) (*User, error) {
	if len(addr) == 0 {
		return nil, fmt.Errorf("wrong address format!")
	}
	keyFilePath := path + "/sm2/" + addr
	if cmn.FileExists(keyFilePath) {
		log.Printf("Found SM2 keyfile: %s\n", keyFilePath)
		uk, err := loadUserKey(keyFilePath, addr)
		if err != nil {
			return nil, err
		}
		return uk, nil
	}

	return nil, fmt.Errorf("SM2 Keyfile does not exist!")
}

// 从文件装入key
func GenSM2Key(path string, addr string, secret string) (*User, error) {
	if len(addr) == 0 {
		return nil, fmt.Errorf("wrong address format!")
	}

	// 建目录
	path = path + "/sm2"
	if err := cmn.EnsureDir(path, 0700); err != nil {
		return nil, err
	}

	keyFilePath := path + "/" + addr
	if cmn.FileExists(keyFilePath) {
		return nil, fmt.Errorf("SM2 Keyfile already exists!")
	}

	// 生成新的密钥文件
	log.Println("Make new SM2 Keyfile: " + keyFilePath)	
	uk := new(User)
	myReader := strings.NewReader(secret)
	signKey, err := sm2.GenerateKey(myReader) // 生成密钥对
	if err != nil {
		return nil, err
	}
	uk.SignKey = *signKey
	pubKey := sm2.Compress(&uk.SignKey.PublicKey) // 33 bytes
	priKey := uk.SignKey.D.Bytes()

	// 加密私钥
	key := []byte(addr)[:16]
	encrypted, err := sm4.Sm4CFB(key, priKey, true)
	if err != nil {
		return nil, fmt.Errorf("sm4 encrypt error: %s", err)
	}

	uk.CryptoPair = cryptoPair{PrivKey: &encrypted, PubKey: &pubKey}
	jsonBytes, err := cdc.MarshalJSON(uk.CryptoPair)
	if err != nil {
		return nil, err
	}
	err = ioutil.WriteFile(keyFilePath, jsonBytes, 0644)
	if err != nil {
		return nil, err
	}

	uk.CryptoPair.PrivKey = &priKey // 恢复私钥

	return uk, nil
}

// 从 byte 恢复密钥对
func RestoreKey(priv *[]byte) *sm2.PrivateKey {
	curve := sm2.P256Sm2()
	key := new(sm2.PrivateKey)
	key.PublicKey.Curve = curve
	key.D = new(big.Int).SetBytes(*priv)
	key.PublicKey.X, key.PublicKey.Y = curve.ScalarBaseMult(*priv)
	return key
}

// 从 byte 恢复公钥
func RestorePublicKey(public []byte) *sm2.PublicKey {
	key := sm2.Decompress(public)
	return key
}

// 从文件导入用户密钥
func loadUserKey(keyFilePath string, addr string) (*User, error) {
	jsonBytes, err := ioutil.ReadFile(keyFilePath)
	if err != nil {
		return nil, err
	}
	uk := new(User)
	err = cdc.UnmarshalJSON(jsonBytes, &uk.CryptoPair)
	if err != nil {
		return nil, fmt.Errorf("Error reading SM2 UserKey from %v: %v", keyFilePath, err)
	}

	// 解密私钥
	key := []byte(addr)[:16]
	plain, err := sm4.Sm4CFB(key, *uk.CryptoPair.PrivKey, false)
	if err != nil {
		return nil, fmt.Errorf("sm4 decrypt error: %s", err)
	}

	uk.CryptoPair.PrivKey = &plain

	// 恢复 privateKey
	uk.SignKey = *RestoreKey(uk.CryptoPair.PrivKey)

	return uk, nil
}

