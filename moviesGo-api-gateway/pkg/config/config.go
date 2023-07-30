package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	Port            string `mapstructure:"PORT"`
	AuthSvcUrl      string `mapstructure:"AUTH_SVC_URL"`
	ProductSvcUrl   string `mapstructure:"PRODUCT_SVC_URL"`
	OrderSvcUrl     string `mapstructure:"ORDER_SVC_URL"`
	AdminAuthSvcUrl string `mapstructure:"ADMIN_AUTH_SVC_URL"`
	CartSvcUrl      string `mapstructure:"CART_SCV_URL"`
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil

}
