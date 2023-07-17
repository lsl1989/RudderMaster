package application

import "gorm.io/gorm"

// DevLang - 开发语言
type DevLang struct {
	gorm.Model
	Lang string `gorm:"type:varchar(32);not null;comment:项目名称" json:"lang"`
}

func (d DevLang) TableName() string {
	return "dev_lang"
}

type Application struct {
	gorm.Model
	Lang    uint `gorm:"default:0;not nll;comment:开发语言" json:"lang"`
	Project uint `gorm:"default:0;not null;comment:所属项目" json:"project_id"`
}
