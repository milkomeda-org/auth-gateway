// Copyright The Milkomeda Org. All rights reserved.
// Created by vinson on 2020/11/6.

package group

import (
	rrt2 "auth-gateway/enums/rrt"
	"auth-gateway/initializer/db"
	"auth-gateway/model"
	"auth-gateway/service/rrt"
	"errors"
	"sync"
)

// RoleService 组织角色服务
type ModuleService struct {
	GroupID  int `form:"groupId" json:"groupId" binding:"required"`
	ModuleID int `form:"moduleId" json:"moduleId" binding:"required"`
}

func (receiver *ModuleService) Add() (err error) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		var count int
		db.DB.Model(&model.Group{}).Where("id = ?", receiver.GroupID).Count(&count)
		if count < 1 {
			err = errors.New("创建失败，用户组不存在")
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
	return rrt.Add(receiver.GroupID, rrt2.GroupModule, receiver.ModuleID)
}

func (receiver *ModuleService) Remove() error {
	return rrt.Remove(receiver.GroupID, rrt2.GroupModule, receiver.ModuleID)
}
