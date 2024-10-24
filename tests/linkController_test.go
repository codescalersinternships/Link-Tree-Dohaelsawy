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

func (suite *DatabaseTestSuite) TestCreateLink() {
	testcase := []struct {
		title  string
		link   controllers.LinkReq
		status int
	}{
		{
			title: "valid link input",
			link: controllers.LinkReq{
				Name: "website",
				Url:  "website.com",
			},
			status: http.StatusOK,
		},
	}

	for _, test := range testcase {

		router := SetupAccountRouter(suite)
		dbService := controllers.NewDBService(&suite.DbInstance, suite.config)

		router.POST("/create_link", dbService.CreateLink)

		jsonValue, err := json.Marshal(test.link)
		suite.Require().NoError(err, "Error can't marshal user to json")

		req, err := http.NewRequest("POST", "/create_link", bytes.NewBuffer(jsonValue))
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

func (suite *DatabaseTestSuite) TestDeleteLink() {
	router := SetupAccountRouter(suite)
	dbService := controllers.NewDBService(&suite.DbInstance, suite.config)

	router.DELETE("/delete_link/:link_id", dbService.DeleteLink)

	req, err := http.NewRequest("DELETE", "/delete_link/2", nil)
	suite.Require().NoError(err, "Error create http request")
	req.AddCookie(&http.Cookie{
		Name:     "Authorization",
		Value:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjk4NjAwMTcsImlhdCI6MTcyOTc3MzYxNywic3VwIjoxM30.M06Hr3QWA-fKkrbpsiQA9VAwmg4JlNXY_knn6iB6kZE",
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

func (suite *DatabaseTestSuite) TestUpdateLink() {
	testcase := []struct {
		title  string
		link   controllers.LinkReq
		status int
	}{
		{
			title: "valid link input",
			link: controllers.LinkReq{
				Url:  "jdhfajdsh",
				Name: "kjhkfhk",
			},
			status: http.StatusOK,
		},
	}

	for _, test := range testcase {
		router := SetupAccountRouter(suite)
		dbService := controllers.NewDBService(&suite.DbInstance, suite.config)

		router.PUT("/update_link/:link_id", dbService.UpdateLink)

		jsonValue, err := json.Marshal(test.link)
		suite.Require().NoError(err, "Error can't marshal user to json")

		req, err := http.NewRequest("PUT", "/update_link/3", bytes.NewBuffer(jsonValue))
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

}

func (suite *DatabaseTestSuite) TestGetLinks() {

	router := SetupAccountRouter(suite)
	dbService := controllers.NewDBService(&suite.DbInstance, suite.config)

	router.GET("/get_links", dbService.CreateLinkTreeUrl)

	req, err := http.NewRequest("GET", "/get_links", nil)
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

func SetupLinkRouter(suite *DatabaseTestSuite) *gin.Engine {

	router := gin.Default()
	router.Use(middleware.AuthMiddleware(suite.config))
	route.LinkRouters(suite.DbInstance, suite.config, router)

	return router
}
