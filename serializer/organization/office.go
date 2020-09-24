// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/23.

package organization

import (
	"goa/model"

	"github.com/lauvinson/gogo/gogo"
)

type OfficeSerializer struct {
	model.BaseModel
	ParentID uint   `gorm:"not null;comment:'上级ID'"` // 上级ID
	Name     string `gorm:"not null;comment:'组织名称'"` // 组织名称
	Type     uint   `gorm:"not null;comment:'组织类型'"` // 组织类型
	Children []gogo.ForkTreeNode
}

func (receiver *OfficeSerializer) GetID() uint {
	return receiver.ID
}

func (receiver *OfficeSerializer) GetPID() uint {
	return receiver.ParentID
}

func (receiver *OfficeSerializer) Following(v []gogo.ForkTreeNode) {
	receiver.Children = append(receiver.Children, v...)
}
