// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package common

import (
	"oa-auth/serializer"
	"oa-auth/statement"

	"github.com/gin-gonic/gin"
)

// Controller 通用controller
func Controller(c *gin.Context, execute statement.Execute) serializer.Response {
	if e := (execute).Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}
