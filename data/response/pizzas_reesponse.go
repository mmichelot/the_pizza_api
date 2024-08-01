package response

type PizzasResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	//Ingredients []xIngredient `json:ingredients`
}
