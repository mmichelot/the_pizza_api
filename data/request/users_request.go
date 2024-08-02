package request

type CreateUsersRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=4,max=20" json:"password"`
}
