package model

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	suite.Suite
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
	UserGeneratorTearDown()
}

func UserGeneratorTearDown() {
	os.RemoveAll("generated")
}

func (suite UserTestSuite) TestShouldGetAllTheUsers() {
	users, err := GetAllUsers()
	if err != nil {
		suite.Fail("Error While Get All users")
	}
	suite.True(len(users) > 0)

}

func (suite UserTestSuite) TestShouldReturnSpecificUserByUserId() {
	user := GetUserById(1)
	suite.Equal(1, user.ID)
}

func (suite UserTestSuite) TestShouldCreateUser() {
	mockUser := User{
		ID:      3,
		Name:    "John Smith",
		Phone:   "9847578494",
		Address: "Sydney",
	}
	CreateUser(mockUser)
	users, _ := GetAllUsers()
	suite.Equal(3, len(users))
}
