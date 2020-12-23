// Copyright The Milkomeda Org. All rights reserved.
// Created by vinson on 2020/11/6.

package role

import (
	rrt2 "auth-gateway/enums/rrt"
	"auth-gateway/initializer/db"
	"auth-gateway/model"
	"auth-gateway/service/rrt"
	"errors"
	"sync"
)

// ProxyService 角色模块服务
type ProxyService struct {
	RoleID  int `form:"roleId" json:"roleId" binding:"required"`
	ProxyID int `form:"proxyId" json:"proxyId" binding:"required"`
}

func (receiver *ProxyService) Add() (err error) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		var count int
		db.DB.Model(&model.Role{}).Where("id = ?", receiver.RoleID).Count(&count)
		if count < 1 {
			err = errors.New("创建失败，角色不存在")
		}
		wg.Done()
	}()
	go func() {
		var count int
		db.DB.Model(&model.Proxy{}).Where("id = ?", receiver.ProxyID).Count(&count)
		if count < 1 {
			err = errors.New("创建失败，代理不存在")
		}
		wg.Done()
	}()
	wg.Wait()
	if nil != err {
		return err
	}
	return rrt.Add(receiver.RoleID, rrt2.RoleProxy, receiver.ProxyID)
}

func (receiver *ProxyService) Remove() error {
	return rrt.Remove(receiver.RoleID, rrt2.RoleProxy, receiver.ProxyID)
}
