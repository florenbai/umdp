package client

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"umdp/app/manage/biz/handler"
	"umdp/app/manage/biz/mw"
)

func Register(r *server.Hertz) {
	root := r.Group("/api/v1")
	_client := root.Group("/message")
	_client.Use(mw.KeyAuthHandler())
	{
		_client.POST("/send", handler.Send)
	}
}
