// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package organization

import (
	"errors"
	"goa/initializer"
	"goa/model/organization"
)

// OfficeAddService 组织添加服务
type OfficeAddService struct {
	ParentID uint   `form:"parentId" json:"parentId" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required"`
	Type     uint   `form:"type" json:"type" binding:"required"`
}

func (receiver OfficeAddService) Execute() error {
	office := organization.Office{
		ParentID: receiver.ParentID,
		Name:     receiver.Name,
		Type:     receiver.Type,
	}
	var count int
	initializer.DB.Model(&organization.Office{}).Where("id = ?", receiver.ParentID).Count(&count)
	if count < 1 {
		return errors.New("创建失败，上级不存在")
	}
	return initializer.DB.Model(&organization.Office{}).Save(&office).Error
}

// OfficeAddService 组织更新服务
type OfficeUpdateService struct {
	ID       uint   `form:"id" json:"id" binding:"required"`
	ParentID uint   `form:"parentId" json:"parentId" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required"`
	Type     uint   `form:"type" json:"type" binding:"required"`
}

func (receiver OfficeUpdateService) Execute() error {
	office := organization.Office{
		ParentID: receiver.ParentID,
		Name:     receiver.Name,
		Type:     receiver.Type,
	}
	// TODO 检查参数有效性
	return initializer.DB.Where("id = ?", receiver.ID).Updates(&office).Error
}

// OfficeAddService 组织删除服务
type OfficeDeleteService struct {
	ID uint `form:"id" json:"id" binding:"required"`
}

func (receiver OfficeDeleteService) Execute() error {
	return initializer.DB.Where("id = ?", receiver.ID).Unscoped().Delete(&organization.Office{}).Error
}
