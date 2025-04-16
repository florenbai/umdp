package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/cloudwego/hertz/pkg/app"
	"umdp/app/manage/biz/model"
	"umdp/app/manage/biz/pack/protodo"
	"umdp/app/manage/model/request"
	"umdp/conf"
	basepb "umdp/hertz_gen/base"
	pb "umdp/hertz_gen/user"
	"umdp/pkg/jwt"
	"umdp/pkg/response"
)

type UserService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewUserService(ctx context.Context, c *app.RequestContext) *UserService {
	return &UserService{ctx: ctx, c: c}
}

// CreateUser 创建用户
func (service *UserService) CreateUser(req request.PutUserRequest) error {
	ok, err := model.NewUserModel().ExistUserByName(service.ctx, req.Nickname)
	if err != nil {
		return err
	}
	if ok {
		return response.NameAlreadyExistErr
	}

	h := md5.New()
	h.Write([]byte(req.Password))
	encryptedPassword := hex.EncodeToString(h.Sum(nil))

	userModel := new(model.User)
	return model.Create(service.ctx, userModel.TableName(), &model.User{
		Nickname: req.Nickname,
		Username: req.Username,
		Password: encryptedPassword,
	})
}

func (service *UserService) VerifyUserInfo(username string, password string) (string, *model.User, error) {
	// 计算密码MD5
	h := md5.New()
	h.Write([]byte(password))
	encryptedPassword := hex.EncodeToString(h.Sum(nil))

	userModel := new(model.User)
	ok, err := userModel.VerifyUser(service.ctx, username, encryptedPassword)
	if err != nil {
		return "", nil, err
	}
	if !ok {
		return "", nil, response.AuthorizeFailErr
	}
	token, err := jwt.GenerateToken(userModel.Id, userModel.Username, []byte(conf.GetConf().Authentication.AuthSecret))
	if err != nil {
		return "", nil, err
	}
	return token, userModel, nil
}

// UpdateUser 编辑用户
func (service *UserService) UpdateUser(id uint64, req request.PutUserRequest) error {
	var user model.User
	err := model.GetOneById(service.ctx, user.TableName(), id, &user)
	if err != nil {
		return err
	}
	if user.Id < 1 {
		return response.DataNotFoundErr
	}

	h := md5.New()
	h.Write([]byte(req.Password))
	encryptedPassword := hex.EncodeToString(h.Sum(nil))

	user.Nickname = req.Nickname
	user.Password = encryptedPassword
	return model.EditOneById(service.ctx, user.TableName(), id, &user)
}

// GetUserList 获取用户列表
func (service *UserService) GetUserList(req *basepb.BaseListReq) (*pb.UserListResponse, error) {
	var res pb.UserListResponse
	associate := map[string]string{
		"nickname": "nickname LIKE ?",
	}
	var userList []model.User
	userModel := new(model.User)
	scopes := model.ParamWithScope(req.Condition.GetCondition(), associate, nil, false)
	c, err := model.GetPageList(service.ctx, userModel.TableName(), req.GetCurrent(), req.GetPageSize(), &userList, scopes)
	if err != nil {
		return nil, err
	}
	err = protodo.CopyWithLocalTime(&res.List, &userList)
	if err != nil {
		return nil, err
	}
	res.Total = uint64(c)
	return &res, nil
}

// GetUserAll 获取所有用户
func (service *UserService) GetUserAll() (*pb.UserListResponse, error) {
	var res pb.UserListResponse
	var userList []model.User
	userModel := new(model.User)
	err := model.GetAll(service.ctx, userModel.TableName(), &userList)
	if err != nil {
		return nil, err
	}
	err = protodo.CopyWithLocalTime(&res.List, &userList)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// DeleteUser 删除用户
func (service *UserService) DeleteUser(id uint64) error {
	var user model.User
	err := model.GetOneById(service.ctx, user.TableName(), id, &user)
	if err != nil {
		return err
	}
	if user.Id < 1 {
		return response.DataNotFoundErr
	}
	return nil
}
