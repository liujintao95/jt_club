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

type ContactApplication struct {
	ContactId   string `json:"contact_id"`
	ContactType int    `json:"contact_type"`
	Notice      string `json:"notice"`
}

type ContactConfirm struct {
	AppId  string `json:"app_id"`
	Status int    `json:"status"`
}
