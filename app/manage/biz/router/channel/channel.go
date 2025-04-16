package channel

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"umdp/app/manage/biz/handler"
)

func Register(r *server.Hertz) {
	root := r.Group("/api/v1")
	_channel := root.Group("/channel")
	_channel.POST("", handler.CreateChannel)
	_channel.GET("/list", handler.ChannelList)
	_channel.GET("/all", handler.AllChannel)
	_channel.GET("/:id", handler.ChannelDetail)
	_channel.PUT("/:id", handler.EditChannel)
	_channel.DELETE("/:id", handler.DeleteChannel)
}
