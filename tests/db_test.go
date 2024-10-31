package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/database/repository"
	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/utils"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseTestSuite struct {
	suite.Suite
	DbInstance repository.DbInstance
	config     model.Config
}

func TestSuite(t *testing.T) {

	setEnvVariables()

	suite.Run(t, new(DatabaseTestSuite))
}

func (suite *DatabaseTestSuite) SetupSuite() {

	suite.config = NewTestConfigController()
	conString := prepareDbTestingConnectionString(suite.config)
	fmt.Println(conString)

	db, err := gorm.Open(postgres.Open(conString), &gorm.Config{})
	suite.Require().NoError(err, "Error connecting to the test database")

	suite.DbInstance.DB = db.Debug()

	err = suite.DbInstance.DB.AutoMigrate(&model.User{}, &model.Link{}, &model.Analytics{})
	suite.Require().NoError(err, "Error auto-migrating database tables")

	Token_11, err := createTestToken(11, suite.config.JwtSecret)
	suite.Require().NoError(err, "Error token generating err")

	user := model.User{ID: 11, FirstName: "doha", LastName: "elsawy", Email: "aaad@gmail.com", Password: "$2a$14$SqGrotGlHpurAd6c.zfNt./oIW7Bh3fp1DAnh4nNTTEIMwfabqT8i", Username: "newusdfername", Token: Token_11}
	err = suite.DbInstance.AddNewUser(&user)
	suite.Require().NoError(err, "Error adding user before testing")

	Token_13, err := createTestToken(13, suite.config.JwtSecret)
	suite.Require().NoError(err, "Error token generating err")

	deleteUser := model.User{ID: 13, FirstName: "doha", LastName: "elsawy", Email: "delete@gmail.com", Password: "$2a$14$SqGrotGlHpurAd6c.zfNt./oIW7Bh3fp1DAnh4nNTTEIMwfabqT8i", Username: "delete", Token: Token_13}
	err = suite.DbInstance.AddNewUser(&deleteUser)
	suite.Require().NoError(err, "Error adding user before testing")

	link := model.Link{Name: "youtube", Url: "youtube.com", UserID: 11}
	err = suite.DbInstance.AddNewLink(&link)
	suite.Require().NoError(err, "Error creating link record")

	deleteLink := model.Link{Name: "youtube", Url: "youtube.com", UserID: 11}
	err = suite.DbInstance.AddNewLink(&deleteLink)
	suite.Require().NoError(err, "Error creating link record")

	editLink := model.Link{Name: "youtube", Url: "youtube.com", UserID: 11}
	err = suite.DbInstance.AddNewLink(&editLink)
	suite.Require().NoError(err, "Error creating link record")

	analytics := model.Analytics{ClickCount: 3, GuestUsername: "newusdfername", UserID: 11}
	err = suite.DbInstance.AddNewVisitor(&analytics)
	suite.Require().NoError(err, "Error creating link record")
}

func (suite *DatabaseTestSuite) TearDownSuite() {

	err := suite.DbInstance.DB.Exec("DROP TABLE links;").Error
	suite.Require().NoError(err, "Error dropping test table")

	err = suite.DbInstance.DB.Exec("DROP TABLE users;").Error
	suite.Require().NoError(err, "Error dropping test table")

}

// TestSuite runs the test suite.

func prepareDbTestingConnectionString(config model.Config) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DbHost, config.DbUser, config.DbPassword, config.DbName, config.DbPort)
}

func NewTestConfigController() model.Config {
	return model.Config{
		DbHost:            os.Getenv("DB_TEST_HOST"),
		DbUser:            os.Getenv("DB_TEST_USER"),
		DbPassword:        os.Getenv("DB_TEST_PASSWORD"),
		DbName:            os.Getenv("DB_TEST_NAME"),
		DbPort:            os.Getenv("DB_TEST_PORT"),
		Port:              os.Getenv("PORT"),
		JwtSecret:         os.Getenv("JWT_SECRET"),
		TokenHourLifeTime: os.Getenv("TOKEN_HOUR_LIFESPAN"),
		BaseUrl:           os.Getenv("BASE_URL"),
	}
}

func createTestToken(id uint, secretToken string) (string, error) {

	return utils.CreateToken(id, 24, secretToken)
}

func setEnvVariables() {
	os.Setenv("DB_TEST_HOST", "localhost")
	os.Setenv("DB_TEST_USER", "admin")
	os.Setenv("DB_TEST_PASSWORD", "adminpassword")
	os.Setenv("DB_TEST_NAME", "linktreedbtest")
	os.Setenv("DB_TEST_PORT", "4568")
	os.Setenv("PORT", "8010")
	os.Setenv("JWT_SECRET", "super_secure")
	os.Setenv("TOKEN_HOUR_LIFESPAN", "24")
	os.Setenv("BASE_URL", "http://localhost:8010")
}
