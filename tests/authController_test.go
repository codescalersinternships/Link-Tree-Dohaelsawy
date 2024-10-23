package tests

// import (
// 	"log"
// 	"net/http"
// 	"net/http/httptest"

// 	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/controllers"
// 	"github.com/gin-gonic/gin"
// )

// func (suite *DatabaseTestSuite) TestRegister() {

// 	store := controllers.NewDBService(&suite.DbInstance)

// 	router := gin.Default()
// 	router.GET("/auth/register", store.Register)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/auth/register",func(c *gin.Context){
// 			http.StatusOK
// 			c.JSON(200, gin.H{
// 				"first_name": "doha",
// 				"last_name":  "elsawy",
// 				"username":   "dwasdewj9",
// 				"email":      "a@gmail.com",
// 				"password":   "12341234",
// 			}),
// 		})
// 	router.ServeHTTP(w, req)

// 	suite.Require().Equal(w.Result().StatusCode, http.StatusOK)

// }
