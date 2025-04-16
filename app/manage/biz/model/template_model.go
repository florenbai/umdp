package model

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"umdp/app/manage/biz/dal/mysql"
)

type Template struct {
	Model
	TemplateName string `gorm:"column:template_name" json:"templateName"`
	ProfessionId uint64 `gorm:"column:profession_id"  json:"professionId"`
	Retry        int64  `gorm:"column:retry"  json:"retry"`
}

type TemplateDetail struct {
	Model
	TemplateName   string `gorm:"column:template_name" json:"templateName"`
	ProfessionId   uint64 `gorm:"column:profession_id"  json:"professionId"`
	ProfessionName string `gorm:"column:profession_name"  json:"professionName"`
	Token          string `gorm:"column:token" json:"token"`
	Retry          int64  `gorm:"column:retry"  json:"retry"`
}

func NewTemplateModel() *Template {
	return &Template{}
}

func (m *Template) TableName() string {
	return mysql.TemplateTableName
}

// GetTemplateList 获取模板列表
func (m *Template) GetTemplateList(ctx context.Context, page uint64, pageSize uint64, scopes ...func(*gorm.DB) *gorm.DB) ([]TemplateDetail, int64, error) {
	var i int64
	var list []TemplateDetail
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).
		Joins("LEFT JOIN profession ON template.profession_id = profession.id").
		Scopes(scopes...).Count(&i).Error
	if err != nil {
		return list, i, err
	}
	scopes = append(scopes, Paginate(page, pageSize))
	err = mysql.DB.WithContext(ctx).Table(m.TableName()).
		Select("template.*,profession.profession_name,profession.token").
		Joins("LEFT JOIN profession ON template.profession_id = profession.id").
		Scopes(scopes...).Order("template.id DESC").Find(&list).Error
	if err != nil {
		return list, i, err
	}
	return list, i, nil
}

// 检测重名
func (m *Template) ExistName(ctx context.Context, name string, id uint64) (bool, error) {
	var i int64
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Where("template_name = ? AND id != ?", name, id).Count(&i).Error
	if err != nil {
		return true, err
	}
	if i > 0 {
		return true, nil
	}
	return false, nil
}

// 获取模板配置
func (m *Template) GetTemplateConfig(ctx context.Context, channel string) (*TemplateChannel, error) {
	var data TemplateChannel
	err := mysql.DB.WithContext(ctx).Table(mysql.TemplateChannelTableName).
		Joins("LEFT JOIN channel ON channel.id = template_channel.channel_id").
		Where("template_channel.template_id", m.Id).
		Where("channel.channel_name", channel).First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &data, nil
}

func (m *Template) GetTemplateDetail(ctx context.Context, id uint64) error {
	return mysql.DB.WithContext(ctx).Table(m.TableName()).Where(id, id).First(&m).Error
}

// Delete 删除模板
func (m *Template) Delete(ctx context.Context) error {
	return Transaction(ctx, func(tx *gorm.DB) error {
		err := tx.WithContext(ctx).Table(mysql.TemplateChannelTableName).Where("template_id", m.Id).Delete(&TemplateChannel{}).Error
		if err != nil {
			return err
		}
		return tx.WithContext(ctx).Table(m.TableName()).Delete(&m, "id = ?", m.Id).Error
	})
}
