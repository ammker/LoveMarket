package models2

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}

type UserDetail struct {
	UserID         uint   `gorm:"primaryKey;autoIncrement;column:userID" json:"user_id"` // 用户ID，主键，自增
	ProfilePicture string `gorm:"type:varchar(255);default:null" json:"profile_picture"` // 头像URL或路径
	Username       string `gorm:"type:varchar(255);not null" json:"username"`            // 用户名
	Age            uint   `gorm:"type:int unsigned;default:null" json:"age"`             // 年龄
	Gender         string `gorm:"type:enum('男','女');default:NULL" json:"gender"`         // 性别
	Email          string `gorm:"type:varchar(255);not null;unique" json:"email"`        // 邮箱，唯一
	PhoneNumber    string `gorm:"type:varchar(20);not null" json:"phone_number"`         // 手机号码
	Location       string `gorm:"type:varchar(255);default:null" json:"location"`        // 地区
	School         string `gorm:"type:varchar(255);default:null" json:"school"`          // 学校
	Hobbies        string `gorm:"type:text;default:null" json:"hobbies"`                 // 兴趣爱好
	Signature      string `gorm:"type:text;default:null" json:"signature"`               // 个性签名
	Nickname       string `gorm:"type:varchar(255);default:null" json:"nickname"`        // 昵称
}
