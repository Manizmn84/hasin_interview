package tododto

type TodoCreateResponse struct {
	Title string
}

type GetTodoByIDResponse struct {
	ID     uint
	Np     float64
	Title  string
	Dsc    string
	Status string
}
