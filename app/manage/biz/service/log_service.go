package service

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"umdp/app/manage/biz/model"
	"umdp/app/manage/biz/pack/protodo"
	basepb "umdp/hertz_gen/base"
	pb "umdp/hertz_gen/log"
)

type LogService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewLogService(ctx context.Context, c *app.RequestContext) *LogService {
	return &LogService{ctx: ctx, c: c}
}

func (service *LogService) Log(professionId uint64, channelId uint64, templateId uint64, variables string, requestId string, receiver []string, err error) error {
	status := 1
	if err != nil {
		status = 0
	}
	b, err := json.Marshal(receiver)
	if err != nil {
		return err
	}
	log := model.Log{
		ProfessionId: professionId,
		ChannelId:    channelId,
		TemplateId:   templateId,
		Parameters:   variables,
		Receiver:     string(b),
		RequestId:    requestId,
		Status:       int8(status),
	}
	if err != nil {
		log.ErrMessage = err.Error()
	}
	return model.Create(service.ctx, log.TableName(), &log)
}

func (service *LogService) List(req *basepb.BaseListReq, search *pb.SearchRequest) (*pb.LogListResponse, error) {
	var data pb.LogListResponse
	condition := map[string]string{
		"profession": search.GetProfession(),
		"template":   search.GetTemplate(),
		"channel":    search.GetChannel(),
		"status":     search.GetStatus(),
	}
	associate := map[string]string{
		"profession": "profession.profession_name LIKE ?",
		"template":   "template.template_name LIKE ?",
		"channel":    "channel.channel_name LIKE ?",
		"status":     "log.status = ?",
	}
	scopes := model.ParamWithScope(condition, associate, nil, false)
	list, count, err := model.NewLogModel().GetLogList(service.ctx, req.GetCurrent(), req.GetPageSize(), scopes, model.TimeRangeScope("log.created_at", search.GetCreatedAt()))
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
