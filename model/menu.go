package model

type Menu struct {
	ID         int    `json:"id"`
	CreateBy   string `gorm:"size:20;comment:创建者名称"`
	UpdateBy   string `gorm:"size:20;comment:修改者名称"`
	Status     string `gorm:"default:1;comment:状态（1-启用，2-禁用）"`
	ParentID   int    `gorm:"default:1;comment:状态（1-启用，2-禁用）"`
	MenuType   string `gorm:"default:1;comment:类型（1-目录，2-菜单）"`
	MenuName   string `gorm:"not null;comment:菜单名称"`
	RouteName  string `gorm:"not null;comment:路由名称"`
	RoutePath  string `gorm:"not null;comment:路由路径"`
	Component  string `gorm:"comment:组件路径"`
	Order      int    `gorm:"comment:排序"`
	I18nKey    string `gorm:"comment:排序"`
	Icon       string `gorm:"comment:Icon"`
	IconType   string `gorm:"default:1;comment:图标类型（1-iconify图标，2-本地图标）"`
	HideInMenu bool   `gorm:"default:false;comment:隐藏菜单"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}
