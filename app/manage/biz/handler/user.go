package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	"umdp/app/manage/biz/service"
	"umdp/app/manage/model/request"
	"umdp/pkg/response"
)

// SSO
/*
func RedirectLogin(ctx context.Context, c *app.RequestContext) {
	next, ok := c.GetQuery("next")
	if !ok {
		response.SendBaseResp(c, response.RedirectNotFoundErr)
		return
	}
	var state = uuid.New().String()
	var nonce = uuid.New().String()
	v := url.Values{}
	v.Add("client_id", conf.GetConf().Oidc.ClientId)
	v.Add("redirect_uri", next)
	v.Add("response_type", "code")
	v.Add("state", state)
	v.Add("nonce", nonce)
	v.Add("scope", conf.GetConf().Oidc.Scope)
	v.Add("code_challenge_method", "S256")
	v.Add("code_challenge", uuid.New().String())
	providerObj, err := oidc.InitProvider(ctx, conf.GetConf().Oidc.OidcServer)
	if err != nil {
		response.SendBaseResp(c, fmt.Errorf("获取provider错误, %s", err))
		return
	}
	session := sessions.Default(c)
	session.Set("state", state)
	session.Set("nonce", nonce)
	session.Set("redirect_uri", next)
	err = session.Save()
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	authURL := providerObj.Endpoint().AuthURL
	redirectSSOUri := fmt.Sprintf("%s?%s", authURL, v.Encode())
	c.Redirect(http.StatusFound, []byte(redirectSSOUri))
}

func OidcCallback(ctx context.Context, c *app.RequestContext) {
	var err error
	code, ok := c.GetQuery("code")
	if !ok {
		response.SendBaseResp(c, response.OidcCodeErr)
		return
	}
	state, _ := c.GetQuery("state")
	redirectURLStr, ok := c.GetQuery("redirect_uri")
	if !ok {
		redirectURLStr = "/"
	}
	session := sessions.Default(c)
	cacheState, _ := session.Get("state").(string)
	if state != "" && state != cacheState {
		response.SendBaseResp(c, response.OidcCodeErr)
		return
	}
	uri := oidc.GetCallbackUri(c, "")
	oc := oidc.NewOIDC(ctx, conf.GetConf().Oidc.OidcServer, conf.GetConf().Oidc.ClientId,
		conf.GetConf().Oidc.ClientSecret, conf.GetConf().Oidc.Scope, uri)
	info, err := oc.OidcCallbackWithCode(ctx, uri, code)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	var claims json.RawMessage
	err = info.Claims(&claims)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	user := gjson.Parse(string(claims))
	nickname := user.Get("claims.nickname").String()
	session.Set("name", nickname)
	redirectURLStr = fmt.Sprintf("%s:%d%s", conf.AppDomain, conf.AppPort, redirectURLStr)
	c.Redirect(http.StatusFound, []byte(redirectURLStr))
	c.Abort()
}

*/

func LogOut(ctx context.Context, c *app.RequestContext) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	response.SendBaseResp(c, response.Success)
}

func UserInfo(ctx context.Context, c *app.RequestContext) {
	session := sessions.Default(c)
	data := map[string]interface{}{"name": session.Get("name")}
	response.SendDataResp(c, response.Success, data)
}

func Login(ctx context.Context, c *app.RequestContext) {
	var req request.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	token, user, err := service.NewUserService(ctx, c).VerifyUserInfo(req.Username, req.Password)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	data := map[string]interface{}{
		"token": token,
		"user": map[string]string{
			"username": user.Username,
			"nickname": user.Nickname,
		},
	}
	response.SendDataResp(c, response.Success, data)
}
