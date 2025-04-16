package channel

import (
	"context"
	"errors"
	"github.com/tidwall/gjson"
	"umdp/pkg/wework"
)

type WechatBot struct {
	wework.IWeWork
}

func NewWechatBotChannel() MessageChannel {
	return WechatBot{}
}

func (wechat WechatBot) SetConfig(config string) (MessageChannel, error) {
	client := wework.NewWeWork(wework.WeWorkConfig{})

	return &WechatBot{
		client,
	}, nil
}

func (wechat WechatBot) Handle(ctx context.Context, parameters Parameters, tp TemplateParameters, retry int64) error {
	configJson := gjson.Parse(tp.Config)
	var resp wework.MessageSendResponse
	var messageInfo interface{}
	switch configJson.Get("messageType").Int() {
	case TextMessage:
		content := ParameterMatchFiled("content", parameters, tp.Config)
		messageInfo = wework.BotTextMessage{
			Text: wework.Text{
				Content: content,
			},
		}
	case MarkDownMessage:
		content := ParameterMatchFiled("content", parameters, tp.Config)
		messageInfo = wework.BotMarkDownMessage{
			MarkDown: wework.Text{Content: content},
		}
	}
	for _, v := range parameters.Receiver {
		resp = wechat.BotMessageSend(v, messageInfo)
		if resp.ErrCode != 0 {
			return errors.New(resp.ErrorMsg)
		}
	}

	return nil
}
