// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package organization

import (
	"errors"
	"goa/initializer"
	"goa/model/organization"
)

// PositionAddService 职位添加服务
type PositionAddService struct {
	ParentID uint   `form:"parentId" json:"parentId" binding:"required"`
	OfficeID uint   `form:"officeId" json:"officeId" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required"`
	Code     string `form:"code" json:"code" binding:"required"`
}

func (receiver PositionAddService) Execute() error {
	position := organization.Position{
		ParentID: receiver.ParentID,
		OfficeID: receiver.OfficeID,
		Name:     receiver.Name,
		Code:     receiver.Code,
	}
	var count int
	initializer.DB.Model(&organization.Office{}).Where("id = ?", receiver.OfficeID).Count(&count)
	if count < 1 {
		return errors.New("创建失败，组织不存在")
	}
	count = 0
	initializer.DB.Model(&organization.Position{}).Where("id = ?", receiver.ParentID).Count(&count)
	if count < 1 {
		return errors.New("创建失败，上级不存在")
	}
	return initializer.DB.Model(&organization.Position{}).Save(&position).Error
}

// PositionAddService 职位更新服务
type PositionUpdateService struct {
	ID       uint   `form:"id" json:"id" binding:"required"`
	ParentID uint   `form:"parentId" json:"parentId" binding:"required"`
	OfficeID uint   `form:"officeId" json:"officeId" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required"`
	Code     string `form:"code" json:"code" binding:"required"`
}

func (receiver PositionUpdateService) Execute() error {
	position := organization.Position{
		Name: receiver.Name,
	}
	return initializer.DB.Where("id = ?", receiver.ID).Updates(&position).Error
}

// PositionAddService 职位删除服务
type PositionDeleteService struct {
	ID uint `form:"id" json:"id" binding:"required"`
}

func (receiver PositionDeleteService) Execute() error {
	return initializer.DB.Where("id = ?", receiver.ID).Unscoped().Delete(&organization.Position{}).Error
}
