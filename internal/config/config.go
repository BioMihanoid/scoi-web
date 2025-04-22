package config

import "github.com/spf13/viper"

type Config struct {
	Server struct {
		Port string
	}
	Database struct {
		DSN string
	}
	Auth struct {
		JWTSecret string
	}
	S3 struct {
		Region    string
		Bucket    string
		AccessKey string
		SecretKey string
		Endpoint  string
	}
}

func MustLoad() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}
