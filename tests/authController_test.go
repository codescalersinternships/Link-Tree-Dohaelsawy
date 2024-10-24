package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/controllers"
	route "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/routers"

	"github.com/gin-gonic/gin"
)

func (suite *DatabaseTestSuite) TestRegister() {
	testcase := []struct {
		title  string
		user   controllers.RegisterRequest
		status int
	}{
		{
			title: "valid user input",
			user: controllers.RegisterRequest{
				FirstName: "doha",
				LastName:  "elsawy",
				Email:     "newemail@gmail.com",
				Username:  "doha55555",
				Password:  "12345678",
			},
			status: http.StatusOK,
		},
		{
			title: "violates user email unique validation input",
			user: controllers.RegisterRequest{
				FirstName: "doha",
				LastName:  "elsawy",
				Email:     "newemail@gmail.com",
				Username:  "doha55555",
				Password:  "12345678",
			},
			status: http.StatusInternalServerError,
		},
		{
			title: "violates user password less than 8",
			user: controllers.RegisterRequest{
				FirstName: "doha",
				LastName:  "elsawy",
				Email:     "verynewemail@gmail.com",
				Username:  "newusername55555",
				Password:  "5678",
			},
			status: http.StatusBadRequest,
		},
	}

	for _, test := range testcase {

		router := SetupAuthRouter(suite)
		dbService := controllers.NewDBService(&suite.DbInstance, suite.config)
		router.POST("/register", dbService.Register)

		jsonValue, err := json.Marshal(test.user)
		suite.Require().NoError(err, "Error can't marshal user to json")

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
		suite.Require().NoError(err, "Error create http request")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		suite.Require().Equal(test.status, w.Code)
	}
}

func (suite *DatabaseTestSuite) TestLogin() {
	testcase := []struct {
		title  string
		user   controllers.LoginRequest
		status int
	}{
		{
			title: "valid user input",
			user: controllers.LoginRequest{
				Email:    "aaad@gmail.com",
				Password: "12341234",
			},
			status: http.StatusOK,
		},
		{
			title: "violates user email not exist",
			user: controllers.LoginRequest{
				Email:    "notexistemail@gmail.com",
				Password: "12345678",
			},
			status: http.StatusNotFound,
		},
	}

	for _, test := range testcase {

		router := SetupAuthRouter(suite)

		dbService := controllers.NewDBService(&suite.DbInstance, suite.config)
		router.POST("/login", dbService.Login)

		jsonValue, err := json.Marshal(test.user)
		suite.Require().NoError(err, "Error can't marshal user to json")

		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
		suite.Require().NoError(err, "Error create http request")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		suite.Require().Equal(test.status, w.Code)
	}
}

func (suite *DatabaseTestSuite) TestLogout() {
	router := SetupAuthRouter(suite)

	dbService := controllers.NewDBService(&suite.DbInstance, suite.config)
	router.POST("/logout", dbService.Logout)

	req, err := http.NewRequest("POST", "/logout", nil)
	suite.Require().NoError(err, "Error create http request")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Empty(w.Result().Cookies()[0].Value)

}

func SetupAuthRouter(suite *DatabaseTestSuite) *gin.Engine {

	router := gin.Default()
	route.AuthRouters(suite.DbInstance, suite.config, router)

	return router
}
