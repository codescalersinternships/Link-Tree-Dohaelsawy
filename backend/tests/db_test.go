package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/database/repository"
	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/models"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/utils"
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
	suite.Run(t, new(DatabaseTestSuite))
}

func (suite *DatabaseTestSuite) SetupSuite() {

	config, err := NewTestConfigController()
	suite.Require().NoError(err, "Error loading envs")
	suite.config = config
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
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DbHost ,config.DbUser, config.DbPassword, config.DbName, config.DbPort)
}

func NewTestConfigController() (model.Config , error) {
	fmt.Println(os.Environ())
	return model.Config{
		DbHost:             getEnv("DB_TEST_HOST","localhost"),
		DbUser:             getEnv("DB_TEST_USER","admin"),
		DbPassword:         getEnv("DB_TEST_PASSWORD", "adminpassword"),
		DbName:             getEnv("DB_TEST_NAME","linktreedbtest"),
		DbPort:             getEnv("DB_TEST_PORT","4568"),
		Port:               getEnv("PORT","8010"),
		JwtSecret:          os.Getenv("JWT_SECRET"),
		TokenHourLifeTime:  getEnv("TOKEN_HOUR_LIFESPAN","24"),
		BaseUrl:            getEnv("BASE_URL","http://185.206.122.17:31111"),
		AwsRegion:          os.Getenv("AWS_REGION"),
		AwsAccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		AwsSecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}, nil
}

func createTestToken(id uint, secretToken string) (string, error) {
	return utils.CreateToken(id, 24, secretToken)
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}