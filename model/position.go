// Copyright The Milkomeda Org. All rights reserved.
// Created by vinson on 2020/9/14.

package model

// Position 身份,职位
type Position struct {
	BaseModel
	ParentID int    `gorm:"not null;comment:'上级ID';default:0"` // 上级ID
	OfficeID int    `gorm:"not null;comment:'组织ID'"`           // 组织ID
	Name     string `gorm:"not null;comment:'职位名称'"`           // 职位名称
	Code     string `gorm:"not null;comment:'职位编码'"`           //职位编码
}
