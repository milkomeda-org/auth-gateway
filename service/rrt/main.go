// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/11/6.

package rrt

import (
	"auth-gateway/enums/rrt"
	"auth-gateway/initializer/db"
	"auth-gateway/model"
)

func Add(s int, t rrt.ResRelationType, o int) error {
	r := model.ResRelation{
		S: s,
		T: t,
		O: o,
	}
	return db.DB.Model(&model.ResRelation{}).Save(&r).Error
}

func Remove(s int, t rrt.ResRelationType, o int) error {
	return db.DB.Where("s = ? and t = ? and o = ?", s, t, o).Unscoped().Delete(&model.ResRelation{}).Error
}
