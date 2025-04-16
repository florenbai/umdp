package channel

import (
	"context"
	"encoding/json"
	"errors"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dingtalkoauth "github.com/alibabacloud-go/dingtalk/oauth2_1_0"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/redis/go-redis/v9"
	"github.com/tidwall/gjson"
	"net/http"
	"strings"
	"time"
	"umdp/conf"
)

type DingTalk struct {
	*openapi.Client
	*DingTalkOptions
}

const (
	DingTalkTextMessage     = 1
	DingTalkMarkDownMessage = 2
)
const AccessTokenCacheKey = "dingtalk_accessToken"

type DingTalkOptions struct {
	AppKey    string `json:"appKey"`
	AppSecret string `json:"appSecret"`
	AgentID   string `json:"agentID"`
}

type TextMessageRequest struct {
	MsgType string       `json:"msgtype"`
	Text    DingTalkText `json:"text"`
}

type MarkDownMessageRequest struct {
	MsgType  string           `json:"msgtype"`
	Markdown DingTalkMarkDown `json:"markdown"`
}

type DingTalkText struct {
	Content string `json:"content"`
}

type DingTalkMarkDown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func NewDingTalkChannel() MessageChannel {
	return DingTalk{}
}

func (dingTalk DingTalk) SetConfig(config string) (MessageChannel, error) {
	var options DingTalkOptions
	json.Unmarshal([]byte(config), &options)
	dtConfig := &openapi.Config{
		Endpoint:           tea.String("oapi.dingtalk.com"),
		Protocol:           tea.String("HTTPS"),
		Method:             tea.String("POST"),
		SignatureAlgorithm: tea.String("v2"),
	}
	client, err := openapi.NewClient(dtConfig)
	return &DingTalk{
		client,
		&options,
	}, err
}

func (dingTalk DingTalk) Handle(ctx context.Context, parameters Parameters, tp TemplateParameters, retry int64) error {
	if len(parameters.Receiver) < 1 {
		return errors.New("必须含有一个接收者")
	}
	params := &openapi.Params{
		Style:       tea.String("ROA"),
		Method:      tea.String("POST"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/topapi/message/corpconversation/asyncsend_v2"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	queries := map[string]interface{}{}
	bodys := map[string]interface{}{}
	token, err := dingTalk.getAccessToken(*dingTalk.DingTalkOptions)
	if err != nil {
		return err
	}
	queries["access_token"] = token
	bodys["agent_id"] = dingTalk.AgentID
	bodys["userid_list"] = strings.Join(parameters.Receiver, ",")

	configJson := gjson.Parse(tp.Config)
	var messageInfo interface{}
	switch configJson.Get("messageType").Int() {
	case DingTalkTextMessage:
		content := ParameterMatchFiled("content", parameters, tp.Config)
		messageInfo = TextMessageRequest{
			MsgType: "text",
			Text: DingTalkText{
				Content: content,
			},
		}
	case DingTalkMarkDownMessage:
		title := ParameterMatchFiled("title", parameters, tp.Config)
		content := ParameterMatchFiled("content", parameters, tp.Config)
		messageInfo = MarkDownMessageRequest{
			MsgType: "markdown",
			Markdown: DingTalkMarkDown{
				Title: title,
				Text:  content,
			},
		}
	}

	bodys["msg"] = messageInfo
	runtime := &util.RuntimeOptions{}
	request := &openapi.OpenApiRequest{
		Query: openapiutil.Query(queries),
		Body:  bodys,
	}
	result, err := dingTalk.CallApi(params, request, runtime)
	if err != nil {
		return err
	}
	b, _ := json.Marshal(result)
	resp := gjson.ParseBytes(b).Get("body")
	if resp.Exists() {
		if resp.Get("errcode").String() != "0" {
			return errors.New(resp.Get("errmsg").String())
		}
	} else {
		return errors.New("未知返回值")
	}
	return nil
}

func (dingTalk DingTalk) getAccessToken(options DingTalkOptions) (string, error) {
	cache := dingTalk.GetCache()
	var accessToken string
	accessToken, err := cache.Get(context.TODO(), AccessTokenCacheKey).Result()
	if err != nil {
		dtConfig := &openapi.Config{
			Protocol: tea.String("https"),
			RegionId: tea.String("central"),
		}
		client, err := dingtalkoauth.NewClient(dtConfig)
		if err != nil {
			return "", err
		}
		req := dingtalkoauth.GetAccessTokenRequest{
			AppKey:    &options.AppKey,
			AppSecret: &options.AppSecret,
		}
		resp, err := client.GetAccessToken(&req)
		if err != nil {
			return "", err
		}
		if *resp.StatusCode != http.StatusOK {
			return "", errors.New(resp.String())
		}
		accessToken = *resp.Body.AccessToken
		err = cache.Set(context.TODO(), AccessTokenCacheKey, accessToken, time.Second*time.Duration(*resp.Body.ExpireIn)).Err()
		if err != nil {
			return "", err
		}
	}
	return accessToken, nil
}

func (dingTalk DingTalk) GetCache() redis.UniversalClient {
	conn := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:        []string{conf.GetConf().Redis.Address},
		DB:           conf.GetConf().Redis.Db,
		Password:     conf.GetConf().Redis.Password,
		DialTimeout:  time.Second * 20,
		MinIdleConns: conf.GetConf().Redis.MaxIdle,
	})
	return conn
}
