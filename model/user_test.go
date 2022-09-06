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
	suite.Equal(2, len(users))

}

func (suite UserTestSuite) TestShouldReturnSpecificUserByUserId() {
	user := GetUserById(2)
	suite.Equal(2, user.ID)

}
