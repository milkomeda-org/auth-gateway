// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/17.

package permission

import (
	"errors"
	"goa/initializer"
	"goa/serializer"

	"github.com/gin-gonic/gin"
)

// AddAccess 添加访问控制
func AddAccess(context *gin.Context) serializer.Response {
	sub, act, obj := context.Request.PostFormValue("sub"),
		context.Request.PostFormValue("act"),
		context.Request.PostFormValue("obj")
	if "" == sub || "" == act || "" == obj {
		return serializer.ParamErr("", errors.New("缺少参数"))
	}
	ok, _ := initializer.Enforcer.AddPolicySafe(sub, act, obj)
	return serializer.Success(ok)
}

// AddAccess 添加访问控制
func RemoveAccess(context *gin.Context) serializer.Response {
	sub, act, obj := context.Query("sub"),
		context.Query("act"),
		context.Query("obj")
	if "" == sub || "" == act || "" == obj {
		return serializer.ParamErr("", errors.New("缺少参数"))
	}
	ok, _ := initializer.Enforcer.RemovePolicySafe(sub, act, obj)
	return serializer.Success(ok)
}
