package tests

import (
	"net/http"
	"net/http/httptest"

	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/controllers"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/middleware"
	route "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/routers"
	"github.com/gin-gonic/gin"
)

func (suite *DatabaseTestSuite) TestGetAnalytics() {
	router := SetupAnalyticsRouter(suite)
	dbService := controllers.NewController(&suite.DbInstance, suite.config)

	Token_11, err := createTestToken(11, suite.config.JwtSecret)
	suite.Require().NoError(err, "Error token generating err")

	router.GET("/get_analytics/:user_id", dbService.GetAnalytics)

	req, err := http.NewRequest("GET", "/get_analytics/11", nil)
	suite.Require().NoError(err, "Error create http request")
	req.AddCookie(&http.Cookie{
		Name:     "Authorization",
		Value:   Token_11,
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

func (suite *DatabaseTestSuite) TestCalculateAnalytics() {
	testcase := []struct {
		guestUsername string
		userID        int
		err           error
	}{
		{
			guestUsername: "unknown",
			userID:        11,
			err:           nil,
		},
		{
			guestUsername: "newusdfername",
			userID:        11,
			err:           nil,
		},
	}

	for _, test := range testcase {
		dbService := controllers.NewController(&suite.DbInstance, suite.config)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// Call the function
		dbService.CalculateAnalytics(c, test.guestUsername, test.userID)

		suite.Require().Empty(w.Body.String())
	}

}

func SetupAnalyticsRouter(suite *DatabaseTestSuite) *gin.Engine {

	router := gin.Default()
	router.Use(middleware.AuthMiddleware(suite.config))
	route.AnalyticsRouters(suite.DbInstance, suite.config, router)

	return router
}
