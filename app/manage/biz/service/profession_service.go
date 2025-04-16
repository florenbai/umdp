package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/jinzhu/copier"
	"umdp/app/manage/biz/model"
	"umdp/app/manage/biz/pack/protodo"
	basepb "umdp/hertz_gen/base"
	channelpb "umdp/hertz_gen/channel"
	pb "umdp/hertz_gen/profession"
	"umdp/pkg/response"
)

type ProfessionService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewProfessionService(ctx context.Context, c *app.RequestContext) *ProfessionService {
	return &ProfessionService{ctx: ctx, c: c}
}

// Create 创建业务
func (service *ProfessionService) Create(req pb.ProfessionRequest) error {
	professionModel := model.NewProfessionModel()
	count, err := model.GetCountWithScope(service.ctx, professionModel.TableName(), model.WhereWithScope("profession_name", req.GetProfessionName()))
	if err != nil {
		return err
	}
	if count > 0 {
		return response.NameAlreadyExistErr
	}
	professionModel.ProfessionName = req.GetProfessionName()
	professionModel.Token = req.GetToken()
	return professionModel.CreateProfession(service.ctx, req.GetChannels())
}

// Edit 编辑业务
func (service *ProfessionService) Edit(id uint64, req pb.ProfessionRequest) error {
	profession := model.NewProfessionModel()
	err := model.GetOneById(service.ctx, profession.TableName(), id, &profession)
	if err != nil {
		return err
	}
	ok, err := profession.ExistName(service.ctx, req.GetProfessionName(), id)
	if err != nil {
		return err
	}
	if ok {
		return response.ChannelTagAlreadyExistErr
	}
	profession.ProfessionName = req.GetProfessionName()
	profession.Token = req.GetToken()
	return profession.EditProfession(service.ctx, req.GetChannels())
}

// List 获取业务列表
func (service *ProfessionService) List(req *basepb.BaseListReq) (*pb.ProfessionListResponse, error) {
	var data pb.ProfessionListResponse
	associate := map[string]string{
		"professionName": "profession.profession_name LIKE ?",
		"channel":        "profession.id IN (SELECT profession_id FROM profession_channel WHERE channel_id = ?)",
	}
	scopes := model.ParamWithScope(req.Condition.GetCondition(), associate, nil, false)
	list, count, err := model.NewProfessionModel().GetProfessionList(service.ctx, req.GetCurrent(), req.GetPageSize(), scopes)
	if err != nil {
		return nil, err
	}
	protodo.CopyWithLocalTime(&data.List, &list)
	if err != nil {
		return nil, err
	}
	data.Total = uint64(count)
	return &data, nil
}

// Channels 获取业务渠道map
func (service *ProfessionService) Channels(id uint64) ([]*channelpb.ChannelMap, error) {
	var data []*channelpb.ChannelMap
	profession := model.NewProfessionModel()
	err := model.GetOneById(service.ctx, profession.TableName(), id, &profession)
	if err != nil {
		return nil, err
	}
	channels, err := model.NewProfessionChannelModel().GetChannelsDetailByProfession(service.ctx, id)
	if err != nil {
		return nil, err
	}
	err = copier.CopyWithOption(&data, &channels, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (service *ProfessionService) ChannelsMap(id uint64) (map[uint64]model.Channel, error) {
	var data = make(map[uint64]model.Channel)
	profession := model.NewProfessionModel()
	err := model.GetOneById(service.ctx, profession.TableName(), id, &profession)
	if err != nil {
		return nil, err
	}
	channels, err := model.NewProfessionChannelModel().GetChannelsDetailByProfession(service.ctx, id)
	if err != nil {
		return nil, err
	}
	for _, v := range channels {
		data[v.Id] = v
	}
	return data, nil
}

// All 获取所有业务
func (service *ProfessionService) All() ([]*pb.Profession, error) {
	var professions []model.Profession
	var data []*pb.Profession
	err := model.GetAll(service.ctx, model.NewProfessionModel().TableName(), &professions)
	if err != nil {
		return nil, err
	}
	protodo.CopyWithLocalTime(&data, &professions)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Detail 获取渠道详情
func (service *ProfessionService) Detail(id uint64) (*pb.ProfessionDetail, error) {
	profession := model.NewProfessionModel()
	err := model.GetOneById(service.ctx, profession.TableName(), id, &profession)
	if err != nil {
		return nil, err
	}
	channels, err := model.NewProfessionChannelModel().GetChannelsByProfession(service.ctx, id)
	if err != nil {
		return nil, err
	}
	return &pb.ProfessionDetail{
		Id:             profession.Id,
		ProfessionName: profession.ProfessionName,
		Token:          profession.Token,
		Channels:       channels,
		CreatedAt:      profession.CreatedAt.Format(),
		UpdatedAt:      profession.UpdatedAt.Format(),
	}, nil
}

// Delete 删除业务
func (service *ProfessionService) Delete(id uint64) error {
	profession := model.NewProfessionModel()
	err := model.GetOneById(service.ctx, profession.TableName(), id, &profession)
	if err != nil {
		return err
	}
	return profession.Delete(service.ctx)
}
