package model

import "time"

type Role struct {
	ID         int    `json:"id"`
	CreateBy   string `gorm:"size:20;comment:创建者名称"`
	UpdateBy   string `gorm:"size:20;comment:修改者名称"`
	Status     int    `gorm:"default:1;comment:状态（1-启用，2-禁用）"`
	RoleName   string `gorm:"size:50;not null;comment:角色名称"`
	RoleCode   string `gorm:"size:50;not null;uniqueIndex;comment:角色code"`
	RoleDesc   string `gorm:"size:100;comment:角色描述"`
	CreateTime time.Time
	UpdateTime time.Time
}
