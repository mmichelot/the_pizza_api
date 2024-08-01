package request

type CreatePizzasRequest struct {
	Name        string `validate:"required,min=1,max=200" json:"name"`
	Description string `validate:"max=200" json:"description"`
}
