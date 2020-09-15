// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package organization

import (
	"goa/serializer"
	"goa/service/organization"

	"github.com/gin-gonic/gin"
)

// OfficeCreate 创建组织
func OfficeCreate(c *gin.Context) serializer.Response {
	var os organization.OfficeAddService
	if err := c.ShouldBind(&os); err != nil {
		return serializer.ParamErr("", err)

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
		return serializer.ParamErr("", err)
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
		return serializer.ParamErr("", err)

	}
	if e := os.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}
