package request

type SignInRequest struct {
	Email    string
	Password string
}

type SignUpRequest struct {
	Name     string
	Email    string
	Password string
}
