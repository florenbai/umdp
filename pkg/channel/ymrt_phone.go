package channel

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/tidwall/gjson"
	"log"
	"strconv"
	"time"
	"umdp/app/manage/biz/model"
	"umdp/pkg/utils"
	"umdp/pkg/zlog"
)

const YMRT_API_PHONE_URL = "http://voice.b2m.cn/voice/templateVariableSend"
const YMRT_API_REPORT_URL = "http://voice.b2m.cn/voice/getReport"

type Ymrt struct {
	Option    YmrtOptions
	CallCount int `json:"callCount"`
	Interval  int `json:"interval"`
	HasCall   int `json:"hasCall"`
}

type YmrtOptions struct {
	AppId      string `json:"appId"`
	AppSecret  string `json:"appSecret"`
	TemplateId string `json:"templateId"`
}

type YmrtPhoneParam struct {
	CallCount int    `json:"callCount"`
	Interval  int    `json:"interval"`
	Content   string `json:"content" validate:"required"`
}

type YmrtPhoneResponse struct {
	VoiceId       string `json:"voiceId"`
	Mobile        string `json:"mobile"`
	CustomVoiceId string `json:"customVoiceId"`
}

type YmrtPhoneRequest struct {
	AppId             string               `json:"appId"`
	Timestamp         string               `json:"timestamp"`
	Sign              string               `json:"sign"`
	TemplateId        string               `json:"templateId"`
	Variables         []YmrtMobileVariable `json:"variables"`
	TriggerConditions string               `json:"triggerConditions"`
	ConnectTime       string               `json:"connectTime"`
}

type YmrtMobileVariable struct {
	Mobile        string            `json:"mobile"`
	Variable      YmrtPhoneVariable `json:"variable"`
	CustomVoiceId string            `json:"customVoiceId"`
}

type YmrtPhoneVariable struct {
	AlertTitle string `json:"alertTitle"`
	AlertCount string `json:"alertCount"`
	Duration   string `json:"duration"`
}

func NewYmrtChannel() MessageChannel {
	return Ymrt{}
}

func (ymrt Ymrt) SetConfig(config string) (MessageChannel, error) {
	var options YmrtOptions
	json.Unmarshal([]byte(config), &options)

	return &Ymrt{
		options,
		0, 0, 0,
	}, nil
}

func (ymrt Ymrt) Handle(ctx context.Context, parameters Parameters, tp TemplateParameters, retry int64) error {
	if len(parameters.Receiver) < 1 {
		return errors.New("必须含有一个接收者")
	}
	var param YmrtPhoneParam
	err := json.Unmarshal([]byte(tp.Config), &param)
	if err != nil {
		return errors.New("电话通道参数不正确")
	}
	ymrt.CallCount = param.CallCount
	ymrt.Interval = param.Interval
	_, err = ymrt.sendPhone(ctx, parameters.Receiver, parameters.Variables, parameters.Id, "")
	if err != nil {
		return err
	}
	return nil
}

