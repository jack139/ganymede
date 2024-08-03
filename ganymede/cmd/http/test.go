package http

import (
	"github.com/valyala/fasthttp"
	"log"

	"github.com/jack139/ganymede/ganymede/cmd/http/helper"
)

/* 空接口, 只进行签名校验 */
func doNonthing(ctx *fasthttp.RequestCtx) {
	log.Println("doNonthing")

	// POST 的数据
	content := ctx.PostBody()

	// 验签
	data, err := helper.CheckSign(content)
	if err != nil {
		helper.RespError(ctx, 9000, err.Error())
		return
	}
	log.Printf("%v\n", *data)

	helper.RespJson(ctx, data)
}
