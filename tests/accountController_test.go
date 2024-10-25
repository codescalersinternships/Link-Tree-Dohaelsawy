package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"

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
				Bio:       "it's me",
			},
			status: http.StatusBadRequest,
		},
	}

	for _, test := range testcase {

		router := SetupAccountRouter(suite)
		dbService := controllers.NewDBService(&suite.DbInstance, suite.config)

		router.PUT("/edit_account", dbService.EditAccount)

		jsonValue, err := json.Marshal(test.user)
		suite.Require().NoError(err, "Error can't marshal user to json")

		req, err := http.NewRequest("PUT", "/edit_account", bytes.NewBuffer(jsonValue))
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

	router.DELETE("/delete_account", dbService.DeleteAccount)

	req, err := http.NewRequest("DELETE", "/delete_account", nil)
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

func (suite *DatabaseTestSuite) TestGetAccount() {
	router := SetupAccountRouter(suite)
	dbService := controllers.NewDBService(&suite.DbInstance, suite.config)

	router.GET("/get_account/", dbService.GetAccount)

	req, err := http.NewRequest("GET", "/get_account/", nil)
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

	router.GET("/create_link_tree_url", dbService.CreateLinkTreeUrl)

	req, err := http.NewRequest("GET", "/create_link_tree_url", nil)
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

func (suite *DatabaseTestSuite) TestUploadUserPhoto() {
	router := SetupAccountRouter(suite)
	dbService := controllers.NewDBService(&suite.DbInstance, suite.config)
	router.POST("/add_photo", dbService.UploadUserImage)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	file, err := os.CreateTemp("testdata", "image.jpeg")
	suite.Require().NoError(err)
	defer os.Remove(file.Name())

	_, err = file.Write([]byte("This is a test image"))
	suite.Require().NoError(err)
	file.Seek(0, io.SeekStart)

	part, err := writer.CreateFormFile("image", filepath.Base(file.Name()))
	suite.Require().NoError(err)

	_, err = io.Copy(part, file)
	suite.Require().NoError(err)

	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/add_photo", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

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
