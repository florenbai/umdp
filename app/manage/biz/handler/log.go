package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"umdp/app/manage/biz/service"
	basepb "umdp/hertz_gen/base"
	pb "umdp/hertz_gen/log"
	"umdp/pkg/response"
)

func LogList(ctx context.Context, c *app.RequestContext) {
	pager := response.BasePager(c)
	var err error
	var baseSearch basepb.BaseSearch
	var search pb.SearchRequest
	err = c.Bind(&search)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	data, err := service.NewLogService(ctx, c).List(&basepb.BaseListReq{
		Current:   pager.Page,
		PageSize:  pager.PageSize,
		Condition: &baseSearch,
	}, &search)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}
