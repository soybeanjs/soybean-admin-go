package model

import "time"

type User struct {
	ID         int
	CreateBy   string `gorm:"size:20;comment:创建者名称"`
	UpdateBy   string `gorm:"size:20;comment:修改者名称"`
	Status     int    `gorm:"default:1;comment:状态（1-正常，2-禁用）"`
	UserName   string `gorm:"size:20;not null;comment:用户名"`
	UserGender int    `gorm:"default:1;comment:性别（1-男，2-女）"`
	NickName   string `gorm:"size:20;comment:昵称"`
	UserPhone  string `gorm:"size:11;uniqueIndex;comment:电话"`
	UserEmail  string `gorm:"size:30;comment:邮箱"`
	CreateTime time.Time
	UpdateTime time.Time
}