// 发送电话语音
func (ymrt Ymrt) sendPhone(ctx context.Context, receiver []string, variables string, channelId string, customVoiceId string) ([]YmrtPhoneResponse, error) {
	logModel := new(model.PhoneLog)
	//
	if customVoiceId != "" {
		err := model.GetOneWithScope(ctx, logModel.TableName(), &logModel, model.WhereWithScope("custom_voice_id", customVoiceId))
		if err != nil {
			return nil, err
		}
		if logModel.CallLimit < logModel.CallCount {
			return nil, err
		}
	}
	var sendResp []YmrtPhoneResponse
	now := time.Now().Format("20060102150405")
	sign := utils.MD5HashString(fmt.Sprintf("%s%s%s", ymrt.Option.AppId, ymrt.Option.AppSecret, now))
	vj := gjson.Parse(variables)
	alertTitle := vj.Get("alertTitle").String()
	alertCount := vj.Get("alertCount").String()
	duration := vj.Get("duration").String()

	ypv := YmrtPhoneVariable{
		AlertTitle: alertTitle,
		AlertCount: alertCount,
		Duration:   duration,
	}
	var ymrtRequestVariables []YmrtMobileVariable
	for _, mobile := range receiver {
		item := YmrtMobileVariable{
			Mobile:   mobile,
			Variable: ypv,
		}
		if customVoiceId == "" {
			item.CustomVoiceId = uuid.New().String()
		} else {
			item.CustomVoiceId = customVoiceId
		}

		ymrtRequestVariables = append(ymrtRequestVariables, item)
	}
	params := YmrtPhoneRequest{
		AppId:             ymrt.Option.AppId,
		Timestamp:         now,
		Sign:              sign,
		TemplateId:        ymrt.Option.TemplateId,
		Variables:         ymrtRequestVariables,
		TriggerConditions: "4",
		ConnectTime:       "60",
	}

	header := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
	b, _ := json.Marshal(params)
	client := resty.New()
	zlog.Debug(fmt.Sprintf("发送电话请求：%v", string(b)))
	resp, err := client.R().SetHeaders(header).SetBody(params).Post(YMRT_API_PHONE_URL)
	if err != nil {
		return nil, err
	}
	zlog.Debug(string(resp.Body()))
	body := gjson.ParseBytes(resp.Body())
	if !body.Get("code").Exists() {
		log.Println(resp.String())
		return nil, errors.New("未找到返回信息")
	}
	if body.Get("code").String() != "SUCCESS" {
		return nil, errors.New(body.Get("code").String())
	}

	for _, v := range ymrtRequestVariables {
		if customVoiceId == "" {
			cid, _ := strconv.Atoi(channelId)
			err := model.Create(ctx, logModel.TableName(), &model.PhoneLog{
				ChannelId:     uint64(cid),
				CustomVoiceId: v.CustomVoiceId,
				Phone:         v.Mobile,
				CallCount:     1,
				CallLimit:     ymrt.CallCount,
			})
			if err != nil {
				log.Println(err)
				return nil, err
			}
		} else {
			err = logModel.IncrByCustomId(ctx, customVoiceId)
			if err != nil {
				log.Println(err)
				return nil, err
			}
		}
	}
	err = json.Unmarshal([]byte(body.Get("data").String()), &sendResp)
	if err != nil {
		return nil, err
	}
	go ymrt.getCallReport(ctx, sendResp, variables)
	return sendResp, nil
}

func (ymrt Ymrt) getCallReport(ctx context.Context, sendResp []YmrtPhoneResponse, variables string) {
	logModel := new(model.PhoneLog)
	// 等待一分钟查询
	var m = make(map[string]YmrtPhoneResponse)
	time.Sleep(time.Minute * 3)

	for _, v := range sendResp {
		m[v.CustomVoiceId] = v
	}
	callReport, err := ymrt.getPhoneStatus()
	if err != nil {
		return
	}
	// 获取报告
	for _, r := range callReport {
		if _, ok := m[r.CustomVoiceId]; ok {
			err := model.GetOneWithScope(ctx, logModel.TableName(), &logModel, model.WhereWithScope("custom_voice_id", r.CustomVoiceId))
			if err != nil {
				log.Println(err)
				return
			}
			// 当成功时，记录日志
			if r.State != "SUCCESS" {
				_, err = ymrt.sendPhone(ctx, []string{r.Mobile}, variables, fmt.Sprintf("%d", logModel.ChannelId), r.CustomVoiceId)
				if err != nil {
					log.Println(err)
					return
				}
			} else {
				logModel.CallResponse = r
				err = model.EditByScopes(ctx, logModel.TableName(), &logModel, model.WhereWithScope("custom_voice_id", r.CustomVoiceId))
				if err != nil {
					log.Println(err)
					return
				}
			}
		}
	}
}

func (ymrt Ymrt) getPhoneStatus() ([]model.CallReport, error) {
	var cr []model.CallReport
	now := time.Now().Format("20060102150405")
	sign := utils.MD5HashString(fmt.Sprintf("%s%s%s", ymrt.Option.AppId, ymrt.Option.AppSecret, now))
	sendurl := fmt.Sprintf("%s?appId=%s&timestamp=%s&sign=%s&number=200", YMRT_API_REPORT_URL, ymrt.Option.AppId, now, sign)
	client := resty.New()
	resp, err := client.R().Post(sendurl)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	body := gjson.ParseBytes(resp.Body())
	if !body.Get("code").Exists() {
		log.Println(resp.String())
		return nil, errors.New("未找到返回信息")
	}
	err = json.Unmarshal([]byte(body.Get("data").String()), &cr)
	if err != nil {
		return nil, err
	}
	return cr, nil
}
