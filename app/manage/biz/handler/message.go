package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"umdp/app/manage/biz/service"
	pb "umdp/hertz_gen/message"
	"umdp/pkg/response"
)

func Send(ctx context.Context, c *app.RequestContext) {
	var err error
	var req pb.SendRequest
	err = c.Bind(&req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	requestId, err := service.NewMessageService(ctx, c).Send(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, map[string]string{"requestId": requestId})
}
