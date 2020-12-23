// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package office

import (
	"auth-gateway/serializer"
	"auth-gateway/service/office"
	"github.com/gin-gonic/gin"
)

// Create 创建组织
func Create(c *gin.Context) *serializer.Response {
	var os office.CreateService
	if err := c.ShouldBind(&os); err != nil {
		return serializer.I18Error(err)
	}
	if e := os.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// Update 修改组织
func Update(c *gin.Context) *serializer.Response {
	var os office.UpdateService
	if err := c.ShouldBind(&os); err != nil {
		return serializer.I18Error(err)
	}
	if e := os.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// Delete 删除组织
func Delete(c *gin.Context) *serializer.Response {
	var os office.DeleteService
	if err := c.ShouldBind(&os); err != nil {
		return serializer.I18Error(err)
	}
	if e := os.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// View 查看组织
func View(c *gin.Context) *serializer.Response {
	var os office.ViewService
	v, err := os.Execute()
	if nil != err {
		return serializer.I18Error(err)
	}
	return serializer.Success(v)
}
