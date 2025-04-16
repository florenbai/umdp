package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/tidwall/gjson"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"gorm.io/gorm"
	"umdp/app/manage/biz/model"
	"umdp/app/manage/biz/pack/protodo"
	basepb "umdp/hertz_gen/base"
	pb "umdp/hertz_gen/template"
	"umdp/pkg/response"
)

type TemplateService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewTemplateService(ctx context.Context, c *app.RequestContext) *TemplateService {
	return &TemplateService{ctx: ctx, c: c}
}

// Create 创建业务模板
func (service *TemplateService) Create(req pb.TemplateRequest) error {
	templateModel := model.NewTemplateModel()
	count, err := model.GetCountWithScope(service.ctx, templateModel.TableName(), model.WhereWithScope("template_name", req.GetTemplateName()))
	if err != nil {
		return err
	}
	if count > 0 {
		return response.NameAlreadyExistErr
	}
	templateModel.TemplateName = req.GetTemplateName()
	templateModel.ProfessionId = req.GetProfessionId()
	err = model.Transaction(service.ctx, func(tx *gorm.DB) error {
		tcModel := model.NewTemplateChannelModel()
		err := tx.Table(templateModel.TableName()).Create(&templateModel).Error
		if err != nil {
			return err
		}
		for _, config := range req.GetConfig() {
			conf, err := config.MarshalJSON()
			if err != nil {
				return err
			}
			id := gjson.ParseBytes(conf).Get("id").Uint()
			err = tx.Table(tcModel.TableName()).Create(&model.TemplateChannel{
				TemplateId: templateModel.Id,
				ChannelId:  id,
				Config:     string(conf),
			}).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// Edit 编辑模板
func (service *TemplateService) Edit(id uint64, req pb.TemplateRequest) error {
	templateModel := model.NewTemplateModel()
	err := model.GetOneById(service.ctx, templateModel.TableName(), id, &templateModel)
	if err != nil {
		return err
	}
	ok, err := templateModel.ExistName(service.ctx, req.GetTemplateName(), id)
	if err != nil {
		return err
	}
	if ok {
		return response.ChannelTagAlreadyExistErr
	}
	templateModel.TemplateName = req.GetTemplateName()
	templateModel.ProfessionId = req.GetProfessionId()
	templateModel.Retry = req.GetRetry()
	err = model.Transaction(service.ctx, func(tx *gorm.DB) error {
		tcModel := model.NewTemplateChannelModel()
		err := tx.Table(templateModel.TableName()).Save(&templateModel).Error
		if err != nil {
			return err
		}
		err = tx.Table(tcModel.TableName()).Where("template_id", templateModel.Id).Delete(model.TemplateChannel{}).Error
		if err != nil {
			return err
		}
		for _, config := range req.GetConfig() {
			conf, err := config.MarshalJSON()
			if err != nil {
				return err
			}
			id := gjson.ParseBytes(conf).Get("id").Uint()
			err = tx.Table(tcModel.TableName()).Create(&model.TemplateChannel{
				TemplateId: templateModel.Id,
				ChannelId:  id,
				Config:     string(conf),
			}).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// List 模板列表
func (service *TemplateService) List(req *basepb.BaseListReq) (*pb.TemplateListResponse, error) {
	associate := map[string]string{
		"templateName":   "template_name LIKE ?",
		"professionName": "profession.profession_name LIKE ?",
	}
	var data pb.TemplateListResponse
	scopes := model.ParamWithScope(req.Condition.GetCondition(), associate, nil, false)
	list, count, err := model.NewTemplateModel().GetTemplateList(service.ctx, req.GetCurrent(), req.GetPageSize(), scopes)
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

// Detail 详情
func (service *TemplateService) Detail(id uint64) (*pb.Template, error) {
	template := model.NewTemplateModel()
	err := model.GetOneById(service.ctx, template.TableName(), id, &template)
	if err != nil {
		return nil, err
	}
	list, err := model.NewTemplateChannelModel().GetTemplateChannelConfigsByTemplateId(service.ctx, id)
	if err != nil {
		return nil, err
	}
	var configStructArr []*structpb.Struct
	for _, v := range list {
		configStruct := new(structpb.Struct)
		err = protojson.Unmarshal([]byte(v), configStruct)
		if err != nil {
			return nil, err
		}
		configStructArr = append(configStructArr, configStruct)
	}
	return &pb.Template{
		Id:           template.Id,
		TemplateName: template.TemplateName,
		ProfessionId: template.ProfessionId,
		Retry:        template.Retry,
		CreatedAt:    template.CreatedAt.Format(),
		UpdatedAt:    template.UpdatedAt.Format(),
		Config:       configStructArr,
	}, nil
}

// Delete 删除模板
func (service *TemplateService) Delete(id uint64) error {
	template := model.NewTemplateModel()
	err := model.GetOneById(service.ctx, template.TableName(), id, &template)
	if err != nil {
		return err
	}
	return template.Delete(service.ctx)
}
