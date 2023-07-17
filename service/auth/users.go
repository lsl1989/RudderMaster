package auth

import (
	"RudderMaster/forms/auth"
	svc2 "RudderMaster/service/base"
)

func GetUserInfo(userId string) (interface{}, error) {
	svc := svc2.NewSvc()
	userInfo, err := svc.FindOne("auth_user", "id", userId, auth.UserDetailForm{})
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}
