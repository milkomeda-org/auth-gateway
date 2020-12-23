// Copyright The Milkomeda Org. All rights reserved.
// Created by vinson on 2020/11/6.

package user

import (
	rrt2 "auth-gateway/enums/rrt"
	"auth-gateway/initializer/db"
	"auth-gateway/model"
	"auth-gateway/service/rrt"
	"errors"
	"sync"
)

// UserService 用户角色服务
type RoleService struct {
	UserID int `form:"userId" json:"userId" binding:"required"`
	RoleID int `form:"roleId" json:"roleId" binding:"required"`
}

func (receiver *RoleService) Add() (err error) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		var count int
		db.DB.Model(&model.User{}).Where("id = ?", receiver.UserID).Count(&count)
		if count < 1 {
			err = errors.New("创建失败，用户不存在")
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
	return rrt.Add(receiver.UserID, rrt2.UserRole, receiver.RoleID)
}

func (receiver *RoleService) Remove() error {
	return rrt.Remove(receiver.UserID, rrt2.UserRole, receiver.RoleID)
}
