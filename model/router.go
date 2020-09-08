package model

// Router 路由模型
type Router struct {
	BaseModel
	Path   string
	Method string
}
