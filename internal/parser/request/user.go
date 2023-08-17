package request

type SignInRequest struct {
	Account  string
	Password string
}

type SignUpRequest struct {
	Name     string
	Email    string
	Password string
}
