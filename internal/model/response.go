package model

// ErrorResponse представляет структуру ответа для ошибки
type ErrorResponse struct {
	Message string `json:"message"`
}

// DeleteResponse представляет структуру ответа для успешного удаления
type DeleteResponse struct {
	Message string `json:"message"`
}
