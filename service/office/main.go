// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package office

import (
	"auth-gateway/initializer/db"
	"auth-gateway/model"
	"auth-gateway/serializer/office"
	"errors"
	"github.com/lauvinson/gogo/gogo"
)

// CreateService 组织添加服务
type CreateService struct {
	ParentID int    `form:"parentId" json:"parentId"`
	Name     string `form:"name" json:"name" binding:"required"`
	Type     int    `form:"type" json:"type" binding:"required"`
}

func (receiver CreateService) Execute() error {
	if receiver.ParentID != 0 {
		var count int
		db.DB.Model(&model.Office{}).Where("id = ?", receiver.ParentID).Count(&count)
		if count < 1 {
			return errors.New("创建失败，上级不存在")
		}
	}
	o := model.Office{
		ParentID: receiver.ParentID,
		Name:     receiver.Name,
		Type:     receiver.Type,
	}
	return db.DB.Model(&model.Office{}).Save(&o).Error
}

// UpdateService 组织更新服务
type UpdateService struct {
	ID       int    `form:"id" json:"id" binding:"required"`
	ParentID int    `form:"parentId" json:"parentId"`
	Name     string `form:"name" json:"name"`
	Type     int    `form:"type" json:"type"`
}

func (receiver UpdateService) Execute() error {
	o := model.Office{
		ParentID: receiver.ParentID,
		Name:     receiver.Name,
		Type:     receiver.Type,
	}
	// TODO 检查参数有效性
	return db.DB.Model(&model.Office{}).Where("id = ?", receiver.ID).Updates(&o).Error
}

// DeleteService 组织删除服务
type DeleteService struct {
	ID int `form:"id" json:"id" binding:"required"`
}

func (receiver DeleteService) Execute() error {
	return db.DB.Where("id = ?", receiver.ID).Unscoped().Delete(&model.Office{}).Error
}

// ViewService 组织查看服务
type ViewService struct {
}

func (receiver ViewService) Execute() (interface{}, error) {
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
