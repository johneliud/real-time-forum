package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/johneliud/forum/backend/util"
)

const (
	GOOGLE_AUTH_URL  = "https://accounts.google.com/o/oauth2/v2/auth"
	GOOGLE_TOKEN_URL = "https://oauth2.googleapis.com/token"
	GOOGLE_USER_INFO = "https://www.googleapis.com/oauth2/v3/userinfo"
)

type GoogleUser struct {
	Name, Email string
}

/*
Initiates the Google OAuth 2.0 flow, sending the user to Google's authorization page to grant access.
*/
func GoogleAuth(w http.ResponseWriter, r *http.Request) {
	state := generateStateCookie(w)

	googleClientID, _ := util.LoadCredentials()

	// Construct the Google OAuth 2.0 authorization URL with necessary parameters
	redirectURL := fmt.Sprintf(
		"%s?client_id=%s&redirect_uri=%s&response_type=code&scope=openid email profile&state=%s&prompt=select_account&access_type=offline",
		GOOGLE_AUTH_URL,
		googleClientID,
		url.QueryEscape(REDIRECT_URL+"/auth/google/callback"),
		state,
	)
	w.Header().Set("Access-Control-Allow-Origin", REDIRECT_URL)
	http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
}

/*
Sends a GET request to the Google User Info API with the provided token, decodes the JSON response into a GoogleUser object, and returns it.
*/
func getGoogleUser(token string) (*GoogleUser, error) {
	req, _ := http.NewRequest("GET", GOOGLE_USER_INFO, nil)
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var user GoogleUser
	json.NewDecoder(resp.Body).Decode(&user)
	return &user, nil
}
