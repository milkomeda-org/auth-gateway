// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package organization

import (
	"goa/serializer"
	"goa/service/organization"

	"github.com/gin-gonic/gin"
)

// PositionCreate 创建职位
func PositionCreate(c *gin.Context) serializer.Response {
	var ps organization.PositionAddService
	if err := c.ShouldBind(&ps); err != nil {
		return serializer.ParamErr("", err)

	}
	if e := ps.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// PositionUpdate 修改职位
func PositionUpdate(c *gin.Context) serializer.Response {
	var ps organization.PositionUpdateService
	if err := c.ShouldBind(&ps); err != nil {
		return serializer.ParamErr("", err)
	}
	if e := ps.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// PositionDelete 删除职位
func PositionDelete(c *gin.Context) serializer.Response {
	var ps organization.PositionDeleteService
	if err := c.ShouldBind(&ps); err != nil {
		return serializer.ParamErr("", err)

	}
	if e := ps.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}
