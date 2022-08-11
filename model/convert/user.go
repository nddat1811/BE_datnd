package convert

import (
	"BE_datnd/data"
	"BE_datnd/model/tablemodel"
)

func User(userModel tablemodel.User) *data.User {
	var user data.User
	user.Id = userModel.Id
	user.Name = userModel.Name
	user.Email = userModel.Email
	user.Role = userModel.Role
	user.Password = userModel.Password
	user.Phone = userModel.Phone
	user.CreatedAt = userModel.CreatedAt
	user.UpdatedAt = userModel.UpdatedAt
	user.DeletedAt = userModel.DeletedAt
	return &user
}

func UpdateUser(userModel tablemodel.User) *data.UserUpdateAccount {
	var user data.UserUpdateAccount
	user.Name = userModel.Name
	user.Phone = userModel.Phone

	return &user
}

func Users(t []*tablemodel.User) []*data.User {
	res := make([]*data.User, 0, len(t))
	for _, v := range t {
		res = append(res, User(*v))
	}

	return res
}

func UserChangePassword(userModel tablemodel.User) *data.UserChangePasswordInput {
	var user data.UserChangePasswordInput
	user.Password = userModel.Password

	return &user
}
