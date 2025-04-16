package profession

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"umdp/app/manage/biz/handler"
)

func Register(r *server.Hertz) {
	root := r.Group("/api/v1")
	_profession := root.Group("/profession")
	_profession.POST("", handler.CreateProfession)
	_profession.GET("/list", handler.ListProfession)
	_profession.PUT("/:id", handler.EditProfession)
	_profession.GET("/:id", handler.ProfessionDetail)
	_profession.GET("/channels/:id", handler.ProfessionChannels)
	_profession.GET("/channels-map/:id", handler.ProfessionChannelsMap)
	_profession.GET("/all", handler.AllProfession)
	_profession.DELETE("/:id", handler.DeleteProfession)
}
