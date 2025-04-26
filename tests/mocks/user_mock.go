package mocks

import (
	"database/sql"
	"time"

	"github.com/brianmorais/go-user-importation/domain/enums/user_access_type"
	"github.com/brianmorais/go-user-importation/domain/enums/user_type"
	userModel "github.com/brianmorais/go-user-importation/domain/models/user"
)

func GetUsers() userModel.Users {
	return userModel.Users{
		userModel.User{
			UserId:               sql.NullString{String: "user1.test"},
			UserTypeId:           sql.NullInt32{Int32: user_type.ActiveDirectory},
			Password:             sql.NullString{String: ""},
			PasswordReset:        sql.NullBool{Bool: false},
			PasswordModifiedDate: sql.NullTime{Time: time.Date(2021, 10, 1, 12, 0, 0, 0, time.UTC)},
			CreatedDate:          sql.NullTime{Time: time.Date(2021, 10, 1, 12, 0, 0, 0, time.UTC)},
			EmployeeId:           sql.NullInt64{Int64: 20},
			ModifiedDate:         sql.NullTime{Time: time.Date(2021, 10, 1, 12, 0, 0, 0, time.UTC)},
			ModifiedUser:         sql.NullString{String: "Admin"},
			AccessTypeId:         sql.NullInt32{Int32: user_access_type.Employee},
			Locked:               sql.NullBool{Bool: false},
			ExpirationDate:       sql.NullTime{Time: time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC)},
		},
		userModel.User{
			UserId:               sql.NullString{String: "user2.test"},
			UserTypeId:           sql.NullInt32{Int32: user_type.Database},
			Password:             sql.NullString{String: "123456"},
			PasswordReset:        sql.NullBool{Bool: false},
			PasswordModifiedDate: sql.NullTime{Time: time.Date(2021, 10, 1, 12, 0, 0, 0, time.UTC)},
			CreatedDate:          sql.NullTime{Time: time.Date(2021, 10, 1, 12, 0, 0, 0, time.UTC)},
			EmployeeId:           sql.NullInt64{Int64: 22},
			ModifiedDate:         sql.NullTime{Time: time.Date(2021, 10, 1, 12, 0, 0, 0, time.UTC)},
			ModifiedUser:         sql.NullString{String: "Admin"},
			AccessTypeId:         sql.NullInt32{Int32: user_access_type.Excecao},
			Locked:               sql.NullBool{Bool: false},
			ExpirationDate:       sql.NullTime{Time: time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC)},
		},
		userModel.User{
			UserId:               sql.NullString{String: "user3.test"},
			UserTypeId:           sql.NullInt32{Int32: user_type.ActiveDirectory},
			Password:             sql.NullString{String: ""},
			PasswordReset:        sql.NullBool{Bool: false},
			PasswordModifiedDate: sql.NullTime{Time: time.Date(2021, 10, 1, 12, 0, 0, 0, time.UTC)},
			CreatedDate:          sql.NullTime{Time: time.Date(2021, 10, 1, 12, 0, 0, 0, time.UTC)},
			EmployeeId:           sql.NullInt64{Int64: 26},
			ModifiedDate:         sql.NullTime{Time: time.Date(2021, 10, 1, 12, 0, 0, 0, time.UTC)},
			ModifiedUser:         sql.NullString{String: "Admin"},
			AccessTypeId:         sql.NullInt32{Int32: user_access_type.Employee},
			Locked:               sql.NullBool{Bool: false},
			ExpirationDate:       sql.NullTime{Time: time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC)},
		},
	}
}
