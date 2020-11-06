// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/11/6.

package group

import (
	"errors"
	rrt2 "oa-auth/enums/rrt"
	"oa-auth/initializer/db"
	"oa-auth/model"
	"oa-auth/service/rrt"
	"sync"
)

// RoleService 组织角色服务
type RoleService struct {
	GroupID int `form:"groupId" json:"groupId" binding:"required"`
	RoleID  int `form:"roleId" json:"roleId" binding:"required"`
}

func (receiver *RoleService) Add() (err error) {
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
		db.DB.Model(&model.Role{}).Where("id = ?", receiver.RoleID).Count(&count)
		if count < 1 {
			err = errors.New("创建失败，角色不存在")
		}
		wg.Done()
	}()
	wg.Wait()
	if nil != err {
		return err
	}
	return rrt.Add(receiver.GroupID, rrt2.GroupRole, receiver.RoleID)
}

func (receiver *RoleService) Remove() error {
	return rrt.Remove(receiver.GroupID, rrt2.GroupRole, receiver.RoleID)
}
