package main

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	"github.com/hertz-contrib/sessions"
	"time"
	"umdp/app/manage/biz/dal"
	"umdp/app/manage/biz/dal/redisStore"
	"umdp/app/manage/biz/mw"
	"umdp/app/manage/biz/router"
	"umdp/conf"
)

func init() {
	//初始化数据库
	dal.Init()

}

func main() {
	h := server.Default(server.WithHostPorts(conf.GetConf().Server.Addr))
	h.Use(gzip.Gzip(gzip.DefaultCompression))
	h.Use(accesslog.New(accesslog.WithFormat("[${time}] ${status} - ${latency} ${method} ${path} ${queryParams}")))
	h.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	if conf.GetConf().Authentication.EnableSession {
		h.Use(sessions.New(fmt.Sprintf("%s_session", conf.Dom), redisStore.Store))
	}
	h.Use(mw.JWTAuth([]string{"/api/v1/user/login"}))
	router.GeneratedRegister(h)
	h.Spin()
}
