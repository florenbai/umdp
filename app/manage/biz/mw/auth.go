package mw

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"strings"
	"umdp/conf"
	"umdp/pkg/jwt"
	"umdp/pkg/response"
)

var (
	HeaderAuthorization = "Authorization"
	HeaderTag           = "userinfo"
)

func shouldSkip(path string, skip []string) bool {
	for _, p := range skip {
		if p == path {
			return true
		}
		// 处理通配符情况，如 /api/public/*
		if strings.HasSuffix(p, "/*") {
			prefix := strings.TrimSuffix(p, "/*")
			if strings.HasPrefix(path, prefix) {
				return true
			}
		}
	}
	return false
}

func JWTAuth(skip []string) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		path := string(c.Request.URI().Path())
		// 检查是否在排除路径中
		if shouldSkip(path, skip) {
			c.Next(ctx)
			return
		}
		token := string(c.GetHeader("Authorization"))
		if token == "" {
			response.SendBaseResp(c, response.UnauthorizedErr)
			c.Abort()
			return
		}
		// 移除Bearer前缀
		token = strings.Replace(token, "Bearer ", "", 1)
		claims, err := jwt.ParseToken(token, []byte(conf.GetConf().Authentication.AuthSecret))
		if err != nil {
			response.SendBaseResp(c, response.UnauthorizedErr)
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Next(ctx)
	}
}
