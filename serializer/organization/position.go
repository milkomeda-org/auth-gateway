// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/24.

package organization

import (
	"goa/model"

	"github.com/lauvinson/gogo/gogo"
)

type PositionSerializer struct {
	model.BaseModel
	ParentID uint   `gorm:"not null;comment:'上级ID'"` // 上级ID
	OfficeID uint   `gorm:"not null;comment:'组织ID'"` // 组织ID
	Name     string `gorm:"not null;comment:'职位名称'"` // 职位名称
	Code     string `gorm:"not null;comment:'职位编码'"` //职位编码
	Children []gogo.ForkTreeNode
}

func (receiver *PositionSerializer) GetID() uint {
	return receiver.ID
}

func (receiver *PositionSerializer) GetPID() uint {
	return receiver.ParentID
}

func (receiver *PositionSerializer) Following(v []gogo.ForkTreeNode) {
	receiver.Children = append(receiver.Children, v...)
}
