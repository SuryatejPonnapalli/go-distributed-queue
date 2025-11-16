package users

type RegisterResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type LoginResponse struct {
	Email string `json:"email"`	
	Token string `json:"token"`
}
