package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost              string        `mapstructure:"DB_HOST"`
	DBPort              int           `mapstructure:"DB_PORT"`
	DBUser              string        `mapstructure:"DB_USER"`
	DBPassword          string        `mapstructure:"DB_PASSWORD"`
	DBName              string        `mapstructure:"DB_NAME"`
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	FrontendDomain      string        `mapstructure:"FRONTEND_DOMAIN"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	S3Bucket            string        `mapstructure:"S3_BUCKET"`
	AWSRegion           string        `mapstructure:"AWS_REGION"`
	AWSAccessKeyID      string        `mapstructure:"AWS_ACCESS_KEY_ID"`
	AWSSecretAccessKey  string        `mapstructure:"AWS_SECRET_ACCESS_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
