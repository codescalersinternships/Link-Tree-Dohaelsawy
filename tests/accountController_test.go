package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/controllers"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/middleware"
	route "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/routers"

	"github.com/gin-gonic/gin"
)

func (suite *DatabaseTestSuite) TestEditAccount() {
	testcase := []struct {
		title  string
		user   controllers.AccountReq
		status int
	}{
		{
			title: "valid user input",
			user: controllers.AccountReq{
				FirstName: "doha",
				LastName:  "elsawy",
				Phone:     "12345678901",
				Photo:     "",
				Bio:       "it's me",
			},
			status: http.StatusOK,
		},
		{
			title: "violates user phone validation less than 11 numbers",
			user: controllers.AccountReq{
				FirstName: "doha",
				LastName:  "elsawy",
				Phone:     "12345",
				Photo:     "",
				Bio:       "it's me",
			},
			status: http.StatusBadRequest,
		},
	}

	for _, test := range testcase {

		router := SetupAccountRouter(suite)
		dbService := controllers.NewDBService(&suite.DbInstance, suite.config)

		router.POST("/edit_account/:user_id", dbService.EditAccount)

		jsonValue, err := json.Marshal(test.user)
		suite.Require().NoError(err, "Error can't marshal user to json")

		req, err := http.NewRequest("POST", "/edit_account/11", bytes.NewBuffer(jsonValue))
		suite.Require().NoError(err, "Error create http request")
		req.AddCookie(&http.Cookie{
			Name:     "Authorization",
			Value:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjk4NTU0ODIsImlhdCI6MTcyOTc2OTA4Miwic3VwIjoxMX0.gtfXET5b2AFUqAja2Dv8T2VM3tR7YNtq6EPIlsmvV3Q",
			Path:     "",
			Domain:   "",
			Secure:   false,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		})

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		suite.Require().Equal(test.status, w.Code)
		suite.Require().NotEmpty(w.Body.String())
	}
}

func (suite *DatabaseTestSuite) TestDeleteAccount() {
	router := SetupAccountRouter(suite)
	dbService := controllers.NewDBService(&suite.DbInstance, suite.config)

	router.POST("/delete_account/:user_id", dbService.DeleteAccount)

	req, err := http.NewRequest("POST", "/delete_account/1", nil)
	suite.Require().NoError(err, "Error create http request")
	req.AddCookie(&http.Cookie{
		Name:     "Authorization",
		Value:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjk4NTU0ODIsImlhdCI6MTcyOTc2OTA4Miwic3VwIjoxMX0.gtfXET5b2AFUqAja2Dv8T2VM3tR7YNtq6EPIlsmvV3Q",
		Path:     "",
		Domain:   "",
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	suite.Require().Equal(http.StatusOK, w.Code)
}

func (suite *DatabaseTestSuite) TestGetAccount() {
	router := SetupAccountRouter(suite)
	dbService := controllers.NewDBService(&suite.DbInstance, suite.config)

	router.POST("/get_account/", dbService.GetAccount)

	req, err := http.NewRequest("POST", "/get_account/", nil)
	suite.Require().NoError(err, "Error create http request")
	req.AddCookie(&http.Cookie{
		Name:     "Authorization",
		Value:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjk4NTU0ODIsImlhdCI6MTcyOTc2OTA4Miwic3VwIjoxMX0.gtfXET5b2AFUqAja2Dv8T2VM3tR7YNtq6EPIlsmvV3Q",
		Path:     "",
		Domain:   "",
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	suite.Require().Equal(http.StatusOK, w.Code)
}

func (suite *DatabaseTestSuite) TestCreateLinkTreeUrl() {

	router := SetupAccountRouter(suite)
	dbService := controllers.NewDBService(&suite.DbInstance, suite.config)

	router.POST("/create_link_tree_url", dbService.CreateLinkTreeUrl)

	req, err := http.NewRequest("POST", "/create_link_tree_url", nil)
	suite.Require().NoError(err, "Error create http request")
	req.AddCookie(&http.Cookie{
		Name:     "Authorization",
		Value:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjk4NTU0ODIsImlhdCI6MTcyOTc2OTA4Miwic3VwIjoxMX0.gtfXET5b2AFUqAja2Dv8T2VM3tR7YNtq6EPIlsmvV3Q",
		Path:     "",
		Domain:   "",
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	suite.Require().Equal(http.StatusOK, w.Code)
}

func SetupAccountRouter(suite *DatabaseTestSuite) *gin.Engine {

	router := gin.Default()
	router.Use(middleware.AuthMiddleware(suite.config))
	route.AccountRouters(suite.DbInstance, suite.config, router)

	return router
}
