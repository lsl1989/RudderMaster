package auth

import (
	"RudderMaster/utils/encryption"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(32);comment:用户名(唯一);unique" json:"username"`
	Password string `gorm:"type:varchar(128);not null;comment:密码" json:"password"`
	Name     string `gorm:"type:varchar(16);comment:中文名;not null" json:"name"`
	Avatar   string `gorm:"type:varchar(256);comment:头像" json:"avatar"`
	RoleId   uint   `gorm:"default:0;comment:角色ID" json:"role_id"`
	DepId    uint   `gorm:"default:0;comment:部门ID" json:"dep_id"`
	CreateBy string `gorm:"type:varchar(32);comment:创建人" json:"create_by"`
}

func (User) TableName() string {
	return "auth_user"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// 创建前密码加密
	pwd, err := encryption.PasswordHash(u.Password)
	if err != nil {
		errStr := fmt.Sprintf("encryption password failed: %v", err)
		err = errors.New(errStr)
	} else {
		u.Password = pwd
	}
	return
}
