package back

type AuthenticateInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type VerifyEmail struct {
	Token string `form:"token" validate:"required"`
}
