package model

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"umdp/app/manage/biz/dal/mysql"
)

type TemplateChannel struct {
	TemplateId uint64 `gorm:"column:template_id" json:"templateId"`
	ChannelId  uint64 `gorm:"column:channel_id"  json:"channelId"`
	Config     string `gorm:"column:config"  json:"config"`
}

type TemplateConfigDetail struct {
	Retry         int64  `json:"retry"`
	ChannelId     uint64 `gorm:"column:channel_id"  json:"channelId"`
	ChannelName   string `gorm:"column:channel_name"  json:"channelName"`
	Config        string `gorm:"column:config"  json:"config"`
	ChannelConfig string `gorm:"column:channel_config"  json:"channelConfig"`
	ChannelStatus bool   `gorm:"column:channel_status" json:"channelStatus"`
}

func NewTemplateChannelModel() *TemplateChannel {
	return &TemplateChannel{}
}

func (m *TemplateChannel) TableName() string {
	return mysql.TemplateChannelTableName
}

// GetTemplateConfigByChannelTag 获取模板配置
func (m *TemplateChannel) GetTemplateConfigByChannelTag(ctx context.Context, tag string, professionId uint64, templateId uint64) (*TemplateConfigDetail, error) {
	var tcd TemplateConfigDetail
	err := mysql.DB.WithContext(ctx).Table(mysql.TemplateTableName).
		Select("template.retry,channel.channel_config,channel.channel_name,channel.channel_status,channel.id AS channel_id,template_channel.config").
		Joins("LEFT JOIN template_channel ON template_channel.template_id = template.id").
		Joins("LEFT JOIN channel ON channel.id = template_channel.channel_id").
		Where("channel.channel_tag", tag).
		Where("template.profession_id", professionId).
		Where("template.id", templateId).
		First(&tcd).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("未找到当前账号绑定的渠道信息")
		}
		return nil, err
	}
	return &tcd, nil
}

func (m *TemplateChannel) GetTemplateConfigByTemplateIdAndChannelId(ctx context.Context, templateId string, channelId string) error {
	return mysql.DB.WithContext(ctx).Table(m.TableName()).Where("template_id = ? AND channel_id = ?", templateId, channelId).First(&m).Error
}

func (m *TemplateChannel) GetTemplateChannelConfigsByTemplateId(ctx context.Context, templateId uint64) ([]string, error) {
	var list []string
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Select("config").Where("template_id", templateId).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, err
}
