package system

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	RoleName string `gorm:"type:varchar(32);not null;comment:角色名" json:"role_name"`
	Alise    string `gorm:"type:varchar(32);unique;comment:角色标识" json:"alise"`
	CreateBy string `gorm:"type:varchar(32);comment:创建者" json:"create_by"`
	UpdateBy string `gorm:"type:varchar(32);comment:更新人" json:"update_by"`
}

func (Role) TableName() string {
	return "sys_role"
}
