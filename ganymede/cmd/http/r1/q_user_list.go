package r1

import (
	"github.com/jack139/ganymede/ganymede/cmd/http/helper"
	persontypes "github.com/jack139/ganymede/ganymede/x/ganymede/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/query"

	"bytes"
	"context"
	"encoding/json"
	"github.com/valyala/fasthttp"
	//"log"
)



/* 查询用户清单 */
func QueryUserList(ctx *fasthttp.RequestCtx) {

	// POST 的数据
	content := ctx.PostBody()

	// 验签
	reqData, err := helper.CheckSign(content)
	if err != nil {
		helper.RespError(ctx, 9000, err.Error())
		return
	}

	var page, limit float64

	// 检查参数
	page, ok := (*reqData)["page"].(float64)
	if !ok {
		page = 1.0
	}
	limit, ok = (*reqData)["limit"].(float64)
	if !ok {
		limit = 50.0
	}
	status, ok := (*reqData)["status"].(string)


	if page < 1 || limit < 1 {
		helper.RespError(ctx, 9002, "page and limit need begin from 1")
		return		
	}

	// 查询链上数据
	respData2, err := queryUserListPage(uint64(page), uint64(limit), status)
	if err!=nil{
		helper.RespError(ctx, 9021, err.Error())
		return
	}	

	dataList := *respData2

	// 构建返回结构
	respData := make([]map[string]interface{}, 0) 

	for _, item0 := range dataList {
		item := item0.(map[string]interface{})

		newItem := map[string]interface{} {
			"chain_addr" : item["chainAddr"],
			"key_name"   : item["keyName"],
			"user_type"  : item["userType"],
			"reg_date"   : item["regDate"],
			"last_date"  : item["lastDate"],
			"status"     : item["status"],
		}
		respData = append(respData, newItem)
	}

	resp := map[string] interface{} {
		"user_list" : respData,
	}

	helper.RespJson(ctx, &resp)
}


// 查询链上数据, 返回 map
func queryUserListPage(page uint64, limit uint64, status string) (*[]interface{}, error) {
	// 获取 ctx 上下文
	clientCtx := client.GetClientContextFromCmd(helper.HttpCmd)

	// 准备查询
	queryClient := persontypes.NewQueryClient(clientCtx)

	// 设置 接收输出
	buf := new(bytes.Buffer)
	clientCtx.Output = buf
	clientCtx.OutputFormat = "json"

	if len(status)==0 { // 查所有的
		pageReq := query.PageRequest{
			Key:        []byte(""),
			Offset:     (page - 1) * limit,
			Limit:      limit,
			CountTotal: true,
		}

		params := &persontypes.QueryAllUsersRequest{
			Pagination: &pageReq,
		}

		res, err := queryClient.UsersAll(context.Background(), params)
		if err != nil {
			return nil, err
		}

		// 转换输出
		clientCtx.PrintProto(res)
	} else {
		params := &persontypes.QueryListByStatusRequest{
			Page: page,
			Limit: limit,
			Status: status,
		}		

		res, err := queryClient.ListByStatus(context.Background(), params)
		if err != nil {
			return nil, err
		}

		// 转换输出
		clientCtx.PrintProto(res)
	}

	// 输出的字节流
	respBytes := []byte(buf.String())

	//log.Println("output: ", buf.String())

	// 转换成map, 生成返回数据
	var respData map[string]interface{}

	if err := json.Unmarshal(respBytes, &respData); err != nil {
		return nil, err
	}

	userMapList := respData["users"].([]interface{})

	return &(userMapList), nil
}
