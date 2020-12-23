// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/17.

package access

import (
	"auth-gateway/initializer/db"
	"auth-gateway/serializer"
	"errors"

	"github.com/gin-gonic/gin"
)

// Add 添加访问控制
func Add(context *gin.Context) *serializer.Response {
	sub, act, obj := context.Request.PostFormValue("sub"),
		context.Request.PostFormValue("act"),
		context.Request.PostFormValue("obj")
	if sub == "" || act == "" || obj == "" {
		return serializer.ParamErr("", errors.New("缺少参数"))
	}
	ok, _ := db.Enforcer.AddPolicySafe(sub, act, obj)
	return serializer.Success(ok)
}

// Remove 添加访问控制
func Remove(context *gin.Context) *serializer.Response {
	sub, act, obj := context.Query("sub"),
		context.Query("act"),
		context.Query("obj")
	if sub == "" || act == "" || obj == "" {
		return serializer.ParamErr("", errors.New("缺少参数"))
	}
	ok, _ := db.Enforcer.RemovePolicySafe(sub, act, obj)
	return serializer.Success(ok)
}
