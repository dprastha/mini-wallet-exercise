package pkg

type BaseResponse struct {
	Status string       `json:"status"`
	Data   *interface{} `json:"data"`
}

type ErrorValidation struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

type ErrorMessage struct {
	Error string `json:"error"`
}
