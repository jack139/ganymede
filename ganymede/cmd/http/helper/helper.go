package helper

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"
	"os"
	"context"
	"math"
	"regexp"
	"strings"

	"github.com/valyala/fasthttp"
	"github.com/Ferluci/fast-realip"
	"github.com/spf13/cobra"
	"golang.org/x/sync/semaphore"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
    Limit  = 1 // 同時并行运行的goroutine上限
    Weight = 1 // 每个goroutine获取信号量资源的权重
)

var (
	/* 保存命令行context */
	HttpCmd *cobra.Command

	/* 日志输出使用 */
	output  = log.New(os.Stdout, "", 0)

	/* 返回值的 content-type */
	strContentType     = []byte("Content-Type")
	strApplicationJSON = []byte("application/json")

	/* 信号量 */
	sem = map[string]*semaphore.Weighted{}
)

/* 获取信号量 */
func AcquireSem(creator string) {
	log.Println("Acquire S ...", creator)
	if _, ok := sem[creator]; !ok {
		sem[creator] = semaphore.NewWeighted(Limit)
		log.Println("New S ...", creator, " Length: ", len(sem))
	}
	sem[creator].Acquire(context.Background(), Weight)
	log.Println("Got S ...", creator)
}

/* 释放信号量 */
func ReleaseSem(creator string) {
	log.Println("Release semaphore ...", creator)
	sem[creator].Release(Weight)
}

/* 处理返回值，返回json */
func RespJson(ctx *fasthttp.RequestCtx, data *map[string]interface{}) {
	respJson := map[string]interface{}{
		"code": 0,
		"msg":  "success",
		"data": *data,
	}
	doJSONWrite(ctx, fasthttp.StatusOK, respJson)
}

func RespError(ctx *fasthttp.RequestCtx, code int, msg string) {
	log.Println("Error: ", code, msg)
	respJson := map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": "",
	}
	doJSONWrite(ctx, fasthttp.StatusOK, respJson)
}

