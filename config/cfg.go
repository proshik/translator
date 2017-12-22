package config

import "github.com/kelseyhightower/envconfig"

const (
	// SERVICENAME contains a service name prefix which used in ENV variables
	SERVICENAME = "TRANSLATOR"
)


type Config struct {
	LogDir                string `default:"/"`
	YandexDictionaryToken string `required:"true"`
	YandexTranslatorToken string `required:"false"`
}

func Load() (*Config, error) {

	cfg := new(Config)

	err := envconfig.Process(SERVICENAME, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
