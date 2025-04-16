package channel

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"github.com/tidwall/gjson"
)

type TencentSms struct {
	*sms.Client
}

type TencentSmsOptions struct {
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	RegionId        string `json:"regionId"`
	Endpoint        string `json:"endpoint"`
}

type TencentSmsParam struct {
	AppId           string   `json:"appId" validate:"required"`
	SignName        string   `json:"signName" validate:"required"`
	TemplateCode    string   `json:"templateCode" validate:"required"`
	TemplateParam   []string `json:"templateParam"`
	SmsUpExtendCode string   `json:"smsUpExtendCode"`
}

type TencentOptions func(tencent *TencentSmsOptions)

func NewTencentSmsChannel() MessageChannel {
	return TencentSms{}
}

func (tencentSms TencentSms) SetConfig(config string) (MessageChannel, error) {
	var options TencentSmsOptions
	json.Unmarshal([]byte(config), &options)

	credential := common.NewCredential(
		options.AccessKeyId,
		options.AccessKeySecret,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.Endpoint = options.Endpoint
	cpf.SignMethod = "HmacSHA1"
	client, err := sms.NewClient(credential, options.RegionId, cpf)
	return &TencentSms{
		client,
	}, err
}

func (tencentSms TencentSms) Handle(ctx context.Context, parameters Parameters, tp TemplateParameters, retry int64) error {
	if len(parameters.Receiver) < 1 {
		return errors.New("必须含有一个手机号码")
	}
	var config TencentSmsParam
	json.Unmarshal([]byte(tp.Config), &config)
	err := validator.New().Struct(config)
	if err != nil {
		return err
	}
	request := sms.NewSendSmsRequest()
	var phone []string
	for _, v := range parameters.Receiver {
		phone = append(phone, fmt.Sprintf("+86%s", v))
	}
	var templateParamSet []string
	if gjson.Valid(parameters.Variables) {
		variables := gjson.Parse(parameters.Variables).String()
		for _, v := range config.TemplateParam {
			pm := ParameterMatch(variables, v)
			templateParamSet = append(templateParamSet, pm)
		}
	} else {
		templateParamSet = config.TemplateParam
	}
	request.PhoneNumberSet = common.StringPtrs(phone)
	request.SmsSdkAppId = common.StringPtr(config.AppId)
	request.SignName = common.StringPtr(config.SignName)
	request.TemplateId = common.StringPtr(config.TemplateCode)
	request.TemplateParamSet = common.StringPtrs(templateParamSet)
	request.ExtendCode = common.StringPtr(config.SmsUpExtendCode)

	response, err := tencentSms.SendSms(request)
	if err != nil {
		return err
	}
	b, _ := json.Marshal(response.Response)
	resp := gjson.ParseBytes(b)
	var errMsg string
	resp.Get("SendStatusSet").ForEach(func(key, value gjson.Result) bool {
		if value.Get("Code").String() != "Ok" {
			errMsg = value.Get("Message").String()
			return false
		}
		return true
	})
	if errMsg != "" {
		return errors.New(errMsg)
	}

	return nil
}
