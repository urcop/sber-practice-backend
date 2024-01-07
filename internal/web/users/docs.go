package users

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
	StatusCode  int    `json:"statusCode"`
	Message     string `json:"message"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseError struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}
