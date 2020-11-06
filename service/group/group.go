// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package group

import (
	"oa-auth/initializer/db"
	group2 "oa-auth/model/group"
	"oa-auth/serializer/group"
)

// GroupCreateService 用户组添加服务
type GroupCreateService struct {
	Name string `form:"name" json:"name" binding:"required"`
	Code string `form:"code" json:"code" binding:"required"`
}

func (receiver GroupCreateService) Execute() error {
	group := group2.Group{
		Name: receiver.Name,
		Code: receiver.Code,
	}
	// TODO 检查参数有效性
	return db.DB.Model(&group2.Group{}).Save(&group).Error
}

// GroupAddService 用户组更新服务
type GroupUpdateService struct {
	ID   int    `form:"id" json:"id" binding:"required"`
	Name string `form:"name" json:"name" binding:"required"`
	Code string `form:"code" json:"code" binding:"required"`
}

func (receiver GroupUpdateService) Execute() error {
	group := group2.Group{
		Name: receiver.Name,
		Code: receiver.Code,
	}
	// TODO 检查参数有效性
	return db.DB.Where("id = ?", receiver.ID).Updates(&group).Error
}

// GroupAddService 用户组删除服务
type GroupDeleteService struct {
	ID int `form:"id" json:"id" binding:"required"`
}

func (receiver GroupDeleteService) Execute() error {
	return db.DB.Where("id = ?", receiver.ID).Unscoped().Delete(&group2.Group{}).Error
}

// GroupViewService 用户组查看服务
type GroupViewService struct {
}

func (receiver GroupViewService) Execute() (interface{}, error) {
	var result []group.GroupSerializer
	err := db.DB.Table("groups").Find(&result).Error
	return result, err
}
