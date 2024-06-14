package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8080/callback",
	ClientID:     "139454420038-erl6ujmciq5g29v3p9htjbu48c1ifhm5.apps.googleusercontent.com",
	ClientSecret: "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",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}
