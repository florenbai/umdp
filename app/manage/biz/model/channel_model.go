package model

import (
	"context"
	"umdp/app/manage/biz/dal/mysql"
)

type Channel struct {
	Model
	ChannelName   string `gorm:"column:channel_name" json:"channelName"`
	ChannelTag    string `gorm:"column:channel_tag"  json:"channelTag"`
	ChannelConfig string `gorm:"column:channel_config"  json:"channelConfig"`
	ChannelStatus *int8  `gorm:"column:channel_status"  json:"channelStatus"`
}

func NewChannelModel() *Channel {
	return &Channel{}
}

func (m *Channel) TableName() string {
	return mysql.ChannelTableName
}

// ExistChannelByTag 检测渠道tag是否存在
func (m *Channel) ExistChannelByTag(ctx context.Context, tag string, id uint64) (bool, error) {
	var i int64
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Where("channel_tag = ? AND id != ?", tag, id).Count(&i).Error
	if err != nil {
		return true, err
	}
	if i > 0 {
		return true, nil
	}
	return false, nil
}

// GetChannelById 获取渠道信息
func (m *Channel) GetChannelById(ctx context.Context, id uint64) (Channel, error) {
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Where("id", id).First(&m).Error
	if err != nil {
		return *m, err
	}
	return *m, nil
}

func (m *Channel) GetChannelByAgentId(ctx context.Context, agentId string) error {
	return mysql.DB.WithContext(ctx).Table(m.TableName()).Where("JSON_EXTRACT(channel_config,'$.agentid') = ?", agentId).First(&m).Error
}
