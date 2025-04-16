package model

import (
	"context"
	"umdp/app/manage/biz/dal/mysql"
)

type ProfessionChannel struct {
	ProfessionId uint64 `gorm:"column:profession_id"  json:"professionId"`
	ChannelId    uint64 `gorm:"column:channel_id"   json:"channelId"`
}

func NewProfessionChannelModel() *ProfessionChannel {
	return &ProfessionChannel{}
}

func (m *ProfessionChannel) TableName() string {
	return mysql.ProfessionChannelTableName
}

// GetChannelsByProfession 根据业务编号获取业务渠道编号数组
func (m *ProfessionChannel) GetChannelsByProfession(ctx context.Context, id uint64) ([]uint64, error) {
	var data []uint64
	err := GetPluck(ctx, m.TableName(), "channel_id", &data, WhereWithScope("profession_id", id))
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetChannelsDetailByProfession 根据业务获取渠道详情
func (m *ProfessionChannel) GetChannelsDetailByProfession(ctx context.Context, id uint64) ([]Channel, error) {
	var channels []Channel
	err := mysql.DB.WithContext(ctx).Select("channel.*").Table(m.TableName()).Joins("LEFT JOIN channel ON channel.id = profession_channel.channel_id").Where("profession_id", id).Find(&channels).Error
	if err != nil {
		return nil, err
	}
	return channels, nil
}
