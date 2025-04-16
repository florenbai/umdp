package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"umdp/app/manage/biz/service"
	"umdp/pkg/channel"
	"umdp/pkg/response"
	"umdp/pkg/wework"
)

// VerifyURL 微信回调
func VerifyURL(ctx context.Context, c *app.RequestContext) {
	var params wework.EventPushQueryBinding
	agentid := c.Param("agentid")
	err := c.Bind(&params)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	wechatConfig, channelName, err := service.NewChannelService(ctx, c).GetChannelConfigByAgent(agentid)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	channelHandler := channel.GetChannelHandler(channelName)
	handler, err := channelHandler.SetConfig(wechatConfig)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	wechatObj := handler.(*channel.Wechat)
	echoStr, err := wechatObj.VerifyURL(params)
	if err != nil {
		wechatObj.Logger().Sugar().Error(err.Error())
		response.SendBaseResp(c, err)
		return
	}
	c.WriteString(echoStr)
}

// Receive 微信POST回调
func Receive(ctx context.Context, c *app.RequestContext) {
	var params wework.EventPushQueryBinding
	agentid := c.Param("agentid")
	err := c.Bind(&params)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewMessageService(ctx, c).HandleWechatReceive(agentid, params)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
}
