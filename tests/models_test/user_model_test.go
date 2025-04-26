package models_test

import (
	"testing"

	userModel "github.com/brianmorais/go-user-importation/domain/models/user"
	"github.com/brianmorais/go-user-importation/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestMustFindDatabaseUser(t *testing.T) {
	users := mocks.GetUsers()

	res := users.FindDatabaseUser()

	assert.NotEqual(t, res, userModel.User{})
}

func TestMustFindActiveDirectoryUser(t *testing.T) {
	users := mocks.GetUsers()

	res := users.FindActiveDirectoryUser()

	assert.NotEqual(t, res, userModel.User{})
}
