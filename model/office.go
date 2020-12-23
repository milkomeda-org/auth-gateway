// Copyright The Milkomeda Org. All rights reserved.
// Created by vinson on 2020/9/14.

package model

// Office 组织
type Office struct {
	BaseModel
	ParentID int    `gorm:"not null;comment:'上级ID';default:0"` // 上级ID
	Name     string `gorm:"not null;comment:'组织名称'"`           // 组织名称
	Type     int    `gorm:"not null;comment:'组织类型'"`           // 组织类型
}
