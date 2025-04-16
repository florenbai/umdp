package service

import (
	"context"
	"encoding/xml"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-laoji/wxbizmsgcrypt"
	"github.com/google/uuid"
	"github.com/tidwall/gjson"
	"strings"
	"umdp/app/manage/biz/model"
	pb "umdp/hertz_gen/message"
	"umdp/pkg/channel"
	"umdp/pkg/http"
	"umdp/pkg/response"
	"umdp/pkg/wework"
)

type MessageService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewMessageService(ctx context.Context, c *app.RequestContext) *MessageService {
	return &MessageService{ctx: ctx, c: c}
}

func (service *MessageService) Send(req pb.SendRequest) (string, error) {
	var err error
	requestId := uuid.New().String()
	token, ok := service.c.Get("token")
	if !ok {
		return requestId, response.AuthorizeFailErr
	}
	variables, err := req.GetParameters().GetVariable().MarshalJSON()
	if err != nil {
		return requestId, err
	}

	var professionId, channelId uint64
	var receiver []string
	defer func() {
		logService := NewLogService(service.ctx, service.c)
		logService.Log(professionId, channelId, req.GetTemplateId(), string(variables), requestId, receiver, err)
	}()
	profession, err := model.NewProfessionModel().GetProfessionByToken(service.ctx, token.(string))
	if err != nil {
		return requestId, err
	}
	professionId = profession.Id
	tc, err := model.NewTemplateChannelModel().GetTemplateConfigByChannelTag(service.ctx, req.GetChannel(), profession.Id, req.GetTemplateId())
	if err != nil {
		return requestId, err
	}
	if !tc.ChannelStatus {
		return requestId, response.ChannelStatusCloseErr
	}
	channelId = tc.ChannelId
	channelHandler := channel.GetChannelHandler(tc.ChannelName)
	channelParam := channel.Parameters{
		Id:        req.GetParameters().GetId(),
		Receiver:  req.GetParameters().GetReceiver(),
		Cc:        req.GetParameters().GetCc(),
		Variables: string(variables),
	}
	receiver = req.GetParameters().GetReceiver()
	templateParam := channel.TemplateParameters{
		Id:     req.GetTemplateId(),
		Config: tc.Config,
	}
	handler, err := channelHandler.SetConfig(tc.ChannelConfig)
	if err != nil {
		return requestId, err
	}
	err = handler.Handle(service.ctx, channelParam, templateParam, tc.Retry)
	if err != nil {
		return requestId, err
	}
	return requestId, nil
}

func (service *MessageService) SendTest(req pb.SendRequest) error {
	var err error
	requestId := uuid.New().String()
	template := model.NewTemplateModel()
	err = template.GetTemplateDetail(service.ctx, req.GetTemplateId())
	if err != nil {
		return err
	}
	tc, err := model.NewTemplateChannelModel().GetTemplateConfigByChannelTag(service.ctx, req.GetChannel(), template.ProfessionId, req.GetTemplateId())
	if err != nil {
		return err
	}
	if !tc.ChannelStatus {
		return response.ChannelStatusCloseErr
	}
	variables, err := req.GetParameters().GetVariable().MarshalJSON()
	if err != nil {
		return err
	}
	var receiver []string
	defer func() {
		logService := NewLogService(service.ctx, service.c)
		logService.Log(template.ProfessionId, tc.ChannelId, req.GetTemplateId(), string(variables), requestId, receiver, err)
	}()

	channelHandler := channel.GetChannelHandler(tc.ChannelName)
	channelParam := channel.Parameters{
		Id:        req.GetParameters().GetId(),
		Receiver:  req.GetParameters().GetReceiver(),
		Cc:        req.GetParameters().GetCc(),
		Variables: string(variables),
	}
	receiver = req.GetParameters().GetReceiver()
	templateParam := channel.TemplateParameters{
		Id:     req.GetTemplateId(),
		Config: tc.Config,
	}
	handler, err := channelHandler.SetConfig(tc.ChannelConfig)
	if err != nil {
		return err
	}
	return handler.Handle(service.ctx, channelParam, templateParam, tc.Retry)
}

func (service *MessageService) HandleWechatReceive(agentId string, params wework.EventPushQueryBinding) error {
	wechatConfig, channelName, err := NewChannelService(service.ctx, service.c).GetChannelConfigByAgent(agentId)
	if err != nil {
		return err
	}
	channelHandler := channel.GetChannelHandler(channelName)
	handler, err := channelHandler.SetConfig(wechatConfig)
	if err != nil {
		return err
	}
	wechatObj := handler.(*channel.Wechat)
	wxcpt := wxbizmsgcrypt.NewWXBizMsgCrypt(wechatObj.GetSuiteToken(), wechatObj.GetSuiteEncodingAesKey(),
		wechatObj.GetSuiteId(), wxbizmsgcrypt.XmlType)
	msg, errMsg := wxcpt.DecryptMsg(params.MsgSign, params.Timestamp, params.Nonce, service.c.Request.Body())
	if errMsg != nil {
		return errors.New(errMsg.ErrMsg)
	}
	var msgType wework.Event
	err = xml.Unmarshal(msg, &msgType)
	if err != nil {
		return err
	}
	if msgType.MsgType == "event" {
		switch msgType.Event {
		case "template_card_event":
			var tce wework.ButtonTemplateEvent
			err = xml.Unmarshal(msg, &tce)
			if err != nil {
				return err
			}
			err = wechatObj.UnClickableTemplateCart(tce.ResponseCode)
			if err != nil {
				return err
			}
			url, err := service.getCallbackUrl(tce.TaskId)
			if err != nil {
				return err
			}
			taskArr := strings.Split(tce.TaskId, "-")
			_, err = http.Get(url, map[string]string{"id": taskArr[3], "event": tce.EventKey, "from": tce.FromUserName}, nil)
			if err != nil {
				return err
			}

		}
	}
	return nil
}

func (service *MessageService) getCallbackUrl(task string) (string, error) {
	taskArr := strings.Split(task, "-")
	if len(taskArr) < 4 {
		return "", response.ParamErr
	}
	tc := model.NewTemplateChannelModel()
	err := tc.GetTemplateConfigByTemplateIdAndChannelId(service.ctx, taskArr[1], taskArr[2])
	if err != nil {
		return "", err
	}
	callbackUrl := gjson.Get(tc.Config, "buttonTemplate.callback").String()
	return callbackUrl, nil
}
