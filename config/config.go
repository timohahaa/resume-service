package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	OAuthHH struct {
		BaseApiUrl string `yaml:"baseApiUrl"`
		// OAuth endpoint for hh.ru
		OAuthEndpoint string `yaml:"OAuthEndpointURL"`
		// RedirectURI - where to redirect the user after the authorization is finished
		RedirectURI string `yaml:"redirectURI"`
		// ID of the hh.ru api client(application) and a secret
		CliendID     string `yaml:"clientID"`
		ClientSecret string `yaml:"clientSecret"`
		// Endpoint for refreshing the access token
		RefreshEndpointURL string `yaml:"refreshEndpointURL"`
		// user agent header, that needs to be sent witt every api call
		HHUserAgent string `yaml:"HH-User-Agent"`
	}
	Config struct {
		OAuthHH `yaml:"OAuthHH"`
	}
)

func NewConfig(filePath string) (*Config, error) {
	var conf = &Config{}
	err := cleanenv.ReadConfig(filePath, conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
