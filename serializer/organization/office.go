// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/23.

package organization

import (
	"oa-auth/model"

	"github.com/lauvinson/gogo/gogo"
)

type OfficeSerializer struct {
	model.BaseModel
	ParentID uint                `json:"parent_id"` // 上级ID
	Name     string              `json:"name"`      // 组织名称
	Type     uint                `json:"type"`      // 组织类型
	Children []gogo.ForkTreeNode `json:"children"`
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
