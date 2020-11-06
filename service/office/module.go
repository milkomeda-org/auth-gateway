// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/11/6.

package office

import (
	"errors"
	rrt2 "oa-auth/enums/rrt"
	"oa-auth/initializer/db"
	"oa-auth/model"
	"oa-auth/service/rrt"
	"sync"
)

// RoleService 组织角色服务
type ModuleService struct {
	OfficeID int `form:"officeId" json:"officeId" binding:"required"`
	ModuleID int `form:"moduleId" json:"moduleId" binding:"required"`
}

func (receiver *ModuleService) Add() (err error) {
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
		var count int
		db.DB.Model(&model.Module{}).Where("id = ?", receiver.ModuleID).Count(&count)
		if count < 1 {
			err = errors.New("创建失败，模块不存在")
		}
		wg.Done()
	}()
	wg.Wait()
	if nil != err {
		return err
	}
	return rrt.Add(receiver.OfficeID, rrt2.OfficeModule, receiver.ModuleID)
}

func (receiver *ModuleService) Remove() error {
	return rrt.Remove(receiver.OfficeID, rrt2.OfficeModule, receiver.ModuleID)
}
