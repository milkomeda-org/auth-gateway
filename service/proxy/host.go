package proxy

import (
	"auth-gateway/initializer/db"
	"auth-gateway/model"
	"auth-gateway/serializer"
)

// HostRegisterService 路由主机注册服务
type HostRegisterService struct {
	URLScheme string `form:"scheme" json:"scheme" binding:"required,min=1,max=30"`
	URLHost   string `form:"host" json:"host" binding:"required,min=1,max=30"`
	URLPath   string `form:"path" json:"path" binding:"required,min=1,max=30"`
	//Host      string `form:"host" json:"host" binding:"required,min=1,max=30"`
}

// valid 验证表单
func (service *HostRegisterService) valid() *serializer.Response {
	count := 0
	db.DB.Model(service).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "主机已存在",
		}
	}

	return nil
}

// Register 注册路由
func (service *HostRegisterService) Register() *serializer.Response {
	if err := service.valid(); err != nil {
		return err
	}

	proxy := model.ProxyHost{
		URLScheme: service.URLScheme,
		URLHost:   service.URLHost,
		URLPath:   service.URLPath,
	}

	if err := db.DB.Create(&proxy).Error; err != nil {
		return &serializer.Response{Code: 40001, Msg: err.Error()}
	}

	return &serializer.Response{
		Data: proxy,
	}
}

func HostList() *[]model.ProxyHost {
	var rs []model.ProxyHost
	db.DB.Model(model.ProxyHost{}).Find(&rs)
	return &rs
}
