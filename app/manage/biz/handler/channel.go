package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
	"umdp/app/manage/biz/service"
	basepb "umdp/hertz_gen/base"
	pb "umdp/hertz_gen/channel"
	"umdp/pkg/response"
)

// CreateChannel 创建渠道
func CreateChannel(ctx context.Context, c *app.RequestContext) {
	var err error
	var req pb.ChannelRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewChannelService(ctx, c).Create(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

// EditChannel 编辑渠道
func EditChannel(ctx context.Context, c *app.RequestContext) {
	var err error
	var req pb.ChannelRequest
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
	err = service.NewChannelService(ctx, c).Edit(id, req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
	return
}

// ChannelDetail 获取渠道详情
func ChannelDetail(ctx context.Context, c *app.RequestContext) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	data, err := service.NewChannelService(ctx, c).Detail(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

// ListChannel 获取渠道列表
func ChannelList(ctx context.Context, c *app.RequestContext) {
	pager := response.BasePager(c)
	var err error
	var baseSearch basepb.BaseSearch
	baseSearch.Condition = map[string]string{
		"channelName":   c.Query("channelName"),
		"channelStatus": c.Query("channelStatus"),
	}
	data, err := service.NewChannelService(ctx, c).List(&basepb.BaseListReq{
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

// AllChannel 获取所有渠道
func AllChannel(ctx context.Context, c *app.RequestContext) {
	data, err := service.NewChannelService(ctx, c).All()
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

// DeleteChannel 删除渠道
func DeleteChannel(ctx context.Context, c *app.RequestContext) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewChannelService(ctx, c).Delete(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}
