package model

import (
	"oa-auth/enums/rrt"
	"time"
)

// BaseModel model base
type BaseModel struct {
	ID        int        `gorm:"primary_key;unique" json:"id"`
	CreatedAt time.Time  `gorm:"type:datetime;not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"type:datetime;not null" json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Status    int        `gorm:"not null;default:0" json:"status"`
}

// ResRelation 资源关联
type ResRelation struct {
	BaseModel
	S int                 `gorm:"not null;comment:'主体'"`
	T rrt.ResRelationType `gorm:"not null;comment:'类型'"`
	O int                 `gorm:"not null;comment:'受体'"`
}
