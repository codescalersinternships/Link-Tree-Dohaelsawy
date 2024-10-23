package repository

import (
	"fmt"
	"os"
	"testing"

	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseTestSuite struct {
	suite.Suite
	dbInstance DbInstance
}

func TestSuite(t *testing.T) {
	setEnvVariables()
	suite.Run(t, new(DatabaseTestSuite))
}

func (suite *DatabaseTestSuite) SetupSuite() {

	config := NewTestConfigController()
	conString := prepareDbTestingConnectionString(config)
	fmt.Println(conString)

	db, err := gorm.Open(postgres.Open(conString), &gorm.Config{})
	suite.Require().NoError(err, "Error connecting to the test database")

	suite.dbInstance.DB = db.Debug()

	err = suite.dbInstance.DB.AutoMigrate(&model.User{}, &model.Link{})
	suite.Require().NoError(err, "Error auto-migrating database tables")

	user := model.User{FirstName: "doha", LastName: "elsawy", Email: "doha@gmail.com", Password: "password",Username: "do23"}
	err = suite.dbInstance.AddNewUser(&user)
	suite.Require().NoError(err, "Error adding user before testing")

	link := model.Link{Name: "youtube", Url: "youtube.com", UserID: 1}
	err = suite.dbInstance.AddNewLink(&link)
	suite.Require().NoError(err, "Error creating link record")
}

func (suite *DatabaseTestSuite) TearDownSuite() {

	err := suite.dbInstance.DB.Exec("DROP TABLE links;").Error
	suite.Require().NoError(err, "Error dropping test table")

	err = suite.dbInstance.DB.Exec("DROP TABLE users;").Error
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
	}
}

func setEnvVariables() {
	os.Setenv("DB_TEST_HOST", "localhost")
	os.Setenv("DB_TEST_USER", "admin")
	os.Setenv("DB_TEST_PASSWORD", "adminpassword")
	os.Setenv("DB_TEST_NAME", "linktreedbtest")
	os.Setenv("DB_TEST_PORT", "4568")
}
