package auth

import (
	"encoding/json"
	"net/http"
)

const (
	GOOGLE_AUTH_URL  = "https://accounts.google.com/o/oauth2/v2/auth"
	GOOGLE_TOKEN_URL = "https://oauth2.googleapis.com/token"
	GOOGLE_USER_INFO = "https://www.googleapis.com/oauth2/v3/userinfo"
)

type GoogleUser struct {
	name, email string
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
