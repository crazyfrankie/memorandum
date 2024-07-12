package model

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primary"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
	Name      string     `json:"name" gorm:"unique"`
	Password  string     `json:"password"`
}

type LoginData struct {
	Name     string `json:"name" binding:"required,min=3,max=15" example:"frank"`
	Password string `json:"password" binding:"required,min=5,max=16" example:"frank666"`
}

type UserResp struct {
	ID       uint   `json:"id" form:"id" example:"1"`                    // 用户ID
	UserName string `json:"user_name" form:"user_name" example:"FanOne"` // 用户名
	CreateAt int64  `json:"create_at" form:"create_at"`                  // 创建
}
