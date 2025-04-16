package channel

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-laoji/wxbizmsgcrypt"
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
	"time"
	"umdp/conf"
	"umdp/pkg/wework"
)

type Wechat struct {
	wework.IWeWork
	AgentId uint `json:"agentId"`
}

const (
	TextMessage               = 1
	MarkDownMessage           = 2
	TemplateCardButtonMessage = 3
	TemplateCardTextMessage   = 4
)

type WechatOptions struct {
	CorpId         string `json:"corpId"`
	AgentId        string `json:"agentId"`
	CorpSecret     string `json:"corpSecret"`
	Secret         string `json:"secret"`
	Token          string `json:"token"`
	EncodingAESKey string `json:"encodingAESKey"`
}

type CardButtonInteraction struct {
	Source                string                         `json:"source"`
	MainTitle             wework.MainTitle               `json:"mainTitle"`
	HorizontalContentList []wework.HorizontalContentList `json:"horizontalContentList"`
	ButtonList            []wework.Button                `json:"buttonList"`
	Callback              string                         `json:"callback"`
	CardAction            wework.CardAction              `json:"cardAction"`
}

type CardTextNotice struct {
	Source                string                         `json:"source"`
	MainTitle             wework.MainTitle               `json:"mainTitle"`
	HorizontalContentList []wework.HorizontalContentList `json:"horizontalContentList"`
	EmphasisContent       wework.EmphasisContent         `json:"emphasisContent"`
	SubTitleText          string                         `json:"subTitleText"`
	CardAction            wework.CardAction              `json:"cardAction"`
}

func NewWechatChannel() MessageChannel {
	return Wechat{}
}

func (wechat Wechat) SetConfig(config string) (MessageChannel, error) {
	var options WechatOptions
	json.Unmarshal([]byte(config), &options)

	client := wework.NewWeWork(wework.WeWorkConfig{
		CorpId:              options.CorpId,
		SuiteToken:          options.Token,
		SuiteEncodingAesKey: options.EncodingAESKey,
		Dsn:                 conf.GetConf().MySQL.DSN,
	})
	client.SetCache(wework.RedisOpts{
		Host:     conf.GetConf().Redis.Address,
		Password: conf.GetConf().Redis.Password,
		Database: conf.GetConf().Redis.Db,
		MaxIdle:  conf.GetConf().Redis.MaxIdle,
	})
	agentId, _ := strconv.Atoi(options.AgentId)
	client.SetDebug(true)
	client.SetAppSecretFunc(func(corpId uint) (corpid string, secret string, customizedApp bool) {
		return options.CorpId, options.Secret, true
	})
	client.SetAgentIdFunc(func(corpId uint) (agentId int) {
		agentId, _ = strconv.Atoi(options.AgentId)
		return agentId
	})
	return &Wechat{
		client,
		uint(agentId),
	}, nil
}

// UnClickableTemplateCart 设置模板按钮禁止点击
func (wechat Wechat) UnClickableTemplateCart(responseCode string) error {
	message := wework.TemplateCardUpdateMessage{
		AtAll:        1,
		ResponseCode: responseCode,
		Button: struct {
			ReplaceName string `json:"replace_name" validate:"required"`
		}{
			ReplaceName: "审核完成",
		},
		ReplaceText: "",
	}
	resp := wechat.MessageUpdateTemplateCard(1, message)
	if resp.ErrCode != 0 {
		return errors.New(resp.ErrorMsg)
	}
	return nil
}

// VerifyURL 接受微信回调验证
func (wechat Wechat) VerifyURL(params wework.EventPushQueryBinding) (string, error) {
	wxcpt := wxbizmsgcrypt.NewWXBizMsgCrypt(wechat.GetSuiteToken(), wechat.GetSuiteEncodingAesKey(),
		wechat.GetCorpId(), wxbizmsgcrypt.XmlType)
	echoStr, err := wxcpt.VerifyURL(params.MsgSign, params.Timestamp, params.Nonce, params.EchoStr)
	if err != nil {
		return "", errors.New(err.ErrMsg)
	}
	return string(echoStr), nil
}

