package tododto

type TodoCreateRequest struct {
	Np    float64
	Title string
	Dsc   string
}


type TodoUpdateRequest struct {
	ID    uint
	Np    float64
	Title string
	Dsc   string
}
