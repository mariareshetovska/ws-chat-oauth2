package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"ws/config"
)

var configLinkedin = config.SetupLinkedinConfig()

func HandleLinkedinLogin(w http.ResponseWriter, r *http.Request) {
	url := configLinkedin.AuthCodeURL(oauthStateString)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleLinkedinCallback(w http.ResponseWriter, r *http.Request) {
	err := getUserInfoLinkedin(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		log.Printf("getUserInfo error: %s", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	if Session["code"] != "" {
		http.Redirect(w, r, "/chat", http.StatusFound)
	}
}

func getUserInfoLinkedin(state string, code string) error {
	if state != oauthStateString {
		return fmt.Errorf("invalid oauth state")
	}

	token, err := configLinkedin.Exchange(context.Background(), code)
	if err != nil {
		return fmt.Errorf("code exchange failed: %s", err.Error())
	}

	client := configLinkedin.Client(context.Background(), token)
	resp, err := client.Get("https://api.linkedin.com/v2/emailAddress?q=members&projection=(elements*(handle~))")
	if err != nil {
		return fmt.Errorf("failed getting user info: %s", err.Error())
	}

	Session["code"] = code
	defer resp.Body.Close()

	return nil
}
