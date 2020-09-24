// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/24.

package organization

import (
	"goa/model"
)

type GroupSerializer struct {
	model.BaseModel
	Name string `gorm:"not null;comment:'组名称'"`
	Code string `gorm:"not null;comment:'组编码'"`
}
