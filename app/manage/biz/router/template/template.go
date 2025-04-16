package template

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"umdp/app/manage/biz/handler"
)

func Register(r *server.Hertz) {
	root := r.Group("/api/v1")
	_template := root.Group("/template")
	_template.POST("", handler.CreateTemplate)
	_template.GET("/list", handler.TemplateList)
	_template.DELETE("/:id", handler.DeleteTemplate)
	_template.GET("/:id", handler.TemplateDetail)
	_template.PUT("/:id", handler.EditTemplate)
	_template.POST("/test", handler.TestTemplate)
}
