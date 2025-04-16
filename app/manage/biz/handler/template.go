package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
	"umdp/app/manage/biz/service"
	basepb "umdp/hertz_gen/base"
	messagepb "umdp/hertz_gen/message"
	pb "umdp/hertz_gen/template"
	"umdp/pkg/response"
)

// CreateTemplate 创建模板
func CreateTemplate(ctx context.Context, c *app.RequestContext) {
	var err error
	var req pb.TemplateRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewTemplateService(ctx, c).Create(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

// EditTemplate 编辑模板
func EditTemplate(ctx context.Context, c *app.RequestContext) {
	var err error
	var req pb.TemplateRequest
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
	err = service.NewTemplateService(ctx, c).Edit(id, req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
	return
}

// TemplateDetail 模板详情
func TemplateDetail(ctx context.Context, c *app.RequestContext) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	data, err := service.NewTemplateService(ctx, c).Detail(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

// TemplateList 模板列表
func TemplateList(ctx context.Context, c *app.RequestContext) {
	pager := response.BasePager(c)
	var err error
	var baseSearch basepb.BaseSearch
	baseSearch.Condition = map[string]string{
		"templateName":   c.Query("templateName"),
		"professionName": c.Query("professionName"),
	}
	data, err := service.NewTemplateService(ctx, c).List(&basepb.BaseListReq{
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

// DeleteTemplate 删除模板
func DeleteTemplate(ctx context.Context, c *app.RequestContext) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewTemplateService(ctx, c).Delete(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

func TestTemplate(ctx context.Context, c *app.RequestContext) {
	var err error
	var req messagepb.SendRequest
	err = c.Bind(&req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewMessageService(ctx, c).SendTest(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}
