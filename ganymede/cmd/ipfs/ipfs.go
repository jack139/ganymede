package ipfs

import (
	"fmt"
	"io"
	"log"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/jack139/ganymede/ganymede/cmd/http/helper"
)

func ipfsAdd(filedata []byte, server string) (string, error) {
	if !helper.Settings.Server.IPFS_ENABLE {
		return "", fmt.Errorf("IPFS not available!")
	}

	// 连接api
	sh := shell.NewShell(server)

	// 添加内容
	cid, err := sh.Add(strings.NewReader(string(filedata)))
	if err != nil {
		return "", fmt.Errorf("IPFS error: %s", err)
	}

	log.Printf("new cid %s", cid)

	return cid, nil
}

func ipfsGet(cid string, server string) ([]byte, error) {
	if !helper.Settings.Server.IPFS_ENABLE {
		return nil, fmt.Errorf("IPFS not available!")
	}

	// 连接api
	sh := shell.NewShell(server)

	// 获取文件内容
	data, err := sh.Cat(cid)
	if err != nil {
		return nil, fmt.Errorf("IPFS error: %s", err)
	}
	defer data.Close()

	// 使用缓存读出文件
	var dataBuf []byte
	longBuf := make([]byte, 1024*20)

	for {
		sz, err := data.Read(longBuf)
		if err != nil {
			if err == io.EOF {
				if sz > 0 { // EOF 此时有可能还读出了数据
					log.Printf("EOF: n = %d", sz)
					dataBuf = append(dataBuf, longBuf[:sz]...)
				}
				break
			}
			return nil, fmt.Errorf("IPFS error: %s", err)
		}
		//fmt.Printf("%d %s\n", sz, longBuf)
		dataBuf = append(dataBuf, longBuf[:sz]...)
	}

	log.Printf("get cid %s", cid)

	return dataBuf, nil
}


// 检查数据长度是否超过限制
func CheckSize(data string) (bool) {
	if len(data) > helper.Settings.Api.MAXSIZE_PAYLOAD {
		if !helper.Settings.Server.IPFS_ENABLE { // 无 IPFS
			return false
		}
		if len(data) > helper.Settings.Api.MAXSIZE_IPFS { // 超过IPFS文件尺寸限制
			return false
		}
	}

	return true
}


// 数据存储到 IPFS, 如果可能
func SaveToIpfsIfPossible(data string, ibc bool) (string, error) {
	if len(data) > helper.Settings.Api.MAXSIZE_PAYLOAD {
		var server string
		if ibc {
			server = helper.Settings.Server.IBC_IPFS_SERVER
		} else {
			server = helper.Settings.Server.IPFS_SERVER
		}

		// 能这里 已经检查 IPFS_ENABLE==ture了
		cid, err := ipfsAdd([]byte(data), server)
		if err!=nil {
			return "", err
		}
		data = "IPFS:" + cid
	} else {
		data = "TEXT:" + data
	}

	return data, nil
}


// 从IPFS取回数据，如果可能
func GetFromIpfsIfPossible(data string, ibc bool) (string, error) {
	if strings.HasPrefix(data, "IPFS:") {
		var server string
		if ibc {
			server = helper.Settings.Server.IBC_IPFS_SERVER
		} else {
			server = helper.Settings.Server.IPFS_SERVER
		}

		ipfsData, err := ipfsGet(data[5:], server)
		if err!=nil {
			return "", err
		}
		data = string(ipfsData)
	} else if strings.HasPrefix(data, "TEXT:") {
		data = data[5:]
	}

	return data, nil
}
