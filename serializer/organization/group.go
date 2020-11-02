// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/24.

package organization

import (
	"oa-auth/model"
)

type GroupSerializer struct {
	model.BaseModel
	Name string `json:"name"`
	Code string `json:"code"`
}
