package env

import (
	"github.com/spf13/viper"
)

type Config struct {
	HttpPort       string
	PaymentService *PaymentServiceConfig
}

type PaymentServiceConfig struct {
	Host string
	Port string
}

var Conf *Config

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")

	viper.BindEnv("HTTP_PORT")
	viper.BindEnv("PAYMENT_SERVICE_HOST")
	viper.BindEnv("PAYMENT_SERVICE_PORT")

	Conf = &Config{
		HttpPort: viper.GetString("HTTP_PORT"),
		PaymentService: &PaymentServiceConfig{
			Host: viper.GetString("PAYMENT_SERVICE_HOST"),
			Port: viper.GetString("PAYMENT_SERVICE_PORT"),
		},
	}
}
