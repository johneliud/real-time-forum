package auth

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/johneliud/forum/backend/util"
	"github.com/johneliud/forum/database"
)

const (
	GOOGLE_AUTH_URL  = "https://accounts.google.com/o/oauth2/v2/auth"
	GOOGLE_TOKEN_URL = "https://oauth2.googleapis.com/token"
	GOOGLE_USER_INFO = "https://www.googleapis.com/oauth2/v3/userinfo"
)

type GoogleUser struct {
	Name, Email, Sub string
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
Google OAuth 2.0 callback handler called after a user grants access to their Google account. The function:

1. Validates the state parameter to prevent CSRF attacks.
2. Exchanges the authorization code for an access token.
3. Retrieves the user's information using the access token.
4. Checks if the user already exists in the database and creates a new user if not.
5. Redirects the user to the  page, indicating whether they're a new or returning user.
*/
func GoogleCallback(w http.ResponseWriter, r *http.Request) {
	if err := validateState(r); err != nil {
		log.Printf("State validation failed: %v", err)
		http.Redirect(w, r, "/sign-in?error=invalid_state", http.StatusTemporaryRedirect)
		return
	}

	code := r.URL.Query().Get("code")
	token, err := exchangeGoogleToken(code)
	if err != nil {
		log.Printf("Token exchange failed: %v\n", err)
		http.Redirect(w, r, "/sign-in?error=token_exchange_failed", http.StatusTemporaryRedirect)
		return
	}

	user, err := getGoogleUser(token)
	if err != nil {
		log.Printf("Failed to get user info: %v\n", err)
		http.Redirect(w, r, "/sign-in?error=user_info_failed", http.StatusTemporaryRedirect)
		return
	}

	var (
		userID       int
		authProvider string
	)

	err = database.DB.QueryRow("SELECT id, auth_provider FROM tblUsers WHERE email = ?", user.Email).Scan(&userID, &authProvider)
	// Email exists but has a different oauth provider
	if err == nil && authProvider != "google" {
		log.Printf("Email already registered with %s: %v", authProvider, user.Email)
		http.Redirect(w, r, "/sign-in?error=email_exists&provider="+authProvider, http.StatusTemporaryRedirect)
		return
	}

	isNewUser := false
	// Create a new user if not existing
	if errors.Is(err, sql.ErrNoRows) {
		var count int
		err = database.DB.QueryRow("SELECT COUNT(*) FROM tblUsers WHERE username = ?", user.Name).Scan(&count)
		if err != nil {
			log.Printf("Database error checking username: %v", err)
			http.Redirect(w, r, "/sign-in?error=database_error", http.StatusTemporaryRedirect)
			return
		}

		if count > 0 {
			// Append a random suffix if a username is taken
			user.Name = fmt.Sprintf("%s_%s", user.Name, user.Sub[:6])
		}

		// Create new user
		result, err := database.DB.Exec(
			"INSERT INTO tblUsers(username, email, auth_provider) VALUES(?, ?, ?)",
			user.Name, user.Email, "google",
		)
		if err != nil {
			log.Printf("User creation failed: %v", err)
			http.Redirect(w, r, "/sign-in?error=user_creation_failed", http.StatusTemporaryRedirect)
			return
		}
		id, _ := result.LastInsertId()
		userID = int(id)
		isNewUser = true
	} else if err != nil {
		log.Printf("Database error: %v", err)
		http.Redirect(w, r, "/sign-in?error=database_error", http.StatusTemporaryRedirect)
		return
	}

	if isNewUser {
		http.Redirect(w, r, "/?status=new_user", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/?status=returning_user", http.StatusSeeOther)
	}
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

/*
Exchanges a Google authorization code for an access token. It sends a POST request to the Google Token URL.
*/
func exchangeGoogleToken(code string) (string, error) {
	googleClientID, googleClientSecret := util.LoadCredentials()

	data := url.Values{
		"code":          {code},
		"client_id":     {googleClientID},
		"client_secret": {googleClientSecret},
		"redirect_uri":  {REDIRECT_URL + "/auth/google/callback"},
		"grant_type":    {"authorization_code"},
	}

	resp, err := http.PostForm(GOOGLE_TOKEN_URL, data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		AccessToken string `json:"access_token"`
	}
	json.NewDecoder(resp.Body).Decode(&result)
	return result.AccessToken, nil
}
