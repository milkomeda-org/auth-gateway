// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/23.

package organization

import (
	"goa/model"
	"goa/statement"
)

type OfficeSerializer struct {
	model.BaseModel
	ParentID uint   `gorm:"not null;comment:'上级ID'"` // 上级ID
	Name     string `gorm:"not null;comment:'组织名称'"` // 组织名称
	Type     uint   `gorm:"not null;comment:'组织类型'"` // 组织类型
	Children []statement.Sequence
}

func (receiver *OfficeSerializer) GetID() uint {
	return receiver.ID
}

func (receiver *OfficeSerializer) GetParentID() uint {
	return receiver.ParentID
}

func (receiver *OfficeSerializer) AppendChildren(v []statement.Sequence) {
	receiver.Children = append(receiver.Children, v...)
}
