// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package office

import (
	"github.com/gin-gonic/gin"
	"oa-auth/serializer"
	"oa-auth/service/office"
)

// Create 创建组织
func Create(c *gin.Context) serializer.Response {
	var os office.OfficeCreateService
	if err := c.ShouldBind(&os); err != nil {
		return serializer.I18Error(err)
	}
	if e := os.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// Update 修改组织
func Update(c *gin.Context) serializer.Response {
	var os office.OfficeUpdateService
	if err := c.ShouldBind(&os); err != nil {
		return serializer.I18Error(err)
	}
	if e := os.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// Delete 删除组织
func Delete(c *gin.Context) serializer.Response {
	var os office.OfficeDeleteService
	if err := c.ShouldBind(&os); err != nil {
		return serializer.I18Error(err)
	}
	if e := os.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// View 查看组织
func View(c *gin.Context) serializer.Response {
	var os office.OfficeViewService
	v, err := os.Execute()
	if nil != err {
		return serializer.I18Error(err)
	}
	return serializer.Success(v)
}
