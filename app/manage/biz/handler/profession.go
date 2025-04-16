package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
	"umdp/app/manage/biz/service"
	basepb "umdp/hertz_gen/base"
	pb "umdp/hertz_gen/profession"
	"umdp/pkg/response"
)

// CreateProfession 创建业务
func CreateProfession(ctx context.Context, c *app.RequestContext) {
	var err error
	var req pb.ProfessionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewProfessionService(ctx, c).Create(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

// EditProfession 编辑业务
func EditProfession(ctx context.Context, c *app.RequestContext) {
	var err error
	var req pb.ProfessionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewProfessionService(ctx, c).Edit(id, req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
	return
}

// ListProfession 业务列表
func ListProfession(ctx context.Context, c *app.RequestContext) {
	pager := response.BasePager(c)
	var err error
	var baseSearch basepb.BaseSearch
	baseSearch.Condition = map[string]string{
		"professionName": c.Query("professionName"),
		"channel":        c.Query("channel"),
	}
	data, err := service.NewProfessionService(ctx, c).List(&basepb.BaseListReq{
		Current:   pager.Page,
		PageSize:  pager.PageSize,
		Condition: &baseSearch,
	})
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

// AllProfession 获取所有业务
func AllProfession(ctx context.Context, c *app.RequestContext) {
	data, err := service.NewProfessionService(ctx, c).All()
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

// ProfessionChannelsMap 获取业务渠道
func ProfessionChannelsMap(ctx context.Context, c *app.RequestContext) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	data, err := service.NewProfessionService(ctx, c).ChannelsMap(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

// ProfessionChannels 获取业务渠道信息
func ProfessionChannels(ctx context.Context, c *app.RequestContext) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	data, err := service.NewProfessionService(ctx, c).Channels(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

// ProfessionDetail 获取业务详情
func ProfessionDetail(ctx context.Context, c *app.RequestContext) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	data, err := service.NewProfessionService(ctx, c).Detail(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

// DeleteProfession 删除业务
func DeleteProfession(ctx context.Context, c *app.RequestContext) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewProfessionService(ctx, c).Delete(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}
