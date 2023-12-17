package main

import (
	"fmt"
	"log"

	"github.com/timohahaa/resume-service/config"
	"github.com/timohahaa/resume-service/internal/OAuth"
)

func main() {
	config, err := config.NewConfig("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	code := "KR1V6F7EP8A54R5KLT17VP68918H21E1NGJEPQ4RIL84JQOHT5354CLIBDC31FM2"

	oas := oauth.NewOAuthService(config.OAuthEndpointURL, config.CliendID, config.ClientSecret, config.RedirectURI, config.BaseApiUrl)

	token, err := oas.GetToken(code)
	fmt.Printf("token: %+v\n", token)
	fmt.Println(err)

	token, err = oas.RefreshToken(token)
	fmt.Printf("after refresh: %+v\n", token)
	fmt.Println(err)

	err = oas.InvalidateToken(token)
	fmt.Println(err)
}
