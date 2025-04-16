package model

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"umdp/app/manage/biz/dal/mysql"
	"umdp/pkg/response"
)

type Profession struct {
	Model
	ProfessionName string `gorm:"column:profession_name" json:"ProfessionName"`
	Token          string `gorm:"column:token"  json:"token"`
}

type ProfessionDetail struct {
	Model
	ProfessionName string `gorm:"column:profession_name" json:"ProfessionName"`
	Token          string `gorm:"column:token"  json:"token"`
	Channels       string `json:"channels"`
}

func NewProfessionModel() *Profession {
	return &Profession{}
}

func (m *Profession) TableName() string {
	return mysql.ProfessionTableName
}

// CreateProfession 创建业务
func (m *Profession) CreateProfession(ctx context.Context, channels []uint64) error {
	return Transaction(ctx, func(tx *gorm.DB) error {
		err := tx.WithContext(ctx).Table(m.TableName()).Create(&m).Error
		if err != nil {
			return err
		}
		var pcs []ProfessionChannel
		for _, v := range channels {
			pc := ProfessionChannel{
				ProfessionId: m.Id,
				ChannelId:    v,
			}
			pcs = append(pcs, pc)
		}
		return tx.WithContext(ctx).Table(mysql.ProfessionChannelTableName).Create(&pcs).Error
	})
}

// EditProfession 编辑业务
func (m *Profession) EditProfession(ctx context.Context, channels []uint64) error {
	return Transaction(ctx, func(tx *gorm.DB) error {
		err := tx.WithContext(ctx).Table(m.TableName()).Save(&m).Error
		if err != nil {
			return err
		}
		err = tx.WithContext(ctx).Table(mysql.ProfessionChannelTableName).Where("profession_id", m.Id).Delete(ProfessionChannel{}).Error
		if err != nil {
			return err
		}
		var pcs []ProfessionChannel
		for _, v := range channels {
			pc := ProfessionChannel{
				ProfessionId: m.Id,
				ChannelId:    v,
			}
			pcs = append(pcs, pc)
		}
		return tx.WithContext(ctx).Table(mysql.ProfessionChannelTableName).Create(&pcs).Error
	})
}

// ExistName 名称是否存在
func (m *Profession) ExistName(ctx context.Context, name string, id uint64) (bool, error) {
	var i int64
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Where("profession_name = ? AND id != ?", name, id).Count(&i).Error
	if err != nil {
		return true, err
	}
	if i > 0 {
		return true, nil
	}
	return false, nil
}

// GetProfessionList 获取业务列表
func (m *Profession) GetProfessionList(ctx context.Context, page uint64, pageSize uint64, scopes ...func(*gorm.DB) *gorm.DB) ([]ProfessionDetail, int64, error) {
	var i int64
	var list []ProfessionDetail
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).
		Joins("LEFT JOIN profession_channel ON profession_channel.profession_id = profession.id").
		Joins("LEFT JOIN channel ON channel.id = profession_channel.channel_id").
		Scopes(scopes...).Group("profession.id").Count(&i).Error
	if err != nil {
		return list, i, err
	}
	scopes = append(scopes, Paginate(page, pageSize))
	err = mysql.DB.WithContext(ctx).Table(m.TableName()).
		Select("profession.*,group_concat(channel.channel_name separator ',') as channels").
		Joins("LEFT JOIN profession_channel ON profession_channel.profession_id = profession.id").
		Joins("LEFT JOIN channel ON channel.id = profession_channel.channel_id").
		Scopes(scopes...).Group("profession.id").Find(&list).Error
	if err != nil {
		return list, i, err
	}
	return list, i, nil
}

// GetProfessionDetail 获取业务详情
func (m *Profession) GetProfessionDetail(ctx context.Context, id uint64) (ProfessionDetail, error) {
	var pd ProfessionDetail
	scopes := WhereWithScope("profession.id", id)
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).
		Select("profession.*,group_concat(channel.channel_name separator ',') as channels").
		Joins("LEFT JOIN profession_channel ON profession_channel.profession_id = profession.id").
		Joins("LEFT JOIN channel ON channel.id = profession_channel.channel_id").
		Scopes(scopes).Group("profession.id").Find(&pd).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return pd, response.DataNotFoundErr
		}
		return pd, err
	}
	return pd, nil
}

// Delete 删除业务及关联数据
func (m *Profession) Delete(ctx context.Context) error {
	return Transaction(ctx, func(tx *gorm.DB) error {
		err := tx.WithContext(ctx).Table(mysql.ProfessionChannelTableName).Where("profession_id", m.Id).Delete(&ProfessionChannel{}).Error
		if err != nil {
			return err
		}
		return tx.WithContext(ctx).Table(m.TableName()).Delete(&m, "id = ?", m.Id).Error
	})
}

// ExistToken 检测token是否存在
func (m *Profession) ExistToken(ctx context.Context, token string) (bool, error) {
	var i int64
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Where("token", token).Count(&i).Error
	if err != nil {
		return true, err
	}
	if i > 0 {
		return true, nil
	}
	return false, nil
}

// GetProfessionByToken 根据token获取业务信息
func (m *Profession) GetProfessionByToken(ctx context.Context, token string) (*Profession, error) {
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Where("token", token).First(&m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}
