// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/11/6.

package user

import (
	"errors"
	rrt2 "oa-auth/enums/rrt"
	"oa-auth/initializer/db"
	"oa-auth/model"
	"oa-auth/service/rrt"
	"sync"
)

// UserService 用户职位服务
type PositionService struct {
	UserID     int `form:"userId" json:"userId" binding:"required"`
	PositionID int `form:"positionId" json:"positionId" binding:"required"`
}

func (receiver *PositionService) Add() (err error) {
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
		db.DB.Model(&model.Position{}).Where("id = ?", receiver.PositionID).Count(&count)
		if count < 1 {
			err = errors.New("创建失败，职位不存在")
		}
		wg.Done()
	}()
	wg.Wait()
	if nil != err {
		return err
	}
	return rrt.Add(receiver.UserID, rrt2.UserPosition, receiver.PositionID)
}

func (receiver *PositionService) Remove() error {
	return rrt.Remove(receiver.UserID, rrt2.UserPosition, receiver.PositionID)
}
