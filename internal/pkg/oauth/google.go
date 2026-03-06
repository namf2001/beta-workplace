package oauth

import (
	"github.com/namf2001/beta-workplace/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	GoogleOauthConfig *oauth2.Config
	StateString       = "random-string"
	Scopes            = []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"}
)

func Init() {
	cfg := config.GetConfig()
	GoogleOauthConfig = &oauth2.Config{
		RedirectURL:  cfg.GoogleRedirectURL,
		ClientID:     cfg.GoogleClientID,
		ClientSecret: cfg.GoogleClientSecret,
		Scopes:       Scopes,
		Endpoint:     google.Endpoint,
	}
}
