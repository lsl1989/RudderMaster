package auth

type LoginForm struct {
	Username string `json:"username" form:"username" binding:"required,min=2,max=32"`
	Password string `json:"password" form:"password" binding:"required,min=5,max=128"`
}

type LoginUserForm struct {
	ID       uint   `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	RoleId   uint   `json:"role_id" form:"role_id"`
	RoleName string `json:"role_name" form:"role_name"`
	IsAdmin  bool   `json:"is_admin" form:"is_admin"`
}
