package onboard

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Type     string `json:"type" validate:""`
}

type Loginresp struct {
	Token string `json:"token"`
}