// Handle 处理企业微信消息发送
func (wechat Wechat) Handle(ctx context.Context, parameters Parameters, tp TemplateParameters, retry int64) error {
	configJson := gjson.Parse(tp.Config)
	message := wework.Message{
		ToUser:                 strings.Join(parameters.Receiver, "|"),
		EnableIDTrans:          0,
		EnableDuplicateCheck:   0,
		DuplicateCheckInterval: 0,
	}
	var resp wework.MessageSendResponse
	var messageInfo interface{}
	switch configJson.Get("messageType").Int() {
	case TextMessage:
		content := ParameterMatchFiled("content", parameters, tp.Config)
		messageInfo = wework.TextMessage{
			Message: message,
			Safe:    0,
			Text:    wework.Text{Content: content},
		}
	case MarkDownMessage:
		content := ParameterMatchFiled("content", parameters, tp.Config)
		messageInfo = wework.MarkDownMessage{
			Message:  message,
			MarkDown: wework.Text{Content: content},
		}
	case TemplateCardButtonMessage:
		cardParam, task := wechat.handleButtonTemplate(parameters, tp)
		messageInfo = wework.TemplateCardMessage{
			Message: message,
			TemplateCard: wework.TemplateCard{
				CardType: wework.CardTypeButtonInteraction,
				Source: wework.Source{
					IconURL:   "",
					Desc:      cardParam.Source,
					DescColor: 0,
				},
				TaskID:                task,
				MainTitle:             cardParam.MainTitle,
				HorizontalContentList: cardParam.HorizontalContentList,
				ButtonList:            cardParam.ButtonList,
				CardAction:            cardParam.CardAction,
			},
		}
	case TemplateCardTextMessage:
		cardParam := wechat.handleTextNotice(parameters, tp)
		messageInfo = wework.TemplateCardMessage{
			Message: message,
			TemplateCard: wework.TemplateCard{
				CardType: wework.CardTypeTextNotice,
				Source: wework.Source{
					IconURL:   "",
					Desc:      cardParam.Source,
					DescColor: 0,
				},
				MainTitle:             cardParam.MainTitle,
				EmphasisContent:       &cardParam.EmphasisContent,
				SubTitleText:          cardParam.SubTitleText,
				HorizontalContentList: cardParam.HorizontalContentList,
				CardAction:            cardParam.CardAction,
			},
		}
	}
	resp = wechat.MessageSend(wechat.AgentId, messageInfo)
	if resp.ErrCode != 0 {
		return errors.New(resp.ErrorMsg)
	}
	return nil
}

func (wechat Wechat) handleTextNotice(parameters Parameters, tp TemplateParameters) CardTextNotice {
	var cardParam CardTextNotice
	configJson := gjson.Parse(tp.Config)
	tntJson := configJson.Get("textNoticeTemplate").String()
	json.Unmarshal([]byte(tntJson), &cardParam)
	variables := gjson.Parse(parameters.Variables).String()
	if gjson.Valid(parameters.Variables) {
		for k, v := range cardParam.HorizontalContentList {
			cardParam.HorizontalContentList[k].Value = ParameterMatch(variables, v.Value)
		}
		cardParam.SubTitleText = ParameterMatch(variables, cardParam.SubTitleText)
		cardParam.CardAction.URL = ParameterMatch(variables, cardParam.CardAction.URL)
	}
	return cardParam
}

func (wechat Wechat) handleButtonTemplate(parameters Parameters, tp TemplateParameters) (CardButtonInteraction, string) {
	var cardParam CardButtonInteraction
	configJson := gjson.Parse(tp.Config)
	btJson := configJson.Get("buttonTemplate").String()
	json.Unmarshal([]byte(btJson), &cardParam)
	variables := gjson.Parse(parameters.Variables).String()
	if gjson.Valid(parameters.Variables) {
		for k, v := range cardParam.HorizontalContentList {
			cardParam.HorizontalContentList[k].Value = ParameterMatch(variables, v.Value)
			cardParam.CardAction.URL = ParameterMatch(variables, cardParam.CardAction.URL)
		}
	}
	return cardParam, fmt.Sprintf("%d-%d-%s-%s", time.Now().Unix(), tp.Id, configJson.Get("id").String(), parameters.Id)
}
