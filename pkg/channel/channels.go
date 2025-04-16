package channel

import (
	"context"
	"fmt"
	"github.com/tidwall/gjson"
	"regexp"
	"strings"
)

type MessageChannel interface {
	SetConfig(config string) (MessageChannel, error)
	Handle(ctx context.Context, parameters Parameters, tp TemplateParameters, retry int64) error
}

type Parameters struct {
	Receiver  []string `json:"receiver"`  //接收者
	Cc        []string `json:"cc"`        //抄送者
	Variables string   `json:"variables"` //变量
	Id        string   `json:"id"`        //渠道编号
}

type TemplateParameters struct {
	Id     uint64
	Config string
}

func GetChannelHandler(channel string) MessageChannel {
	return handlers[channel]
}

func ParameterMatch(variable string, field string) string {
	if !gjson.Valid(variable) {
		field = StringMatch(variable, field)
	} else {
		field = ObjectMatch(variable, field)
	}
	return field
}

func StringMatch(value string, field string) string {
	re := regexp.MustCompile(`\{\$[a-z0-9]+\}`)
	return re.ReplaceAllString(field, value)
}

func ArrayMatch(variable string, field string) string {
	re := regexp.MustCompile(`\{\$[a-z0-9]+\}`)
	findArr := re.FindAllString(field, -1)
	gjson.Parse(variable).ForEach(func(key, value gjson.Result) bool {
		field = strings.ReplaceAll(field, findArr[key.Int()], value.String())
		return true
	})
	return field
}

func ObjectMatch(variable string, field string) string {
	gjson.Parse(variable).ForEach(func(key, value gjson.Result) bool {
		field = strings.ReplaceAll(field, fmt.Sprintf("{$%s}", key.String()), value.String())
		return true
	})
	return field
}

func ParameterMatchFiled(field string, parameters Parameters, config string) string {
	configJson := gjson.Parse(config)
	content := configJson.Get(field).String()

	if gjson.Valid(parameters.Variables) {
		content = ParameterMatch(parameters.Variables, content)
	}
	return content
}
