package handler

type (
	SignUpRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	SignInRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)
