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

type UserInfoToken struct {
	ID   int64    `json:"id"`
	Role UserRole `json:"role"`
}

func (u *UserInfoToken) IsEqualRole(role UserRole) bool {
	return u.Role == role
}