func doJSONWrite(ctx *fasthttp.RequestCtx, code int, obj interface{}) {
	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
	ctx.Response.SetStatusCode(code)
	start := time.Now()
	if err := json.NewEncoder(ctx).Encode(obj); err != nil {
		elapsed := time.Since(start)
		log.Printf("", elapsed, err.Error(), obj)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}

/*
	接口验签，返回data数据
*/
func CheckSign(content []byte) (*map[string]interface{}, error) {
	fields := make(map[string]interface{})
	if err := json.Unmarshal(content, &fields); err != nil {
		return nil, err
	}

	// 参数输出到日志
	log.Println(fields)

	var appId, version, signType, signData string
	var timestamp int64
	var data map[string]interface{}
	var ok bool

	// 检查参数
	if appId, ok = fields["appid"].(string); !ok {
		return nil, fmt.Errorf("need appid")
	}
	if version, ok = fields["version"].(string); !ok {
		return nil, fmt.Errorf("need version")
	}
	if signType, ok = fields["sign_type"].(string); !ok {
		return nil, fmt.Errorf("need sign_type")
	}
	if signData, ok = fields["sign_data"].(string); !ok {
		return nil, fmt.Errorf("need sign_data")
	}
	if _, ok = fields["timestamp"].(float64); !ok {
		return nil, fmt.Errorf("need timestamp")
	} else {
		timestamp = int64(fields["timestamp"].(float64)) // 返回整数
	}
	if data, ok = fields["data"].(map[string]interface{}); !ok {
		return nil, fmt.Errorf("need data")
	}

	// 调用时间不能超过前后5分钟
	if math.Abs(float64(time.Now().Unix()-timestamp)) > Settings.Api.REQ_TIME_DIFF {
		return nil, fmt.Errorf("timestamp is too old")
	}

	// 获取 secret，用户密钥的签名串
	secret, ok := Settings.Api.SECRET_KEY[appId]
	if !ok {
		return nil, fmt.Errorf("wrong appId")
	}

	// 检查版本
	if version != "1" {
		return nil, fmt.Errorf("wrong version")
	}

	/*
	// 检查签名类型
	if signType != "SHA256" {
		return nil, fmt.Errorf("unknown signType")
	}
	*/

	// 生成参数的key，并排序
	keys := getMapKeys(fields)
	sort.Strings(*keys)
	//fmt.Println(*keys)

	// data 串，用于验签， map已按key排序
	dataStr, _ := json.Marshal(data)

	// 拼接验签串
	var signString = string("")
	for _, k := range *keys {
		if k == "sign_data" {
			continue
		}
		if k == "data" {
			signString += k + "=" + string(dataStr) + "&"
		} else if k == "timestamp" {
			signString += k + "=" + strconv.FormatInt(timestamp, 10) + "&"
		} else {
			signString += k + "=" + fields[k].(string) + "&"
		}
	}
	signString += "key=" + secret
	//fmt.Println(signString)

	// 验签
	log.Println("signature type: ", signType)
	switch signType {
	case "SHA256":
		h := sha256.New()
		h.Write([]byte(signString))
		sum := h.Sum(nil)
		sha256Str := fmt.Sprintf("%x", sum)
		signStr := base64.StdEncoding.EncodeToString([]byte(sha256Str))
		//fmt.Println(sha256Str)

		if signStr != signData {
			fmt.Println(signStr)
			fmt.Println(signData)
			return nil, fmt.Errorf("wrong SHA256 signature")
		}
	case "SM2":
		ok := sm2VerifyBase64([]byte(signString), signData)
		if ok != true {
			fmt.Println(signString)
			fmt.Println(signData)
			return nil, fmt.Errorf("wrong SM2 signature")
		}
	default: // 未知签名算法
		return nil, fmt.Errorf("unknown signType")
	}


	return &data, nil
}


// 日志格式处理

// "github.com/AubSs/fasthttplogger"
func getHttp(ctx *fasthttp.RequestCtx) string {
	if ctx.Response.Header.IsHTTP11() {
		return "HTTP/1.1"
	}
	return "HTTP/1.0"
}

// Combined format:
// [<time>] <remote-addr> | <HTTP/http-version> | <method> <url> - <status> - <response-time us> | <user-agent>
// [2017/05/31 - 13:27:28] 127.0.0.1:54082 | HTTP/1.1 | GET /hello - 200 - 48.279µs | Paw/3.1.1 (Macintosh; OS X/10.12.5) GCDHTTPRequest
func Combined(req fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		begin := time.Now()
		req(ctx)
		end := time.Now()
		output.Printf("[%v] %v (%v) | %s | %s %s - %v - %v | %s",
			end.Format("2006/01/02 - 15:04:05"),
			ctx.RemoteAddr(),
			realip.FromRequest(ctx),
			getHttp(ctx),
			ctx.Method(),
			ctx.RequestURI(),
			ctx.Response.Header.StatusCode(),
			end.Sub(begin),
			ctx.UserAgent(),
		)
	})
}


// 返回 map 所有 key
func getMapKeys(m map[string]interface{}) *[]string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	return &keys
}

/* 获取key信息 */
// TODO: 这个函数只能检查本地存储的 keyring，如果验证用户地址需要是用 链上user数据
func FetchKey(kb keyring.Keyring, keyref string) (*keyring.Record, error) {
	info, err := kb.Key(keyref)
	if err != nil {
		accAddr, err := sdk.AccAddressFromBech32(keyref)
		if err != nil {
			return info, err
		}

		info, err = kb.KeyByAddress(accAddr)
		if err != nil {
			return info, errors.New("key not found")
		}
	}
	return info, nil
}


func IsAlphanum(word string) bool {
    return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(word)
}


// 封装 GenerateOrBroadcastTxCLI， 延时 retry 5次, 约10秒
func BroadcastTxWithRetry(clientCtx client.Context, msgs sdk.Msg) error {
	retry := 0
	for {
		err := tx.GenerateOrBroadcastTxCLI(clientCtx, HttpCmd.Flags(), msgs)
		if err==nil {
			break
		}

		if strings.Contains(err.Error(), "account sequence mismatch") {
			if retry > 4 {
				return err
			}
			log.Println("Tx try again: ", retry+1)
			time.Sleep(2000 * time.Millisecond) // 延时 2 秒
			retry++
		} else {
			return err
		}
	}

	return nil
}