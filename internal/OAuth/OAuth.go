package oauth

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type OAuthService interface {
}

type OAuthServiceImpl struct {
	httpClient         *http.Client
	refreshEndpointURL string
	clientID           string
	clientSecret       string
	redirectURI        string
}

func NewOAuthService(refreshEndpointURL, clientID, clientSecret, redirectURI string) OAuthService {
	oas := &OAuthServiceImpl{
		httpClient:         &http.Client{},
		refreshEndpointURL: refreshEndpointURL,
		clientID:           clientID,
		clientSecret:       clientSecret,
		redirectURI:        redirectURI,
	}
	return oas
}

func (oas *OAuthServiceImpl) GetOAuthToken(authorizationCode string) (OAuthToken, error) {
	/*
		После получения authorization_code приложению необходимо сделать сервер-сервер запрос
		POST https://hh.ru/oauth/token для обмена полученного authorization_code на access_token.

		В теле запроса необходимо передать дополнительные параметры:

		grant_type=authorization_code
		client_id и client_secret - необходимо заполнить значениями, выданными при регистрации приложения
		redirect_uri - если параметр был уточнен на шаге получения авторизации, необходимо передать уточненный redirect_uri или вернется ошибка,
		если уточнения не было, параметр можно не посылать.
		code – значение authorization_code, полученное при перенаправлении пользователя

		Тело запроса необходимо передавать в стандартном application/x-www-form-urlencoded с указанием соответствующего заголовка Content-Type.
	*/
	body := url.Values{}
	body.Add("grant_type", "authorization_code")
	body.Add("client_id", oas.clientID)
	body.Add("client_secret", oas.clientSecret)
	body.Add("redirect_uri", oas.redirectURI)
	body.Add("code", authorizationCode)

	req, err := http.NewRequest(http.MethodPost, oas.refreshEndpointURL, strings.NewReader(body.Encode()))
	if err != nil {
		return OAuthToken{}, fmt.Errorf("(*OAuthService).GetOAuthToken - http.NewRequest: %w", err)
	}

}
