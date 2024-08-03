package client

import (
	"log"
	"errors"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/client"
	//"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bip39 "github.com/cosmos/go-bip39"

	//"github.com/jack139/ganymede/ganymede/x/ganymede/types"
)

// 建新用户 user， 建key，建account
// 返回： address, mnemonic
// 参考： cosmos-sdk-0.46.11/client/keys/add.go
func AddUserAccount(cmd *cobra.Command, name string) (string, string, string, error) {
	// 节点home
	flagHome, err := cmd.Flags().GetString(flags.FlagHome)
	if err != nil {
		return "", "", "", err
	}

	// context
	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return "", "", "", err
	}
	kb := clientCtx.Keyring

	// 注册新的 key
	keyringAlgos, _ := kb.SupportedAlgorithms()
	algo, err := keyring.NewSigningAlgoFromString(string(hd.Secp256k1Type), keyringAlgos)
	if err != nil {
		return "", "", "", err
	}

	hdPath := hd.CreateHDPath(sdk.GetConfig().GetCoinType(), 0, 0).String()

	_, err = kb.Key(name)
	if err == nil {
		return "", "", "", errors.New("user name already existed!")
	}

	// read entropy seed straight from tmcrypto.Rand and convert to mnemonic
	mnemonicEntropySize := 256
	entropySeed, err := bip39.NewEntropy(mnemonicEntropySize)
	if err != nil {
		return "", "", "", err
	}

	// Get bip39 mnemonic
	var mnemonic, bip39Passphrase string

	mnemonic, err = bip39.NewMnemonic(entropySeed)
	if err != nil {
		return "", "", "", err
	}

	info, err := kb.NewAccount(name, mnemonic, bip39Passphrase, hdPath, algo)
	if err != nil {
		return "", "", "", err
	}

	pk, err := info.GetPubKey() // secp256k1
	if err != nil {
		return "", "", "", err
	}

	// 序列化， 消息参数中只传递 PubKey
	aminoCdc := codec.NewLegacyAmino()
	serializedPubkey, err := aminoCdc.MarshalJSON(pk)
	if err != nil {
		return "", "", "", err
	}


	// 取得地址字符串： 例如 artchain1zfqgxtujvpy92prtzgmzs3ygta9y2cl3w8hxlh
	ko_new, err := keyring.MkAccKeyOutput(info)
	if err != nil {
		return "", "", "", err
	}
	log.Println(ko_new)

	addr_new := ko_new.Address


	// 生成sm2密钥， 由mnemonic生成
	_, err = GenSM2Key(flagHome, addr_new, mnemonic)
	if err != nil {
		return "", "", "", err
	}

	/*
	_, err = GetSM2Key(flagHome, addr_new)
	if err != nil {
		return "", "", err
	}
	*/

	/* msg_users 实现 msg_add_account 的功能，因此这里不调用 add_account 的 tx -- 2023-05-09
	// ganymede addAccount 消息
	msg := types.NewMsgAddAccount(clientCtx.GetFromAddress().String(), string(serializedPubkey))
	if err := msg.ValidateBasic(); err != nil {
		return "", "", err
	}

	clientCtx.BroadcastMode = flags.BroadcastBlock // 默认是 flags.BroadcastSync
	// block 方式将废弃， 使用 flags.BroadcastSync 可能会报错（两次tx太近时？）：  account sequence mismatch
	// 见： https://github.com/cosmos/cosmos-sdk/issues/13621

	// 调用 RPC 服务 添加 auth.account
	return addr_new, mnemonic, tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
	*/

	return addr_new, mnemonic, string(serializedPubkey), nil
}


// 比对用户 user， 
// 返回： bool
func VerifyUserAccount(cmd *cobra.Command, userAddr string, mnemonic string) (bool, error) {
	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return false, err
	}
	kb := clientCtx.Keyring

	// 注册新的 key
	keyringAlgos, _ := kb.SupportedAlgorithms()
	algo, err := keyring.NewSigningAlgoFromString(string(hd.Secp256k1Type), keyringAlgos)
	if err != nil {
		return false, err
	}

	hdPath := hd.CreateHDPath(sdk.GetConfig().GetCoinType(), 0, 0).String()

	// Get bip39 mnemonic
	var bip39Passphrase string

	if !bip39.IsMnemonicValid(mnemonic) && mnemonic != "" {
		return false, errors.New("invalid mnemonic")
	}

	// 生成私钥
	derivedPriv, err := algo.Derive()(mnemonic, bip39Passphrase, hdPath)
	if err != nil {
		return false, err
	}
	privKey := algo.Generate()(derivedPriv)

	// 从公钥生成acc地址
	accAddr := sdk.AccAddress(privKey.PubKey().Address().Bytes())
	//log.Println(accAddr.String())

	return accAddr.String()==userAddr, nil
}


/* 通过key name获取地址 */
func GetAddrStr(cmd *cobra.Command, keyref string) (string, error) {
	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return "", err
	}
	kb := clientCtx.Keyring

	// 获取 地址
	//keyref := "faucet"
	info0, err := kb.Key(keyref)
	if err != nil {
		return "", err
	}
	//addr0 := info0.GetAddress() // AccAddress

	ko, err := keyring.MkAccKeyOutput(info0)
	if err != nil {
		return "", err
	}

	// 取得地址字符串： 例如 artchain1zfqgxtujvpy92prtzgmzs3ygta9y2cl3w8hxlh
	addr0 := ko.Address

	return addr0, nil
}
