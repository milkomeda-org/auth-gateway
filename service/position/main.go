// Copyright The Milkomeda Org. All rights reserved.
// Created by vinson on 2020/9/14.

package position

import (
	"auth-gateway/initializer/db"
	"auth-gateway/model"
	"auth-gateway/serializer/position"
	"errors"
	"github.com/lauvinson/gogo/gogo"
	"sync"
)

// CreateService 职位添加服务
type CreateService struct {
	ParentID int    `form:"parentId" json:"parentId" binding:"required"`
	OfficeID int    `form:"officeId" json:"officeId" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required"`
	Code     string `form:"code" json:"code" binding:"required"`
}

func (receiver CreateService) Execute() (err error) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		var count int
		db.DB.Model(&model.Office{}).Where("id = ?", receiver.OfficeID).Count(&count)
		if count < 1 {
			err = errors.New("创建失败，组织不存在")
		}
		wg.Done()
	}()
	go func() {
		var count = 0
		db.DB.Model(&model.Position{}).Where("id = ?", receiver.ParentID).Count(&count)
		if count < 1 {
			err = errors.New("创建失败，上级不存在")
		}
		wg.Done()
	}()
	wg.Wait()
	if nil != err {
		return err
	}
	p := model.Position{
		ParentID: receiver.ParentID,
		OfficeID: receiver.OfficeID,
		Name:     receiver.Name,
		Code:     receiver.Code,
	}
	return db.DB.Model(&model.Position{}).Save(&p).Error
}

// UpdateService 职位更新服务
type UpdateService struct {
	ID       int    `form:"id" json:"id" binding:"required"`
	ParentID int    `form:"parentId" json:"parentId" binding:"required"`
	OfficeID int    `form:"officeId" json:"officeId" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required"`
	Code     string `form:"code" json:"code" binding:"required"`
}

func (receiver UpdateService) Execute() error {
	p := model.Position{
		Name: receiver.Name,
	}
	return db.DB.Model(&model.Position{}).Where("id = ?", receiver.ID).Updates(&p).Error
}

// DeleteService 职位删除服务
type DeleteService struct {
	ID int `form:"id" json:"id" binding:"required"`
}

func (receiver DeleteService) Execute() error {
	return db.DB.Model(&model.Position{}).Where("id = ?", receiver.ID).Unscoped().Delete(&model.Position{}).Error
}

// ViewService 职位查看服务
type ViewService struct {
	OfficeID int `form:"officeId" json:"officeId" binding:"required"`
}

func (receiver ViewService) Execute() (interface{}, error) {
	var result []position.PositionSerializer
	err := db.DB.Table("positions").Find(&result, "office_id = ?", receiver.OfficeID).Error
	if nil != err {
		return result, err
	}
	se := make([]gogo.ForkTreeNode, 0)
	for _, v := range result {
		temp := v
		se = append(se, &temp)
	}
	a := gogo.BuildTreeByRecursive(se)
	return a, nil
}
