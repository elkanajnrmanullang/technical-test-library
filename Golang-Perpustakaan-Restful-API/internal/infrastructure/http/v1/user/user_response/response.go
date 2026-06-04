package user_response

var (
	ErrInvalidFormatJson = "Data json yang diberikan salah!. Tidak sesuai format"
)

type LoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func NewLoginResponse(username string, token string, refreshToken string) *LoginResponse {
	var loginResponse LoginResponse
	loginResponse.Username = username
	loginResponse.Token = token
	loginResponse.RefreshToken = refreshToken
	return &loginResponse
}
