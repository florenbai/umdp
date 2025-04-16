package model

import (
	"context"
	"gorm.io/gorm"
	"umdp/app/manage/biz/dal/mysql"
)

type CallReport struct {
	AppId           string   `json:"appId"`
	CustomVoiceId   string   `json:"customVoiceId"`
	TemplateId      string   `json:"templateId"`
	VoiceId         string   `json:"voiceId"`
	Mobile          string   `json:"mobile"`
	CallingTime     string   `json:"callingTime"`
	CallingDuration *float64 `json:"callingDuration"`
	Message         string   `json:"message"`
	State           string   `json:"state"`
}

type PhoneLog struct {
	Model
	CallResponse  CallReport `gorm:"column:call_resp;serializer:json"  json:"callResponse"`
	ChannelId     uint64     `gorm:"column:channel_id"  json:"channelId"`
	CustomVoiceId string     `json:"customVoiceId"`
	Phone         string     `json:"phone"`
	CallCount     int        `json:"callCount"`
	CallLimit     int        `json:"callLimit"`
}

func NewPhoneLogModel() *PhoneLog {
	return &PhoneLog{}
}

func (m *PhoneLog) TableName() string {
	return mysql.PhoneLogTableName
}

func (m *PhoneLog) IncrByCustomId(ctx context.Context, customVoiceId string) error {
	return mysql.DB.WithContext(ctx).Table(m.TableName()).Where("custom_voice_id", customVoiceId).UpdateColumn("call_count", gorm.Expr("call_count + ?", 1)).Error
}
