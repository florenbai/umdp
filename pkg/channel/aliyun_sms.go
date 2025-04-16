package channel

import (
	"context"
	"encoding/json"
	"errors"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/go-playground/validator/v10"
	"github.com/tidwall/gjson"
	"strings"
)

type AliyunSms struct {
	*openapi.Client
}

type AliyunSmsOptions struct {
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	RegionId        string `json:"regionId"`
	Endpoint        string `json:"endpoint"`
}

type AliyunSmsTemplateParam struct {
	Name  string `json:"name" validate:"required"`
	Value string `json:"value" validate:"required"`
}

type AliyunSmsParam struct {
	SignName        string                   `json:"signName" validate:"required"`
	TemplateCode    string                   `json:"templateCode" validate:"required"`
	TemplateParam   []AliyunSmsTemplateParam `json:"templateParam"`
	SmsUpExtendCode string                   `json:"smsUpExtendCode"`
	OutId           string                   `json:"outId"`
}

type AliOptions func(aliyun *AliyunSmsOptions)

func NewAliyunSmsChannel() MessageChannel {
	return AliyunSms{}
}

func (aliyunSms AliyunSms) SetConfig(config string) (MessageChannel, error) {
	var options AliyunSmsOptions
	json.Unmarshal([]byte(config), &options)
	smsConfig := &openapi.Config{
		AccessKeyId:     &options.AccessKeyId,
		AccessKeySecret: &options.AccessKeySecret,
		RegionId:        &options.RegionId,
		Endpoint:        &options.Endpoint,
	}
	client, err := openapi.NewClient(smsConfig)
	return &AliyunSms{
		client,
	}, err
}

func (aliyunSms AliyunSms) Handle(ctx context.Context, parameters Parameters, tp TemplateParameters, retry int64) error {
	if len(parameters.Receiver) < 1 {
		return errors.New("必须含有一个手机号码")
	}
	var config AliyunSmsParam
	json.Unmarshal([]byte(tp.Config), &config)
	err := validator.New().Struct(config)
	if err != nil {
		return err
	}
	params := &openapi.Params{
		Action:      tea.String("SendSms"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		Pathname:    tea.String("/"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}

	phoneNumbers := strings.Join(parameters.Receiver, ",")

	queries := map[string]interface{}{}
	queries["PhoneNumbers"] = &phoneNumbers
	queries["SignName"] = &config.SignName
	queries["TemplateCode"] = &config.TemplateCode
	if config.TemplateParam != nil {
		templateParam := make(map[string]string)
		for _, v := range config.TemplateParam {
			templateParam[v.Name] = v.Value
		}
		if gjson.Valid(parameters.Variables) {
			variables := gjson.Parse(parameters.Variables).String()
			for k, v := range templateParam {
				templateParam[k] = ParameterMatch(variables, v)
			}
		}

		tpByte, err := json.Marshal(templateParam)
		if err != nil {
			return err
		}
		queries["TemplateParam"] = tea.String(string(tpByte))
	}
	if config.SmsUpExtendCode != "" {
		queries["SmsUpExtendCode"] = tea.String(config.SmsUpExtendCode)
	}
	if config.OutId != "" {
		queries["OutId"] = tea.String(config.OutId)
	}
	runtime := &util.RuntimeOptions{}
	request := &openapi.OpenApiRequest{
		Query: openapiutil.Query(queries),
	}
	result, err := aliyunSms.CallApi(params, request, runtime)
	if err != nil {
		return err
	}
	b, _ := json.Marshal(result)
	resp := gjson.ParseBytes(b).Get("body")
	if resp.Exists() {
		if resp.Get("Message").String() != "OK" {
			return errors.New(resp.Get("Message").String())
		}
	} else {
		return errors.New("未知返回值")
	}
	return nil
}
