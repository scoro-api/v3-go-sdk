package scoro

import (
	"context"
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

// GetClient return http.APIClient using the provided oauth2.Config
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

	return APIClient{config: config, httpClient: config.oauthConfig.Client(ctx, tok), customHeaders: make(map[string]string)}
}

// GetAPIClientConfig returns oauth2.Config
func GetAPIClientConfig(siteURL string, clientID string, clientSecret string, scope string, redirectURL string, language string) ApiConfig {
	if len(clientID) == 0 || len(clientSecret) == 0 || len(language) == 0 || len(siteURL) == 0 || len(scope) == 0 || len(redirectURL) == 0 {
		log.Fatalf("configuration error: SITE_URL, CLIENT_ID, CLIENT_SECRET, SCOPE, REDIRECT_URL and API_LANG must be set")
	}

	oauthConf := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{scope},
		Endpoint: oauth2.Endpoint{
			TokenURL: siteURL + "/api/" + apiVersion + "/token",
			AuthURL:  siteURL + "/apiAuth",
		},
		RedirectURL: redirectURL,
	}
	return ApiConfig{oauthConf, siteURL}
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
