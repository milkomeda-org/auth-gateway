// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package organization

import (
	"oa-auth/api"
	"oa-auth/serializer"
	"oa-auth/service/organization"

	"github.com/gin-gonic/gin"
)

// OfficeCreate 创建组织
func OfficeCreate(c *gin.Context) serializer.Response {
	var os organization.OfficeCreateService
	if err := c.ShouldBind(&os); err != nil {
		return api.ErrorResponse(err)
	}
	if e := os.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// OfficeUpdate 修改组织
func OfficeUpdate(c *gin.Context) serializer.Response {
	var os organization.OfficeUpdateService
	if err := c.ShouldBind(&os); err != nil {
		return api.ErrorResponse(err)
	}
	if e := os.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// OfficeDelete 删除组织
func OfficeDelete(c *gin.Context) serializer.Response {
	var os organization.OfficeDeleteService
	if err := c.ShouldBind(&os); err != nil {
		return api.ErrorResponse(err)
	}
	if e := os.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// OfficeView 查看组织
func OfficeView(c *gin.Context) serializer.Response {
	var os organization.OfficeViewService
	v, err := os.Execute()
	if nil != err {
		return api.ErrorResponse(err)
	}
	return serializer.Success(v)
}
