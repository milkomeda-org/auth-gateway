package proxy

import (
	"auth-gateway/initializer/db"
	"auth-gateway/model"
	"auth-gateway/serializer"
)

// RegisterService 路由注册服务
type RegisterService struct {
	Path   string `form:"path" json:"path" binding:"required,min=1,max=30"`
	Method string `form:"method" json:"method" binding:"required,min=1,max=30"`
	HostID int    `form:"hostId" json:"hostId" binding:"required"`
}

// valid 验证表单
func (service *RegisterService) valid() *serializer.Response {
	count := 0
	db.DB.Model(&model.Proxy{}).Where("path = ? and method = ? and host_id = ?", service.Path, service.Method, service.HostID).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "路由已存在",
		}
	}

	return nil
}

// Register 注册路由
func (service *RegisterService) Register() *serializer.Response {
	if err := service.valid(); err != nil {
		return err
	}

	proxy := model.Proxy{
		Path:   service.Path,
		Method: service.Method,
		HostID: service.HostID,
	}

	if err := db.DB.Create(&proxy).Error; err != nil {
		return &serializer.Response{Code: 40001, Msg: err.Error()}
	}

	return &serializer.Response{
		Data: proxy,
	}
}

func List() *[]model.Proxy {
	var rs []model.Proxy
	db.DB.Model(model.Proxy{}).Find(&rs)
	return &rs
}
