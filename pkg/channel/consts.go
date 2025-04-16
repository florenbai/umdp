package channel

const (
	WechatChannel     = "企业微信"
	EmailChannel      = "邮箱"
	AliyunSmsChannel  = "阿里云短信"
	TencentSmsChannel = "腾讯云短信"
	DingtalkChannel   = "钉钉"
	FeishuChannel     = "飞书"
	YmrtChannel       = "亿美软通电话"
	WechatBotChannel  = "企业微信-群机器人"
)

var handlers = map[string]MessageChannel{
	WechatChannel:     NewWechatChannel(),
	EmailChannel:      NewEmailChannel(),
	AliyunSmsChannel:  NewAliyunSmsChannel(),
	TencentSmsChannel: NewTencentSmsChannel(),
	DingtalkChannel:   NewDingTalkChannel(),
	FeishuChannel:     NewFeishuChannel(),
	YmrtChannel:       NewYmrtChannel(),
	WechatBotChannel:  NewWechatBotChannel(),
}
