package repositories

import (
	userModel "github.com/brianmorais/go-user-importation/domain/models/user"
)

type UserRepository struct{}

func (UserRepository) FindUserById(userId string) (userModel.User, error) {
	conn := getWriteConnection()

	defer conn.Close()

	queryText := `
		SELECT 
			UserId,
			UserTypeId,
			[Password],
			PasswordReset,
			PasswordModifiedDate,
			ExpirationDate,
			Locked,
			EmployeeId,
			AccessTypeId,
			CreatedDate,
			ModifiedDate,
			ModifiedUser
		FROM
			[User]
		WHERE 
			UserId = ?`

	rows, err := conn.Query(queryText, userId)

	if err != nil {
		return userModel.User{}, err
	}

	defer rows.Close()

	var user userModel.User

	for rows.Next() {
		err = rows.Scan(
			&user.UserId,
			&user.UserTypeId,
			&user.Password,
			&user.PasswordReset,
			&user.PasswordModifiedDate,
			&user.ExpirationDate,
			&user.Locked,
			&user.EmployeeId,
			&user.AccessTypeId,
			&user.CreatedDate,
			&user.ModifiedDate,
			&user.ModifiedUser,
		)
	}

	if err != nil {
		return userModel.User{}, err
	}

	return user, nil
}

func (UserRepository) CreateUser(user userModel.User) (int64, error) {
	conn := getWriteConnection()

	defer conn.Close()

	queryText := `
		INSERT INTO [User] (
			UserId,
			UserTypeId,
			[Password],
			PasswordReset,
			PasswordModifiedDate,
			ExpirationDate,
			Locked,
			EmployeeId,
			AccessTypeId,
			CreatedDate,
			ModifiedDate,
			ModifiedUser
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	res, err := conn.Exec(
		queryText,
		user.UserId,
		user.UserTypeId,
		user.Password,
		user.PasswordReset,
		user.PasswordModifiedDate,
		user.ExpirationDate,
		user.Locked,
		user.EmployeeId,
		user.AccessTypeId,
		user.CreatedDate,
		user.ModifiedDate,
		user.ModifiedUser,
	)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (UserRepository) UpdateUser(user userModel.User) (int64, error) {
	conn := getWriteConnection()

	defer conn.Close()

	queryText := `
		UPDATE [User] SET 
			UserTypeId = ?,
			[Password] = ?,
			PasswordReset = ?,
			PasswordModifiedDate = ?,
			ExpirationDate = ?,
			Locked = ?,
			EmployeeId = ?,
			AccessTypeId = ?,
			ModifiedDate = ?,
			ModifiedUser = ?
		WHERE 
			UserId = ?`

	res, err := conn.Exec(
		queryText,
		user.UserTypeId,
		user.Password,
		user.PasswordReset,
		user.PasswordModifiedDate,
		user.ExpirationDate,
		user.Locked,
		user.EmployeeId,
		user.AccessTypeId,
		user.ModifiedDate,
		user.ModifiedUser,
		user.UserId,
	)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (UserRepository) FindUsersByEmployeeId(employeeId int64) (userModel.Users, error) {
	conn := getWriteConnection()

	defer conn.Close()

	queryText := `
		SELECT 
			UserId,
			UserTypeId,
			[Password],
			PasswordReset,
			PasswordModifiedDate,
			ExpirationDate,
			Locked,
			EmployeeId,
			AccessTypeId,
			CreatedDate,
			ModifiedDate,
			ModifiedUser
		FROM
			[User]
		WHERE 
			EmployeeId = ?`

	rows, err := conn.Query(queryText, employeeId)

	if err != nil {
		return userModel.Users{}, err
	}

	defer rows.Close()

	var users userModel.Users

	for rows.Next() {
		var user userModel.User

		err = rows.Scan(
			&user.UserId,
			&user.UserTypeId,
			&user.Password,
			&user.PasswordReset,
			&user.PasswordModifiedDate,
			&user.ExpirationDate,
			&user.Locked,
			&user.EmployeeId,
			&user.AccessTypeId,
			&user.CreatedDate,
			&user.ModifiedDate,
			&user.ModifiedUser,
		)

		if err != nil {
			return userModel.Users{}, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (UserRepository) DeleteUser(userId string) (int64, error) {
	conn := getWriteConnection()

	defer conn.Close()

	queryText := "DELETE FROM [User] WHERE UserId = ?"

	res, err := conn.Exec(
		queryText,
		userId,
	)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (UserRepository) GetUsersByTypeAndEmployeeId(employeeId int64, userTypeId int32) (userModel.Users, error) {
	conn := getWriteConnection()

	defer conn.Close()

	queryText := `
		SELECT 
			UserId,
			UserTypeId,
			[Password],
			PasswordReset,
			PasswordModifiedDate,
			ExpirationDate,
			Locked,
			EmployeeId,
			AccessTypeId,
			CreatedDate,
			ModifiedDate,
			ModifiedUser
		FROM
			[User]
		WHERE 
			Locked = 0
			AND EmployeeId = ?
			AND UserTypeId = ?`

	rows, err := conn.Query(queryText, employeeId, userTypeId)

	if err != nil {
		return userModel.Users{}, err
	}

	defer rows.Close()

	var users userModel.Users

	for rows.Next() {
		var user userModel.User

		err = rows.Scan(
			&user.UserId,
			&user.UserTypeId,
			&user.Password,
			&user.PasswordReset,
			&user.PasswordModifiedDate,
			&user.ExpirationDate,
			&user.Locked,
			&user.EmployeeId,
			&user.AccessTypeId,
			&user.CreatedDate,
			&user.ModifiedDate,
			&user.ModifiedUser,
		)

		if err != nil {
			return userModel.Users{}, err
		}

		users = append(users, user)
	}

	return users, nil
}
