// Copyright The Milkomeda Org. All rights reserved.
// Created by vinson on 2020/9/14.

package model

// Group 组
type Group struct {
	BaseModel
	Name string `gorm:"not null;comment:'组名称'"`
	Code string `gorm:"not null;comment:'组编码'"`
}
