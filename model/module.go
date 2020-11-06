// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package model

// Module 模块
type Module struct {
	BaseModel
	Name string `gorm:"not null comment:'模块名称'"` // 模块名称
	Code string `gorm:"not null comment:'模块编码'"` // 模块编码
}
