package model

import "time"

type RoleMenu struct {
	ID         int
	RoleId     int `gorm:"not null;comment:角色id"`
	MenuId     int `gorm:"not null;comment:菜单id"`
	CreateTime time.Time
	UpdateTime time.Time
}
