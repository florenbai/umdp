package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"umdp/app/manage/biz/router/channel"
	"umdp/app/manage/biz/router/client"
	"umdp/app/manage/biz/router/log"
	"umdp/app/manage/biz/router/profession"
	"umdp/app/manage/biz/router/template"
	"umdp/app/manage/biz/router/user"
	"umdp/app/manage/biz/router/wechat"
)

func GeneratedRegister(r *server.Hertz) {
	user.Register(r)
	channel.Register(r)
	profession.Register(r)
	template.Register(r)
	client.Register(r)
	log.Register(r)
	wechat.Register(r)
}
