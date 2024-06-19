package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

var GoogleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8080/google/callback",
	ClientID:     "139454420038-erl6ujmciq5g29v3p9htjbu48c1ifhm5.apps.googleusercontent.com",
	ClientSecret: "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",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

var GitHubOauthConfig = &oauth2.Config{
	ClientID:     "Ov23ctrUbMMRRi7XNvLM",
	ClientSecret: "7e2b797a4b406beebb232f191e97bd1b85b1d434",
	RedirectURL:  "http://localhost:8080/github/callback",
	Scopes:       []string{"user:email"},
	Endpoint:     github.Endpoint,
}

var FacebookOauthConfig = &oauth2.Config{
	ClientID:     "831950561657982",
	ClientSecret: "1b9b3a848efa8d2fc69844e3b587b0a2",
	RedirectURL:  "http://localhost:8080/facebook/callback",
	Scopes:       []string{"public_profile"},
	Endpoint:     facebook.Endpoint,
}

var DiscordOauthConfig = &oauth2.Config{
	ClientID:     "1252621445725421569",
	ClientSecret: "XGgUAlkBUydXOVMO3pC5lZhR_XlY09CS",
	RedirectURL:  "http://localhost:8080/discord/callback",
	Scopes:       []string{"identify", "email"},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://discord.com/api/oauth2/authorize",
		TokenURL: "https://discord.com/api/oauth2/token",
	},
}
