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
			status: http.StatusOK,
		},
	}

	for _, test := range testcase {

		router := SetupAccountRouter(suite)
		dbService := controllers.NewDBService(&suite.DbInstance, suite.config)

		Token_11, err := createTestToken(11, suite.config.JwtSecret)
		suite.Require().NoError(err, "Error token generating err")

		router.PUT("/edit_account", dbService.EditAccount)

		jsonValue, err := json.Marshal(test.user)
		suite.Require().NoError(err, "Error can't marshal user to json")

		req, err := http.NewRequest("PUT", "/edit_account", bytes.NewBuffer(jsonValue))
		suite.Require().NoError(err, "Error create http request")
		req.AddCookie(&http.Cookie{
			Name:     "Authorization",
			Value:    Token_11,
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

	Token_13, err := createTestToken(13, suite.config.JwtSecret)
	suite.Require().NoError(err, "Error token generating err")

	router.DELETE("/delete_account", dbService.DeleteAccount)

	req, err := http.NewRequest("DELETE", "/delete_account", nil)
	suite.Require().NoError(err, "Error create http request")
	req.AddCookie(&http.Cookie{
		Name:     "Authorization",
		Value:    Token_13,
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

	Token_11, err := createTestToken(11, suite.config.JwtSecret)
	suite.Require().NoError(err, "Error token generating err")

	router.GET("/get_account/", dbService.GetAccount)

	req, err := http.NewRequest("GET", "/get_account/", nil)
	suite.Require().NoError(err, "Error create http request")
	req.AddCookie(&http.Cookie{
		Name:     "Authorization",
		Value:    Token_11,
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

	Token_11, err := createTestToken(11, suite.config.JwtSecret)
	suite.Require().NoError(err, "Error token generating err")

	router.POST("/add_photo", dbService.UploadUserImage)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	file, err := os.CreateTemp("testdata", "image.png")
	suite.Require().NoError(err)
	defer os.Remove(file.Name())

	_, err = file.Write([]byte(""))
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
		Value:    Token_11,
		Path:     "",
		Domain:   "",
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	suite.Require().Equal(http.StatusOK, w.Body.String())

}

func SetupAccountRouter(suite *DatabaseTestSuite) *gin.Engine {

	router := gin.Default()
	router.Use(middleware.AuthMiddleware(suite.config))
	route.AccountRouters(suite.DbInstance, suite.config, router)

	return router
}
