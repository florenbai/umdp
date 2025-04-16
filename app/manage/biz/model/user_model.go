package model

import (
	"context"
	"gorm.io/gorm"
	"umdp/app/manage/biz/dal/mysql"
)

type User struct {
	Model
	Nickname string `gorm:"column:nickname" json:"nickname" `
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUserModel() *User {
	return &User{}
}

func (m *User) TableName() string {
	return mysql.UserTableName
}

// ExistUserByName 检测用户是否存在
func (m *User) ExistUserByName(ctx context.Context, name string) (bool, error) {
	var i int64
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Where("nickname = ?", name).Count(&i).Error
	if err != nil {
		return true, err
	}
	if i > 0 {
		return true, nil
	}
	return false, nil
}

func (m *User) VerifyUser(ctx context.Context, username string, password string) (bool, error) {
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Where("username = ? AND password = ?", username, password).First(&m).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (m *User) GetUserById(ctx context.Context, id uint64) (User, error) {
	var user User
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Where("id = ?", id).Take(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
