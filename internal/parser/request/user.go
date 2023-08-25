package request

type SignInRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSelect struct {
	Query     string `json:"query"`
	IsContact bool   `json:"is_contact"`
}

type ContactRequest struct {
	Uid    string `json:"uid"`
	Notice string `json:"notice"`
}

type ContactCommit struct {
	RequestId string `json:"request_id"`
	Status    string `json:"status"`
}
