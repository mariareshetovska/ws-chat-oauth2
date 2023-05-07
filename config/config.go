package config

import (
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/linkedin"
)

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}

func SetupLinkedinConfig() *oauth2.Config {
	LoadEnv()
	return &oauth2.Config{
		RedirectURL:  os.Getenv("LINKEDIN_REDIRECT_URL"),
		ClientID:     os.Getenv("LINKEDIN_CLIENT_ID"),
		ClientSecret: os.Getenv("LINKEDIN_CLENT_SECRET"),
		Scopes:       []string{"r_emailaddress", "r_liteprofile"},
		Endpoint:     linkedin.Endpoint,
	}

}

func SetupGoogleConfig() *oauth2.Config {
	LoadEnv()
	return &oauth2.Config{
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLENT_SECRET"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}
