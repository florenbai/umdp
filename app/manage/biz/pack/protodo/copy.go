package protodo

import (
	"errors"
	"github.com/jinzhu/copier"
	"umdp/app/manage/biz/model"
)

func CopyWithLocalTime(to interface{}, from interface{}) error {
	return copier.CopyWithOption(to, from, copier.Option{IgnoreEmpty: true, DeepCopy: true, Converters: []copier.TypeConverter{
		{
			SrcType: model.LocalTime{},
			DstType: copier.String,
			Fn: func(src interface{}) (interface{}, error) {
				s, ok := src.(model.LocalTime)
				if !ok {
					return nil, errors.New("src type not matching")
				}
				return s.Format(), nil
			},
		},
	}})
}
