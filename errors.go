package castos

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"error"`
	Status  int    `json:"status"`
}

func (e ErrorResponse) Error() string {
	return e.Message
}
