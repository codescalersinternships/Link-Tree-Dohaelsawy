package model

type Config struct {
	DbHost            string
	DbUser            string
	DbPassword        string
	DbName            string
	DbPort            string
	Port              string
	JwtSecret         string
	TokenHourLifeTime string
	BaseUrl           string
}
