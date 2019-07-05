package users_test

import (
	"github.com/jonboydell/logzio_client/users"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	test_username = "test@massive.co"
	test_fullname = "Test User"
)

func TestUsers_CreateValidUser(t *testing.T) {
	underTest, err := setupUsersTest()

	if assert.NoError(t, err) {
		u := users.User{
			Username:  test_username,
			Fullname:  test_fullname,
			AccountId: underTest.AccountId,
			Roles:     []int32{users.UserTypeUser},
		}

		user, err := underTest.CreateUser(u)
		assert.NoError(t, err)
		if assert.NotNil(t, user) {
			v, err := underTest.GetUser(user.Id)

			if assert.NoError(t, err) && assert.NotNil(t, v) {
				assert.Equal(t, test_username, v.Username)
				assert.Equal(t, test_fullname, v.Fullname)
				assert.Equal(t, underTest.AccountId, v.AccountId)
				assert.True(t, v.Active)
				assert.Equal(t, user.Id, user.Id)
			}

			err = underTest.DeleteUser(user.Id)
			assert.NoError(t, err)
		}
	}
}

func TestUsers_CreateDeleteDuplicateUser(t *testing.T) {
	underTest, err := setupUsersTest()

	if assert.NoError(t, err) {
		u := users.User{
			Username:  test_username,
			Fullname:  test_fullname,
			AccountId: underTest.AccountId,
			Roles:     []int32{users.UserTypeUser},
		}

		user, err := underTest.CreateUser(u)
		assert.NoError(t, err)
		_, err = underTest.CreateUser(u)

		if assert.Error(t, err) {
			v, err := underTest.GetUser(user.Id)

			if assert.NoError(t, err) && assert.NotNil(t, v) {
				assert.Equal(t, test_username, v.Username)
				assert.Equal(t, test_fullname, v.Fullname)
				assert.Equal(t, underTest.AccountId, v.AccountId)
				assert.True(t, v.Active)
				assert.Equal(t, user.Id, user.Id)
			}

			err = underTest.DeleteUser(user.Id)
			assert.NoError(t, err)
		}
	}
}

func TestUsers_CreateInvalidUser_Email(t *testing.T) {
	underTest, err := setupUsersTest()

	if assert.NoError(t, err) {
		u := users.User{
			Username:  "SomeTestUser",
			Fullname:  "Test User",
			AccountId: underTest.AccountId,
			Roles:     []int32{users.UserTypeUser},
		}

		_, err := underTest.CreateUser(u)
		assert.Error(t, err)
	}
}