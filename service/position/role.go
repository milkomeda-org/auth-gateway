// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/11/6.

package position

import (
	rrt2 "auth-gateway/enums/rrt"
	"auth-gateway/initializer/db"
	"auth-gateway/model"
	"auth-gateway/service/rrt"
	"errors"
	"sync"
)

// RoleService 职位角色服务
type RoleService struct {
	PositionID int `form:"positionId" json:"positionId" binding:"required"`
	RoleID     int `form:"roleId" json:"roleId" binding:"required"`
}

func (receiver *RoleService) Add() (err error) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		var count = 0
		db.DB.Model(&model.Position{}).Where("id = ?", receiver.PositionID).Count(&count)
		if count < 1 {
			err = errors.New("创建失败，职位不存在")
		}
		wg.Done()
	}()
	go func() {
		count := 0
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
	return rrt.Add(receiver.PositionID, rrt2.PositionRole, receiver.RoleID)
}

func (receiver *RoleService) Remove() error {
	return rrt.Remove(receiver.PositionID, rrt2.PositionRole, receiver.RoleID)
}
