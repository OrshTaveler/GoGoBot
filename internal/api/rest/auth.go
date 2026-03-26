package rest

import (
	"encoding/json"
	"fmt"
	"gogobot/internal/api/shared"
	"net/http"
	"net/url"
	"strings"
)

func GetAccessToken(r *http.Request) (string, error) {
	data := url.Values{}

	// Setting type of authorization
	data.Set("grant_type", "authorization_code")
	// Setting code that server gave us after auth
	data.Set("code", r.URL.Query().Get("code"))
	// Setting redirect link
	data.Set("redirect_uri", "http://"+r.Host+r.URL.Path)

	req, _ := http.NewRequest("POST", "https://online-go.com/oauth2/token/", strings.NewReader(data.Encode()))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(shared.CLIENT_ID, shared.CLIENT_SCREET)

	resp, err := http.DefaultClient.Do(req)

	if resp == nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("response not ok")
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result["access_token"].(string), nil
}

func GetJWT(accessToken string) (string, error) {
	req, _ := http.NewRequest("GET", "https://online-go.com/api/v1/ui/config", nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := http.DefaultClient.Do(req)
	if resp == nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("response not ok")
	}

	type usr struct {
		UserJWT string `json:"user_jwt"`
	}
	var config usr
	json.NewDecoder(resp.Body).Decode(&config)

	return config.UserJWT, nil
}
