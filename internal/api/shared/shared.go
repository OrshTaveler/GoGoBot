package shared

const CLIENT_ID = "GNLieTm6TsaOJSGIKimALXfC4QNVpFvF1DzWXXxU"
const CLIENT_SCREET = "fjCqsdMz8kA7Qsgaffh9D4Y99pDYqWwsYPw5tENgRxneUi3BKTwevIpx23p8LXf1aCbtTGG9WVzXF14y9yau8MnSLDgTWFHuMeMGZeFHJWR3FBBH6E7ZE003JBV42Gq5"

const PORT = ":8080"
const REDIRECT_URI = "http://localhost"
const AUTH_ENDPOINT = "/auth"

type Player struct {
	Token    string
	Username string
	JWT      string
	UserId   float64
}
