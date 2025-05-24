package interfaces

import userModel "github.com/brianmorais/go-user-importation/domain/models/user"

type IUserRepository interface {
	FindUserById(userId string) (userModel.User, error)
	CreateUser(user userModel.User) (int64, error)
	UpdateUser(user userModel.User) (int64, error)
	FindUsersByEmployeeId(employeeId int64) (userModel.Users, error)
	DeleteUser(userId string) (int64, error)
	GetUsersByTypeAndEmployeeId(employeeId int64, userTypeId int32) (userModel.Users, error)
}
