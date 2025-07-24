// Package model 提供用户账户数据模型定义
package model

import "lease/internal/model/base"

// Account 用户账户模型
type Account struct {
	base.Base
	Phone    string `gorm:"type:varchar(32);unique;default:null" json:"phone"` // 手机号，次登录方式
	Email    string `gorm:"type:varchar(64);unique;not null" json:"email"`     // 邮箱，主登录方式
	Password string `gorm:"type:varchar(255);not null" json:"password"`        // 加密密码
	Avatar   string `gorm:"type:varchar(255);default:null" json:"avatar"`      // 用户头像
	Nickname string `gorm:"type:varchar(64);not null" json:"nickname"`         // 昵称
	Status   bool   `gorm:"type:boolean;default:false" json:"status"`          // 账号状态
}

// TableName 指定表名
// 返回值：
//   - string: 表名
func (Account) TableName() string {
	return "accounts"
}
