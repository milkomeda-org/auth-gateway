package serializer

import (
	"oa-auth/model/organization"
)

// User 用户序列化器
type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Nickname  string `json:"nickname"`
	Status    int    `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

// BuildUser 序列化用户
func BuildUser(user organization.User) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		Nickname:  user.Nickname,
		Status:    user.Status,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user UserSession) Response {
	return Response{
		Data: user,
	}
}

// UserSession 用户session信息
type UserSession struct {
	UserName string         `json:"user_name"`
	UserID   uint           `json:"user_id"`
	NickName string         `json:"nick_name"`
	Avatar   string         `json:"avatar"`
	Roles    map[int]string `json:"roles"`
}
