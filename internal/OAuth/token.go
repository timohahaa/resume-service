package oauth

import "time"

type OAuthToken struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64
	TokenType    string
	ExpiresAt    time.Time
}

func (t OAuthToken) IsExpired() bool {
	return time.Until(t.ExpiresAt) > 0
}
