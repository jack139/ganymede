// Package provides some helping funcs, suchs as redis-related and settings parsing
package helper

import (
	"log"
	"io/ioutil"
	"gopkg.in/yaml.v3"
)

type apiYaml struct {
	/* http 服务端口和绑定地址 */
	Port int `yaml:"Port"`
	Addr string `yaml:"Addr"`

	/* 接口验签使用 appid : appsecret */
	SECRET_KEY map[string]string `yaml:"AppIdSecret"` 

	/* SM2私钥 */
	SM2Private string `yaml:"SM2PrivateKey"`

	/* api请求timestamp与服务器时间差异(秒)，大于差异绝对值将被拒绝 */
	REQ_TIME_DIFF float64 `yaml:"RequestTimestampDiff"`

	MAXSIZE_KEY int `yaml:"MaxSizeKey"`  // key数据最大尺寸
	MAXSIZE_ASK int `yaml:"MaxSizeAsk"`  // 请求数据最大尺寸
	MAXSIZE_PAYLOAD int `yaml:"MaxSizePayload"` // 数据包最大尺寸
	MAXSIZE_IPFS int `yaml:"MaxSizeIpfs"` // IPFS文件最大尺寸
}

type serverYaml struct {
	IPFS_ENABLE bool `yaml:"IpfsEnable"`
	IPFS_SERVER string `yaml:"IpfsServer"`
	IBC_IPFS_ENABLE bool `yaml:"IBCIpfsEnable"`
	IBC_IPFS_SERVER string `yaml:"IBCIpfsServer"`
}

type chainYaml struct {
	ChainID string `yaml:"ChainID"`
	NodeUser string `yaml:"NodeUser"`
	RelayUser string `yaml:"RelayUser"`
	IBCChannel []string `yaml:"IBCChannel"` /* 可用的IBC channel */
}

type configYaml struct{
	Api apiYaml `yaml:"API"`
	Server serverYaml `yaml:"Server"`
	Chain chainYaml `yaml:"Chain"`
}

// Settings read from local YAML setting file located in 'config/settings.yaml'
var (
	Settings = configYaml{}
)

func ReadSettings(yamlFilepath string){
	config, err := ioutil.ReadFile(yamlFilepath)
	if err != nil {
		log.Fatal("Read settings file FAIL: ", err)
	}

	yaml.Unmarshal(config, &Settings)

	log.Println("Settings loaded: ", yamlFilepath)
}
