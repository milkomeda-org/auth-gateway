// Copyright The ZHIYUN Co. All rights reserved.
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

// UserService 用户组织服务
type OfficeService struct {
	UserID   int `form:"userId" json:"userId" binding:"required"`
	OfficeID int `form:"officeId" json:"officeId" binding:"required"`
}

func (receiver *OfficeService) Add() (err error) {
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
		db.DB.Model(&model.Office{}).Where("id = ?", receiver.OfficeID).Count(&count)
		if count < 1 {
			err = errors.New("创建失败，组织不存在")
		}
		wg.Done()
	}()
	wg.Wait()
	if nil != err {
		return err
	}
	return rrt.Add(receiver.UserID, rrt2.UserOffice, receiver.OfficeID)
}

func (receiver *OfficeService) Remove() error {
	return rrt.Remove(receiver.UserID, rrt2.UserOffice, receiver.OfficeID)
}
