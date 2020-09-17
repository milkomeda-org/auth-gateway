// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package organization

import (
	"errors"
	"goa/initializer"
	"goa/model/authorization"
	"goa/model/organization"
)

// PositionCreateService 职位添加服务
type PositionCreateService struct {
	ParentID uint   `form:"parentId" json:"parentId" binding:"required"`
	OfficeID uint   `form:"officeId" json:"officeId" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required"`
	Code     string `form:"code" json:"code" binding:"required"`
}

func (receiver PositionCreateService) Execute() error {
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
	position := organization.Position{
		ParentID: receiver.ParentID,
		OfficeID: receiver.OfficeID,
		Name:     receiver.Name,
		Code:     receiver.Code,
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

// PositionRoleMappingAddService 职位角色添加服务
type PositionRoleMappingAddService struct {
	PositionID uint `form:"positionId" json:"positionId" binding:"required"`
	RoleID     uint `form:"roleId" json:"roleId" binding:"required"`
}

func (receiver PositionRoleMappingAddService) Execute() error {
	var count = 0
	initializer.DB.Model(&organization.Position{}).Where("id = ?", receiver.PositionID).Count(&count)
	if count < 1 {
		return errors.New("创建失败，身份不存在")
	}
	count = 0
	initializer.DB.Model(&authorization.Role{}).Where("id = ?", receiver.RoleID).Count(&count)
	if count < 1 {
		return errors.New("创建失败，角色不存在")
	}
	positionRole := organization.PositionRoleMapping{
		PositionID: receiver.PositionID,
		RoleID:     receiver.RoleID,
	}
	return initializer.DB.Model(&organization.PositionRoleMapping{}).Save(&positionRole).Error
}

// PositionRoleMappingRemoveService 职位角色删除服务
type PositionRoleMappingRemoveService struct {
	ID uint `form:"id" json:"id" binding:"required"`
}

func (receiver PositionRoleMappingRemoveService) Execute() error {
	return initializer.DB.Where("id = ?", receiver.ID).Unscoped().Delete(&organization.PositionRoleMapping{}).Error
}
