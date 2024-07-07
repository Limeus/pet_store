package model

type UpdatePet struct {
	Price       float64 `json:"price"`
	Description string  `json:"description,omitempty"`
}
