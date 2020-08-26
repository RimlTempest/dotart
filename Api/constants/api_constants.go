package constants

const (
	API_SERVER = ":8080"
	// API
	API         string = "/api"
	API_VERSION string = "/v1"

	// CROS ORIGIN
	api_local               string = "http://localhost:3000"
	API_ENDPOINT            string = api_local //"*"
	API_ENDPOINT_LOGIN      string = api_local + "/login"
	API_ENDPOINT_SNS        string = api_local + "/sns"
	API_ENDPOINT_INFOMATION string = api_local + "/infomaton"
	API_ENDPOINT_CREATOR    string = api_local + "/creator"

	// JWT
	SESSION_SECRET  = "dot_art_secret"
	REALM           = "dotart_api"
	IDENTITY_KEY    = "id"
	TOKEN_HEAD_NAME = "Bearer" //"X-DotArt-Access"
	EXPIRE          = "expire"
	TOKEN           = "token"
	TokenLength     = 64
)
