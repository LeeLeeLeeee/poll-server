package model

import (
	"time"
)

type User struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	UserId      string    `gorm:"not_null;uniqueindex" json:"user_id" binding:"required"`
	Password    string    `gorm:"<-" json:"password" binding:"required"`
	Email       string    `gorm:"<-" json:"email" binding:"required"`
	FirstName   string    `gorm:"<-" json:"firstname" binding:"required"`
	LastName    string    `gorm:"<-" json:"lastname" binding:"required"`
	IsActive    bool      `gorm:"default:true" json:"is_active"`
	LastLogin   time.Time `gorm:"autoCreateTime" json:"last_login"`
	CreatedDate time.Time `gorm:"autoUpdateTime" json:"created_date"`
	UpdatedDate time.Time `gorm:"autoCreateTime" json:"update_date"`
	LoginCount  uint      `gorm:"type:smallint; default:0" json:"login_count"`
	IsSuperuser bool      `json:"is_superuser"`
}

func (User) TableName() string {
	return "tb_user"
}
