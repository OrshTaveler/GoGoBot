package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetUserInfo(accessToken string) (string, float64, error) {
	req, _ := http.NewRequest("GET", "https://online-go.com/api/v1/me/", nil)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := http.DefaultClient.Do(req)

	if resp == nil {
		return "", -1, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", -1, fmt.Errorf("response status not ok")
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", -1, err
	}

	return result["username"].(string), result["id"].(float64), nil
}
