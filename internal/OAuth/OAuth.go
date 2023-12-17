package oauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type OAuthService interface {
}

type OAuthServiceImpl struct {
	httpClient         *http.Client
	refreshEndpointURL string
	clientID           string
	clientSecret       string
	redirectURI        string
	baseApiURL         string
}

func NewOAuthService(refreshEndpointURL, clientID, clientSecret, redirectURI, baseApiURL string) *OAuthServiceImpl {
	oas := &OAuthServiceImpl{
		httpClient:         &http.Client{},
		refreshEndpointURL: refreshEndpointURL,
		clientID:           clientID,
		clientSecret:       clientSecret,
		redirectURI:        redirectURI,
		baseApiURL:         baseApiURL,
	}
	return oas
}

type getOAuthTokenResponce struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

func (oas *OAuthServiceImpl) GetToken(authorizationCode string) (OAuthToken, error) {
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
		return OAuthToken{}, fmt.Errorf("(*OAuthService).GetToken - http.NewRequest: %w", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := oas.httpClient.Do(req)
	if err != nil {
		return OAuthToken{}, fmt.Errorf("(*OAuthService).GetToken - http.Client.Do: %w", err)

	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return OAuthToken{}, fmt.Errorf("(*OAuthService).GetToken - io.ReadAll: %w", err)

	}

	var tokenResp getOAuthTokenResponce
	err = json.Unmarshal(respBody, &tokenResp)
	if err != nil {
		return OAuthToken{}, fmt.Errorf("(*OAuthService).GetToken - json.Unmarshal: %w", err)
	}

	/*
		access_token также имеет срок жизни (ключ expires_in, в секундах),
		при его истечении приложение должно сделать запрос с refresh_token для получения нового.
	*/
	token := OAuthToken{
		AccessToken:  tokenResp.AccessToken,
		RefreshToken: tokenResp.RefreshToken,
		ExpiresIn:    tokenResp.ExpiresIn,
		TokenType:    tokenResp.TokenType,
		ExpiresAt:    time.Now().Add(time.Second * time.Duration(tokenResp.ExpiresIn)),
	}

	return token, nil
}

func (oas *OAuthServiceImpl) RefreshToken(token OAuthToken) (OAuthToken, error) {
	/*
		access_token также имеет срок жизни (ключ expires_in, в секундах),
		при его истечении приложение должно сделать запрос с refresh_token для получения нового.

		Запрос необходимо делать в application/x-www-form-urlencoded.

		POST https://hh.ru/oauth/token
		В теле запроса необходимо передать дополнительные параметры:

		grant_type=refresh_token
		refresh_token – refresh-токен, полученный ранее при получении пары токенов или прошлом обновлении пары
	*/
	if !token.IsExpired() {
		return token, nil
	}

	body := url.Values{}
	body.Add("grant_type", "refresh_token")
	body.Add("refresh_token", token.RefreshToken)

	req, err := http.NewRequest(http.MethodPost, oas.refreshEndpointURL, strings.NewReader(body.Encode()))
	if err != nil {
		return OAuthToken{}, fmt.Errorf("(*OAuthService).RefreshToken - http.NewRequest: %w", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := oas.httpClient.Do(req)
	if err != nil {
		return OAuthToken{}, fmt.Errorf("(*OAuthService).RefreshToken - http.Client.Do: %w", err)

	}

	if resp.StatusCode == 400 {
		return token, errors.New("(*OAuthService).RefreshToken - token not expired!")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return OAuthToken{}, fmt.Errorf("(*OAuthService).RefreshToken - io.ReadAll: %w", err)

	}
	var tokenResp getOAuthTokenResponce
	err = json.Unmarshal(respBody, &tokenResp)
	if err != nil {
		return OAuthToken{}, fmt.Errorf("(*OAuthService).RefreshToken - json.Unmarshal: %w", err)
	}

	newToken := OAuthToken{
		AccessToken:  tokenResp.AccessToken,
		RefreshToken: tokenResp.RefreshToken,
		ExpiresIn:    tokenResp.ExpiresIn,
		TokenType:    tokenResp.TokenType,
		ExpiresAt:    time.Now().Add(time.Second * time.Duration(tokenResp.ExpiresIn)),
	}

	return newToken, nil
}

func (oas *OAuthServiceImpl) InvalidateToken(token OAuthToken) error {
	/*
		Для того, чтобы инвалидировать текущий access-токен, необходимо сделать запрос:

		DELETE https://api.hh.ru/oauth/token
		Передавая его стандартным способом в заголовке в формате:

		Authorization: Bearer ACCESS_TOKEN

		Инвалидация работает только на действующем access-токене.

		После инвалидации токен нельзя будет запросить с помощью refresh-токена - для работы необходимо будет заново авторизоваться в api.

		Таким образом можно инвалидировать только токен пользователя.
	*/
	req, err := http.NewRequest(http.MethodDelete, oas.baseApiURL+"/oauth/token", strings.NewReader(""))
	if err != nil {
		return fmt.Errorf("(*OAuthService).InvalidateToken - http.NewRequest: %w", err)
	}

	req.Header.Add("Authorization", "Bearer "+token.AccessToken)

	resp, err := oas.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("(*OAuthService).InvalidateToken - http.Client.Do: %w", err)

	}

	if resp.StatusCode != 204 {
		return errors.New("(*OAuthService).InvalidateToken - could not invalidate token!")
	}

	return nil
}
