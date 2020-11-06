// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/11/6.

package module

import (
	"oa-auth/initializer/db"
	"oa-auth/model"
	"oa-auth/util/snow"
	"strconv"
)

// CreateService 模块创建服务
type CreateService struct {
	Name string `form:"name" json:"name" binding:"required"`
}

// CreateModule 创建模块
func (r *CreateService) Execute() error {
	s, err := snow.SFlake.GetID()
	if nil != err {
		return err
	}
	module := model.Module{
		Name: r.Name,
		Code: strconv.Itoa(int(s)),
	}
	return db.DB.Model(&model.Module{}).Save(&module).Error
}

// DeleteService 模块删除服务
type DeleteService struct {
	ID int `form:"id" json:"id" binding:"required"`
}

// DeleteModule 删除模块
func (r *DeleteService) Execute() error {
	return db.DB.Where("id = ?", r.ID).Delete(model.Module{}).Error
}

// ViewService 查看服务
type ViewService struct {
}

func (receiver ViewService) Execute() (interface{}, error) {
	var r []model.Module
	return r, db.DB.Model(&model.Module{}).Find(&r, "").Error
}
