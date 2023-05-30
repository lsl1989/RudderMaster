package auth

import (
	formAuth "RudderMaster/forms/auth"
	modelAuth "RudderMaster/models/auth"
	svc2 "RudderMaster/service/base"
	"RudderMaster/utils/encryption"
	"errors"
)

func LoginCheck(formData *formAuth.LoginForm) (*modelAuth.User, error) {
	svc := svc2.NewSvc()
	userData, err := svc.FindOne("auth_user", "username", formData.Username, modelAuth.User{})
	if err != nil {
		return nil, err
	}
	user := userData.(*modelAuth.User)
	if encryption.PasswordVerify(user.Password, formData.Password) {
		return user, nil
	}
	return nil, errors.New("用户名密码认证错误")
}
