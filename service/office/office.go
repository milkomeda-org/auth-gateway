// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package office

import (
	"errors"
	"github.com/lauvinson/gogo/gogo"
	"oa-auth/initializer/db"
	office2 "oa-auth/model/office"
	"oa-auth/serializer/office"
)

// OfficeCreateService 组织添加服务
type OfficeCreateService struct {
	ParentID int    `form:"parentId" json:"parentId"`
	Name     string `form:"name" json:"name" binding:"required"`
	Type     int    `form:"type" json:"type" binding:"required"`
}

func (receiver OfficeCreateService) Execute() error {
	office := office2.Office{
		ParentID: receiver.ParentID,
		Name:     receiver.Name,
		Type:     receiver.Type,
	}
	if 0 != receiver.ParentID {
		var count int
		db.DB.Model(&office2.Office{}).Where("id = ?", receiver.ParentID).Count(&count)
		if count < 1 {
			return errors.New("创建失败，上级不存在")
		}
	}
	return db.DB.Model(&office2.Office{}).Save(&office).Error
}

// OfficeAddService 组织更新服务
type OfficeUpdateService struct {
	ID       int    `form:"id" json:"id" binding:"required"`
	ParentID int    `form:"parentId" json:"parentId" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required"`
	Type     int    `form:"type" json:"type" binding:"required"`
}

func (receiver OfficeUpdateService) Execute() error {
	office := office2.Office{
		ParentID: receiver.ParentID,
		Name:     receiver.Name,
		Type:     receiver.Type,
	}
	// TODO 检查参数有效性
	return db.DB.Where("id = ?", receiver.ID).Updates(&office).Error
}

// OfficeAddService 组织删除服务
type OfficeDeleteService struct {
	ID int `form:"id" json:"id" binding:"required"`
}

func (receiver OfficeDeleteService) Execute() error {
	return db.DB.Where("id = ?", receiver.ID).Unscoped().Delete(&office2.Office{}).Error
}

// OfficeViewService 组织查看服务
type OfficeViewService struct {
}

func (receiver OfficeViewService) Execute() (interface{}, error) {
	var result []office.OfficeSerializer
	err := db.DB.Table("offices").Find(&result).Error
	if nil != err {
		return result, err
	}
	var (
		se []gogo.ForkTreeNode
	)
	for _, v := range result {
		temp := v
		se = append(se, &temp)
	}
	a := gogo.BuildTreeByRecursive(se)
	return a, nil
}
