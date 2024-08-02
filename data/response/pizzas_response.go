package response

type PizzasResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	//Ingredients []xIngredient `json:ingredients`
}
