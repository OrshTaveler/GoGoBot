package utils

import (
	"fmt"
	"gogobot/internal/api/shared"
)

func GenerateAuthURL() string {
	url := fmt.Sprintf("https://online-go.com/oauth2/authorize/?client_id=%s&response_type=code&redirect_uri=%s%s%s&scope=read+write&state=RANDOM_STATE",
		shared.CLIENT_ID, shared.REDIRECT_URI, shared.PORT, shared.AUTH_ENDPOINT)

	return url
}
