package auth

type UserForm struct {
	Name     string `json:"name" form:"name" binding:"required,only_chinese,min=2,max=16" example:"张三"`
	Username string `json:"username" form:"username" binding:"required,min=2,max=32" example:"zhangsan"`
	Avatar   string `json:"avatar"`
	RoleId   uint   `json:"role_id" form:"role_id" binding:"required"`
	DepId    uint   `json:"dep_id" form:"dep_id"`
	CreateBy string `json:"create_by" form:"create_by"`
	IsAdmin  bool   `json:"is_admin" form:"is_admin"`
}

type UserDetailForm struct {
	UserForm
	ID        uint   `json:"id"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}
