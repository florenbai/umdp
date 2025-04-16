package user

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"umdp/app/manage/biz/handler"
)

func Register(r *server.Hertz) {
	root := r.Group("/api/v1")
	_user := root.Group("/user")
	_user.POST("/login", handler.Login)
	_user.POST("/info", handler.UserInfo)
	_user.POST("/logout", handler.LogOut)
}
