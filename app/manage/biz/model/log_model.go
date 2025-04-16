package model

import (
	"context"
	"gorm.io/gorm"
	"umdp/app/manage/biz/dal/mysql"
)

type Log struct {
	Model
	ProfessionId uint64 `gorm:"column:profession_id" json:"professionId"`
	ChannelId    uint64 `gorm:"column:channel_id"  json:"channelId"`
	TemplateId   uint64 `gorm:"column:template_id"  json:"templateId"`
	Parameters   string `gorm:"column:parameters"  json:"parameters"`
	Receiver     string `gorm:"column:receiver" json:"receiver"`
	Status       int8   `gorm:"column:status"  json:"status"`
	RequestId    string `gorm:"column:request_id"  json:"requestId"`
	ErrMessage   string `gorm:"column:err_message"  json:"errMessage"`
}

type LogRes struct {
	Model
	ProfessionName string `gorm:"column:profession_name" json:"professionName"`
	ChannelName    string `gorm:"column:channel_name"  json:"channelName"`
	TemplateName   string `gorm:"column:template_name"  json:"templateName"`
	Parameters     string `gorm:"column:parameters"  json:"parameters"`
	Receiver       string `gorm:"column:receiver" json:"receiver"`
	Status         int8   `gorm:"column:status"  json:"status"`
	RequestId      string `gorm:"column:request_id"  json:"requestId"`
	ErrMessage     string `gorm:"column:err_message"  json:"errMessage"`
}

func NewLogModel() *Log {
	return &Log{}
}

func (m *Log) TableName() string {
	return mysql.LogTableName
}

// GetLogList 获取日志列表
func (m *Log) GetLogList(ctx context.Context, page uint64, pageSize uint64, scopes ...func(*gorm.DB) *gorm.DB) ([]LogRes, int64, error) {
	var i int64
	var list []LogRes
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Joins("LEFT JOIN profession ON profession.id = log.profession_id").
		Joins("LEFT JOIN channel ON channel.id = log.channel_id").
		Joins("LEFT JOIN template ON template.id = log.template_id").Scopes(scopes...).Count(&i).Error
	if err != nil {
		return list, i, err
	}
	scopes = append(scopes, Paginate(page, pageSize))
	err = mysql.DB.WithContext(ctx).Table(m.TableName()).
		Select("profession.profession_name,channel.channel_name,template.template_name,log.*").
		Joins("LEFT JOIN profession ON profession.id = log.profession_id").
		Joins("LEFT JOIN channel ON channel.id = log.channel_id").
		Joins("LEFT JOIN template ON template.id = log.template_id").
		Scopes(scopes...).
		Order("log.id DESC").
		Find(&list).Error
	if err != nil {
		return list, i, err
	}
	return list, i, nil
}
