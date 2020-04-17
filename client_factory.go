package scoro

import (
	"context"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"log"
	"os"
)

const envSiteURL = "SITE_URL"
const envClientID = "CLIENT_ID"
const envClientSecret = "CLIENT_SECRET"
const envScope = "SCOPE"
const envRedirectUrl = "REDIRECT_URL"
const envLang = "API_LANG"

type ApiClientActions interface {
	HandleAuthorization(oauth2.Config) string
	SaveTokens(*oauth2.Token)
	FetchTokens() *oauth2.Token
}

type ApiConfig struct {
	oauthConfig oauth2.Config
	siteUrl     string
}

//GetClient return http.APIClient using the provided oauth2.Config
func GetClient(config ApiConfig, client interface{ ApiClientActions }) APIClient {
	ctx := context.Background()
	tok := client.FetchTokens()
	if len(tok.AccessToken) == 0 {
		connect(config.oauthConfig, client)
		tok = client.FetchTokens()
	}

	tokenSource := config.oauthConfig.TokenSource(ctx, tok)
	newToken, err := tokenSource.Token()
	if err != nil {
		log.Fatalln(err)
	}

	if newToken.AccessToken != tok.AccessToken {
		client.SaveTokens(newToken)
		tok = newToken
	}

	return APIClient{config, config.oauthConfig.Client(ctx, tok)}
}

//GetAPIClientConfigFromEnvFile returns oauth2.Config from .env file
func GetAPIClientConfigFromEnvFile() ApiConfig {
	siteURL := getEnvVariable(envSiteURL)
	clientID := getEnvVariable(envClientID)
	clientSecret := getEnvVariable(envClientSecret)
	scope := getEnvVariable(envScope)
	redirectURL := getEnvVariable(envRedirectUrl)
	language := getEnvVariable(envLang)

	if len(clientID) == 0 || len(clientSecret) == 0 || len(language) == 0 || len(siteURL) == 0 || len(scope) == 0 || len(redirectURL) == 0 {
		log.Fatalf("Error loading values from .env file")
	}

	oauthConf := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{scope},
		Endpoint: oauth2.Endpoint{
			TokenURL: siteURL + "/api/" + apiVersion + "/token",
			AuthURL:  siteURL + "/api/" + apiVersion + "/auth",
		},
		RedirectURL: redirectURL,
	}
	return ApiConfig{oauthConf, siteURL}
}

func getEnvVariable(key string) string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	err = godotenv.Load(path + "/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func connect(config oauth2.Config, client ApiClientActions) {
	ctx := context.Background()
	authorizationCode := client.HandleAuthorization(config)

	if len(authorizationCode) == 0 {
		log.Fatalf("Error getting authorization code")
	}

	tok, err := config.Exchange(ctx, authorizationCode)
	if err != nil {
		log.Fatal(err)
	}

	client.SaveTokens(tok)
}
