package model

type Type string

const (
	Cat Type = "cat"
	Dog Type = "dog"
)

type Pet struct {
	Age   int     `json:"age"`
	Type  Type    `json:"type"`
	Price float64 `json:"price"`
}
