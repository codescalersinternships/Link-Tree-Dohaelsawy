package repository

import (
	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"
)

func (suite *DatabaseTestSuite) TestLinkAddNewLink() {

	link := model.Link{Name: "youtube", Url: "youtube.com", UserID: 1}

	err := suite.dbInstance.AddNewLink(&link)
	suite.Require().NoError(err, "Error creating link record")

	var retrievedLink model.Link
	err = suite.dbInstance.GetOneLink(&retrievedLink, link.ID)
	suite.Require().NoError(err, "Error retrieving link record")
}

func (suite *DatabaseTestSuite) TestLinkDeleteLink() {

	link := model.Link{ID: 1, Name: "youtube", Url: "youtube.com", UserID: 1}

	err := suite.dbInstance.DeleteLink(&link, link.ID)
	suite.Require().NoError(err, "Error deleting link record")

	var retrievedLink model.Link
	err = suite.dbInstance.GetOneLink(&retrievedLink, link.ID)
	suite.Require().Error(err, "Error retrieving link record")
}

func (suite *DatabaseTestSuite) TestLinkPutOneLink() {

	link := model.Link{Name: "instagram", Url: "instagram.com", UserID: 1}

	err := suite.dbInstance.AddNewLink(&link)
	suite.Require().NoError(err, "Error add link record")

	link.Name = "x"

	err = suite.dbInstance.PutOneLink(&link, link.ID)
	suite.Require().NoError(err, "Error update link record")
	suite.Require().Equal("x", link.Name)
}

func (suite *DatabaseTestSuite) TestLinkGetAllLinksForUser() {
	var links []model.Link

	err := suite.dbInstance.GetAllLinksForUser(&links, 1)

	suite.Require().NoError(err, "Error retrieving links for specified user record")
	suite.Require().NotZero(links)
	suite.Require().NotEmpty(links)
}
