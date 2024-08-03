package http

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
	"fmt"

	"github.com/jack139/ganymede/ganymede/cmd/http/helper"
	release1 "github.com/jack139/ganymede/ganymede/cmd/http/r1"
)


/* 入口 */
func RunServer() {

	helper.InitSM2()

	/* router */
	r := router.New()
	r.GET("/", index)
	r.POST("/api/test", doNonthing)

	r.POST("/api/r1/tx/user/new", release1.TxUserRegister)
	r.POST("/api/r1/tx/user/update", release1.TxUserUpdate)
	r.POST("/api/r1/tx/user/audit", release1.TxUserAudit)
	r.POST("/api/r1/tx/kv/new", release1.TxZooNew)
	r.POST("/api/r1/tx/kv/update", release1.TxZooUpdate)
	r.POST("/api/r1/tx/kv/delete", release1.TxZooDelete)
	r.POST("/api/r1/tx/exchange/ask", release1.TxExchangeAsk)
	r.POST("/api/r1/tx/exchange/reply", release1.TxExchangeReply)
	r.POST("/api/r1/tx/post/send", release1.TxPostSend)
	r.POST("/api/r1/tx/post/ask", release1.TxPostAsk)
	r.POST("/api/r1/tx/post/reply", release1.TxPostReply)

	r.POST("/api/r1/q/user/info", release1.QueryUserInfo)
	r.POST("/api/r1/q/user/list", release1.QueryUserList)
	r.POST("/api/r1/q/user/verify", release1.QueryUserVerify)
	r.POST("/api/r1/q/bank/balance", release1.QueryBalance)
	r.POST("/api/r1/q/block/height", release1.QueryBlockHeight)
	r.POST("/api/r1/q/block/tx", release1.QueryBlockTx)
	r.POST("/api/r1/q/block/txs", release1.QueryBlockTxs)
	r.POST("/api/r1/q/kv/show", release1.QueryZooInfo)
	r.POST("/api/r1/q/kv/list", release1.QueryZooList)
	r.POST("/api/r1/q/exchange/ask/list", release1.QueryExchangeAskList)
	r.POST("/api/r1/q/exchange/reply/list", release1.QueryExchangeReplyList)
	r.POST("/api/r1/q/exchange/reply/show", release1.QueryExchangeReplyInfo)
	r.POST("/api/r1/q/post/sent/list", release1.QueryPostSentList)
	r.POST("/api/r1/q/post/timeout/list", release1.QueryPostTimeoutList)
	r.POST("/api/r1/q/post/recv/list", release1.QueryPostRecvList)
	r.POST("/api/r1/q/post/recv/show", release1.QueryPostRecvInfo)


	host := fmt.Sprintf("%s:%d", helper.Settings.Api.Addr, helper.Settings.Api.Port)
	log.Printf("start HTTP server at %s\n", host)

	/* 启动server */
	s := &fasthttp.Server{
		Handler: helper.Combined(r.Handler),
		Name:    "FastHttpLogger",
	}
	log.Fatal(s.ListenAndServe(host))
}

/* 根返回 */
func index(ctx *fasthttp.RequestCtx) {
	log.Printf("%v", ctx.RemoteAddr())
	ctx.WriteString("Hello world.")
}
