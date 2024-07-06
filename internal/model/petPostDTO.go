package model

type PostPet struct {
	Age         int     `json:"age"`
	Type        Type    `json:"type"`
	Price       float64 `json:"price"`
	Description string  `json:"description,omitempty"`
}
