package role

import (
	"auth-gateway/initializer/db"
	"auth-gateway/model"
)

// CreateService 角色创建服务
type CreateService struct {
	Alias string `form:"alias" json:"alias" binding:"required"`
}

// CreateRole 创建角色
func (r *CreateService) Execute() error {
	role := model.Role{
		Alias: r.Alias,
	}
	return db.DB.Model(&model.Role{}).Save(&role).Error
}

// DeleteService 角色删除服务
type DeleteService struct {
	ID int `form:"id" json:"id" binding:"required"`
}

// DeleteRole 删除角色
func (r *DeleteService) Execute() error {
	return db.DB.Where("id = ?", r.ID).Delete(model.Role{}).Error
}

// ViewService 查看服务
type ViewService struct {
}

func (receiver ViewService) Execute() (interface{}, error) {
	var r []model.Role
	return r, db.DB.Model(&model.Role{}).Find(&r, "").Error
}

func Exists(userName string) bool {
	count := 0
	db.DB.Model(&model.Role{}).Where("alias = ?", userName).Count(&count)
	return count > 0
}
