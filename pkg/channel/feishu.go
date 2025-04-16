package channel

import (
	"context"
	"encoding/json"
	"errors"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"github.com/tidwall/gjson"
	"time"
)

type Feishu struct {
	*lark.Client
}

type FeishuOptions struct {
	AppId     string `json:"appId"`
	AppSecret string `json:"appSecret"`
}

type FeishuTextMessageRequest struct {
	Text string `json:"text"`
}

const (
	FeishuTextMessage = 1
)

func NewFeishuChannel() MessageChannel {
	return Feishu{}
}

func (feishu Feishu) SetConfig(config string) (MessageChannel, error) {
	var options FeishuOptions
	json.Unmarshal([]byte(config), &options)
	client := lark.NewClient(options.AppId, options.AppSecret, lark.WithReqTimeout(3*time.Second), lark.WithEnableTokenCache(true))
	return &Feishu{
		client,
	}, nil
}

func (feishu Feishu) Handle(ctx context.Context, parameters Parameters, tp TemplateParameters, retry int64) error {
	if len(parameters.Receiver) < 1 {
		return errors.New("必须含有一个接收者")
	}
	configJson := gjson.Parse(tp.Config)
	req := larkim.NewCreateMessageReqBuilder().ReceiveIdType(`open_id`)
	body := larkim.NewCreateMessageReqBodyBuilder()
	switch configJson.Get("messageType").Int() {
	case FeishuTextMessage:
		content := ParameterMatchFiled("content", parameters, tp.Config)
		text := FeishuTextMessageRequest{Text: content}
		textByte, err := json.Marshal(text)
		if err != nil {
			return err
		}
		body.MsgType("text")
		body.Content(string(textByte))
	}
	for _, v := range parameters.Receiver {
		body.ReceiveId(v)

		resp, err := feishu.Im.Message.Create(context.Background(), req.Body(body.Build()).Build())
		if err != nil {
			return err
		}
		// 服务端错误处理
		if !resp.Success() {
			return errors.New(resp.Msg)
		}
	}
	return nil
}
