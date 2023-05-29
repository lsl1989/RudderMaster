package service

import (
	"RudderMaster/database"
	"RudderMaster/utils/tools"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"reflect"
)

type Service struct {
	db  *gorm.DB
	ctx context.Context
}

func NewSvc(ctx context.Context) *Service {
	return &Service{
		db:  database.DB.WithContext(ctx),
		ctx: ctx,
	}
}

func (d *Service) FilterTips(m interface{}, search map[string]interface{}, order, orderType string) *gorm.DB {
	// 拼接搜索条件
	query := d.db.Model(m)
	if order != "" && orderType != "" {
		query.Order(fmt.Sprintf("%s %s", order, orderType))
	}
	if len(search) > 0 {
		columns, _ := d.db.Migrator().ColumnTypes(&m)
		for _, field := range columns {
			fieldName := field.Name()
			if value := search[fieldName]; value != nil && value != "" {
				stringVal := tools.Number2String(value)
				query.Where(fmt.Sprintf("%s = '%s'", fieldName, stringVal))
			}
		}
	}
	return query
}

func (d *Service) FilterAll(m interface{}, search map[string]interface{}, order, orderType string) (interface{}, error) {
	query := d.FilterTips(m, search, order, orderType)
	modelType := reflect.TypeOf(m)
	data := reflect.New(reflect.SliceOf(modelType)).Elem().Interface()
	if err := query.Find(&data).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return data, nil
}

func (d *Service) Pagination(m interface{}, size, page int, search map[string]interface{}, order, orderType string) (interface{}, int64, error) {
	var total int64
	var err error
	query := d.FilterTips(m, search, order, orderType)
	modelType := reflect.TypeOf(m)
	limit := tools.GetPageSize(size)
	offset := tools.GetPageNum(page, limit)
	data := reflect.New(reflect.SliceOf(modelType)).Elem().Interface()
	if err = query.Count(&total).Error; err != nil {
		return nil, 0, errors.WithStack(err)
	}
	if size > 0 {
		if err = query.Limit(limit).Offset(offset).Find(&data).Error; err != nil {
			return nil, 0, errors.WithStack(err)
		}
	} else {
		if err = query.Find(&data).Error; err != nil {
			return nil, 0, errors.WithStack(err)
		}
	}
	return data, total, nil
}
