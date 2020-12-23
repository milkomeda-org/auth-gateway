// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package group

import (
	"auth-gateway/initializer/db"
	"auth-gateway/model"
	"auth-gateway/serializer/group"
)

// CreateService 用户组添加服务
type CreateService struct {
	Name string `form:"name" json:"name" binding:"required"`
	Code string `form:"code" json:"code" binding:"required"`
}

func (receiver CreateService) Execute() error {
	g := model.Group{
		Name: receiver.Name,
		Code: receiver.Code,
	}
	// TODO 检查参数有效性
	return db.DB.Model(&model.Group{}).Save(&g).Error
}

// UpdateService 用户组更新服务
type UpdateService struct {
	ID   int    `form:"id" json:"id" binding:"required"`
	Name string `form:"name" json:"name" binding:"required"`
	Code string `form:"code" json:"code" binding:"required"`
}

func (receiver UpdateService) Execute() error {
	g := model.Group{
		Name: receiver.Name,
		Code: receiver.Code,
	}
	// TODO 检查参数有效性
	return db.DB.Model(&model.Group{}).Where("id = ?", receiver.ID).Updates(&g).Error
}

// DeleteService 用户组删除服务
type DeleteService struct {
	ID int `form:"id" json:"id" binding:"required"`
}

func (receiver DeleteService) Execute() error {
	return db.DB.Where("id = ?", receiver.ID).Unscoped().Delete(&model.Group{}).Error
}

// ViewService 用户组查看服务
type ViewService struct {
}

func (receiver ViewService) Execute() (interface{}, error) {
	var result []group.GroupSerializer
	err := db.DB.Model(&model.Group{}).Find(&result).Error
	return result, err
}
