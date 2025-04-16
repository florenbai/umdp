package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/tidwall/gjson"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"umdp/app/manage/biz/model"
	basepb "umdp/hertz_gen/base"
	pb "umdp/hertz_gen/channel"
	"umdp/pkg/response"
)

type ChannelService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewChannelService(ctx context.Context, c *app.RequestContext) *ChannelService {
	return &ChannelService{ctx: ctx, c: c}
}

// Create 创建渠道
func (service *ChannelService) Create(req pb.ChannelRequest) error {
	cj, _ := req.GetChannelConfig().MarshalJSON()
	if !gjson.Valid(string(cj)) {
		return response.JsonParamErr
	}
	channelModel := model.NewChannelModel()
	count, err := model.GetCountWithScope(service.ctx, channelModel.TableName(), model.WhereWithScope("channel_tag", req.GetChannelTag()))
	if err != nil {
		return err
	}
	if count > 0 {
		return response.ChannelTagHasFoundErr
	}
	status := int8(req.GetChannelStatus())
	return model.Create(service.ctx, channelModel.TableName(), &model.Channel{
		ChannelName:   req.GetChannelName(),
		ChannelTag:    req.GetChannelTag(),
		ChannelConfig: string(cj),
		ChannelStatus: &status,
	})
}

// Edit 编辑渠道
func (service *ChannelService) Edit(id uint64, req pb.ChannelRequest) error {
	channel := model.NewChannelModel()
	err := model.GetOneById(service.ctx, channel.TableName(), id, &channel)
	if err != nil {
		return err
	}
	if channel == nil {
		return response.DataNotFoundErr
	}
	ok, err := channel.ExistChannelByTag(service.ctx, req.GetChannelTag(), id)
	if err != nil {
		return err
	}
	if ok {
		return response.ChannelTagAlreadyExistErr
	}
	cj, _ := req.GetChannelConfig().MarshalJSON()
	if !gjson.Valid(string(cj)) {
		return response.JsonParamErr
	}
	status := int8(req.GetChannelStatus())
	return model.EditOneById(service.ctx, channel.TableName(), id, &model.Channel{
		ChannelName:   req.GetChannelName(),
		ChannelTag:    req.GetChannelTag(),
		ChannelConfig: string(cj),
		ChannelStatus: &status,
	})
}

// List 获取渠道列表
func (service *ChannelService) List(req *basepb.BaseListReq) (*pb.ChannelListResponse, error) {
	var data pb.ChannelListResponse
	var list []model.Channel
	associate := map[string]string{
		"channelName":   "channel_name LIKE ?",
		"channelStatus": "channel_status = ?",
	}
	scopes := model.ParamWithScope(req.Condition.GetCondition(), associate, nil, false)
	count, err := model.GetPageList(service.ctx, model.NewChannelModel().TableName(), req.GetCurrent(), req.GetPageSize(), &list, scopes, model.OrderScope("id desc"))
	if err != nil {
		return nil, err
	}
	for _, v := range list {
		item, err := service.toChannelPB(v)
		if err != nil {
			return nil, err
		}
		data.List = append(data.List, item)
	}
	data.Total = uint64(count)
	return &data, nil
}

// Detail 获取渠道详情
func (service *ChannelService) Detail(id uint64) (*pb.Channel, error) {
	channel := model.NewChannelModel()
	err := model.GetOneById(service.ctx, channel.TableName(), id, &channel)
	if err != nil {
		return nil, err
	}
	return service.toChannelPB(*channel)
}

// All 获取所有渠道
func (service *ChannelService) All() (*pb.ChannelListResponse, error) {
	var data pb.ChannelListResponse
	var list []model.Channel

	err := model.GetAll(service.ctx, model.NewChannelModel().TableName(), &list)
	if err != nil {
		return nil, err
	}
	for _, v := range list {
		item, err := service.toChannelPB(v)
		if err != nil {
			return nil, err
		}
		data.List = append(data.List, item)
	}
	data.Total = uint64(len(data.List))
	return &data, nil
}

// Delete 删除渠道
func (service *ChannelService) Delete(id uint64) error {
	channel := model.NewChannelModel()
	err := model.GetOneById(service.ctx, channel.TableName(), id, &channel)
	if err != nil {
		return err
	}
	return model.DeleteOneById(service.ctx, channel.TableName(), id, &model.Channel{})
}

// toChannelPB 对象转pb
func (service *ChannelService) toChannelPB(v model.Channel) (*pb.Channel, error) {
	configStruct := new(structpb.Struct)
	err := protojson.Unmarshal([]byte(v.ChannelConfig), configStruct)
	if err != nil {
		return nil, err
	}
	item := pb.Channel{
		Id:            v.Id,
		ChannelName:   v.ChannelName,
		ChannelTag:    v.ChannelTag,
		ChannelConfig: configStruct,
		ChannelStatus: int32(*v.ChannelStatus),
		CreatedAt:     v.CreatedAt.Format(),
		UpdatedAt:     v.UpdatedAt.Format(),
	}
	return &item, nil
}

// GetChannelConfigByAgent 根据微信Agentid获取配置信息
func (service *ChannelService) GetChannelConfigByAgent(agentId string) (string, string, error) {
	//var channelData model.Channel
	//err := model.GetOneWithScope(service.ctx, channelData.TableName(), &channelData, model.WhereWithScope("JSON_EXTRACT(channel_config,'$.agentid')", agentId))
	channelModel := model.NewChannelModel()
	err := channelModel.GetChannelByAgentId(service.ctx, agentId)
	if err != nil {
		return "", "", err
	}
	return channelModel.ChannelConfig, channelModel.ChannelName, nil
}
