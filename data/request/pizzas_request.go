package request

type CreatePizzasRequest struct {
	Name        string `validate:"required,min=1,max=200" json:"name"`
	Description string `validate:"max=200" json:"description"`
}

type UpdatePizzasRequest struct {
	ID          uint   `validate:"required"`
	Name        string `validate:"required,max=200,min=1" json:"name"`
	Description string `validate:"max=200" json:"description"`
}
