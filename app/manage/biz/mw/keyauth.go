package mw

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/savsgio/gotils/strconv"
	"net/http"
	"umdp/app/manage/biz/model"
	"umdp/pkg/response"
)

func KeyAuthHandler() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token := strconv.B2S(c.GetHeader("token"))
		if len(token) < 1 {
			c.AbortWithMsg(response.AuthorizeFailErr.GetErrMsg(), http.StatusUnauthorized)
		}
		ok, err := model.NewProfessionModel().ExistToken(ctx, token)
		if err != nil {
			c.AbortWithMsg(response.Err.GetErrMsg(), http.StatusInternalServerError)
		}
		if !ok {
			c.AbortWithMsg(response.AuthorizeFailErr.GetErrMsg(), http.StatusUnauthorized)
		}
		c.Set("token", token)
	}
}
