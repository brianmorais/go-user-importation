package userModel

import (
	"database/sql"
	"time"

	"github.com/brianmorais/go-user-importation/domain/enums/userType"
)

type User struct {
	UserId               sql.NullString
	UserTypeId           sql.NullInt32
	Password             sql.NullString
	PasswordReset        sql.NullBool
	PasswordModifiedDate sql.NullTime
	CreatedDate          sql.NullTime
	EmployeeId           sql.NullInt64
	ModifiedDate         sql.NullTime
	ModifiedUser         sql.NullString
	AccessTypeId         sql.NullInt32
	Locked               sql.NullBool
	ExpirationDate       sql.NullTime
}

type Users []User

func (users Users) FindDatabaseUser() User {
	for i := range users {
		if users[i].UserTypeId.Int32 == userType.Database {
			return users[i]
		}
	}

	return User{}
}

func (users Users) FindActiveDirectoryUser() User {
	for i := range users {
		if users[i].UserTypeId.Int32 == userType.ActiveDirectory {
			return users[i]
		}
	}

	return User{}
}

func (user *User) SetUserId(userId string) {
	user.UserId = sql.NullString{String: userId, Valid: true}
}

func (user *User) SetUserTypeId(userTypeId int32) {
	user.UserTypeId = sql.NullInt32{Int32: userTypeId, Valid: true}
}

func (user *User) SetPassword(password string) {
	user.Password = sql.NullString{String: password, Valid: true}
}

func (user *User) SetPasswordReset(passwordReset bool) {
	user.PasswordReset = sql.NullBool{Bool: passwordReset, Valid: true}
}

func (user *User) SetPasswordModifiedDate(passwordModifiedDate time.Time) {
	user.PasswordModifiedDate = sql.NullTime{Time: passwordModifiedDate, Valid: true}
}

func (user *User) SetCreatedDate(createdDate time.Time) {
	user.CreatedDate = sql.NullTime{Time: createdDate, Valid: true}
}

func (user *User) SetEmployeeId(employeeId int64) {
	user.EmployeeId = sql.NullInt64{Int64: employeeId, Valid: true}
}

func (user *User) SetModifiedDate(modifiedDate time.Time) {
	user.ModifiedDate = sql.NullTime{Time: modifiedDate, Valid: true}
}

func (user *User) SetModifiedUser(modifiedUser string) {
	user.ModifiedUser = sql.NullString{String: modifiedUser, Valid: true}
}

func (user *User) SetAccessTypeId(accessTypeId int32) {
	user.AccessTypeId = sql.NullInt32{Int32: accessTypeId, Valid: true}
}

func (user *User) SetLocked(locked bool) {
	user.Locked = sql.NullBool{Bool: locked, Valid: true}
}

func (user *User) SetExpirationDate(expirationDate time.Time) {
	user.ExpirationDate = sql.NullTime{Time: expirationDate, Valid: true}
}

func (user *User) GetLocked() bool {
	return user.Locked.Bool
}

func (user *User) GetUserId() string {
	return user.UserId.String
}

func (user *User) IsValid() bool {
	return user.UserId.Valid && (user != &User{})
}
