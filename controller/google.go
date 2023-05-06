package controller

import (
	"context"
	"fmt"
	"log"
	"ws/config"

	"net/http"
)

var googleOauthConfig = config.SetupGoogleConfig()

var (
	oauthStateString = "random"
	Session          = map[string]string{}
)

func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received Google login request from %s\n", r.RemoteAddr)
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	err := getUserInfoGoogle(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		log.Printf("getUserInfo error: %s", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	if Session["code"] != "" {
		http.Redirect(w, r, "/chat", http.StatusFound)
	}
}

func getUserInfoGoogle(state string, code string) error {
	if state != oauthStateString {
		return fmt.Errorf("invalid oauth state")
	}

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return fmt.Errorf("failed getting user info: %s", err.Error())
	}
	Session["code"] = code
	defer response.Body.Close()

	return nil
}
