package wechat

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"umdp/app/manage/biz/handler"
)

func Register(r *server.Hertz) {
	root := r.Group("/api/v1")
	_wechat := root.Group("/wechat")
	_wechat.GET("/callback/:agentid", handler.VerifyURL)
	_wechat.POST("/callback/:agentid", handler.Receive)
}
