package project

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name     string `gorm:"type:varchar(32);not null;comment:项目名称" json:"name"`
	Alias    string `gorm:"type:varchar(32);not null;comment:项目标识;unique" json:"alias"`
	CreateBy string `gorm:"type:varchar(32);comment:创建者" json:"create_by"`
	UpdateBy string `gorm:"type:varchar(32);comment:更新人" json:"update_by"`
}

func (p Project) TableName() string {
	return "project"
}

type ProjJobRole struct {
	Name  string `gorm:"type:varchar(16);unique;comment:职责名称" json:"name"`
	Alias string `gorm:"type:varchar(32);unique;comment:标识" json:"alias"`
}

func (j ProjJobRole) TableName() string {
	return "proj_job_role"
}

type ProjUser struct {
	gorm.Model
	UserId    uint   `gorm:"default:0;not nll;comment:用户" json:"user_id"`
	ProjectId uint   `gorm:"default:0;not null;comment:项目" json:"project_id"`
	JobRole   string `gorm:"type:varchar(32);not null;comment:工作职责" json:"job_role"`
}

func (u ProjUser) TableName() string {
	return "proj_user"
}
