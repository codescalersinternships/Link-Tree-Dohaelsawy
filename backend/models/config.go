package model

type Config struct {
	DbHost             string
	DbUser             string
	DbPassword         string
	DbName             string
	DbPort             string
	Port               string
	JwtSecret          string
	TokenHourLifeTime  string
	BaseUrl            string
	UserImagePath      string
	LinkTreePath       string
	StaticImagesPath   string
	AwsRegion          string
	AwsAccessKeyID     string
	AwsSecretAccessKey string
	Origin             string
	DB_CACHE_ADDR      string
	DB_CACHE_PASSWORD  string
}
