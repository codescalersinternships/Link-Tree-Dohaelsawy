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
	dbService := controllers.NewDBService(&suite.DbInstance, suite.config)

	router.GET("/get_analytics/:user_id", dbService.GetAnalytics)

	req, err := http.NewRequest("GET", "/get_analytics/11", nil)
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
		dbService := controllers.NewDBService(&suite.DbInstance, suite.config)

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
