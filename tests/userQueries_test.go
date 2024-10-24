package tests

import (
	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"
)

func (suite *DatabaseTestSuite) TestLinkAddNewUser() {

	user := model.User{FirstName: "doha", LastName: "elsawy", Email: "x@gmail.com", Password: "password", Username: "do2333"}

	err := suite.DbInstance.AddNewUser(&user)
	suite.Require().NoError(err, "Error adding user before testing")

	var retrievedUser model.User
	err = suite.DbInstance.GetUserEmail(&retrievedUser, user.Email)
	suite.Require().NoError(err, "Error retrieving link record")
	suite.Require().Equal(user.Email, retrievedUser.Email)
	suite.Require().Equal(user.Password, retrievedUser.Password)
}

func (suite *DatabaseTestSuite) TestLinkDeleteUser() {

	user := model.User{FirstName: "doha", LastName: "elsawy", Email: "c@gmail.com", Password: "password", Username: "dosfs"}

	err := suite.DbInstance.DeleteUser(&user, user.ID)
	suite.Require().NoError(err, "Error deleting link record")

	var retrievedUser model.User
	err = suite.DbInstance.GetUserID(&retrievedUser, user.ID)
	suite.Require().Error(err, "Error retrieving link record")
}

func (suite *DatabaseTestSuite) TestLinkPutOneUser() {

	user := model.User{FirstName: "doha", LastName: "elsawy", Email: "c@gmail.com", Password: "password", Username: "dosfs"}

	err := suite.DbInstance.AddNewUser(&user)
	suite.Require().NoError(err, "Error adding user before testing")

	user.FirstName = "soha"

	err = suite.DbInstance.PutOneUser(&user, user.ID)
	suite.Require().NoError(err, "Error update link record")
	suite.Require().Equal("soha", user.FirstName)
}
