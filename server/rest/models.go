package rest

type Event struct {
	ID         int    `json:"id"`
	Government string `json:"government"`
	Event      string `json:"event"`
	Name       string `json:"name"`
	Position   string `json:"position"`
	Address    string `json:"address"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Comment    string `json:"comment"`
}
