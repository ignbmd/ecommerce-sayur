package request

type SignInRequests struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"min=8,required"`
}
