package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port            string `mapstructure:"PORT"`
	AuthSvcUrl      string `mapstructure:"AUTH_SVC_URL"`
	ProductSvcUrl   string `mapstructure:"PRODUCT_SVC_URL"`
	OrderSvcUrl     string `mapstructure:"ORDER_SVC_URL"`
	AdminAuthSvcUrl string `mapstructure:"ADMIN_AUTH_SVC_URL"`
	CartSvcUrl      string `mapstructure:"CART_SVC_URL"`
}

var envs = []string{
	"PORT", "AUTH_SVC", "PRODUCT_SVC_URL", "ORDER_SVC_URL", "ADMIN_AUTH_SVC_URL", "CART_SVC_URL",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}


	return config, nil

}
