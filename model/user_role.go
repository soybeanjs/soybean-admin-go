package model

import "time"

type UserRole struct {
	ID         int
	RoleId     int `gorm:"not null;comment:角色id"`
	UserId     int `gorm:"not null;comment:用户id"`
	CreateTime time.Time
	UpdateTime time.Time
}
