package models

type AppErr struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Error   error  `json:"error"`
}

func NewAppErr(msg string, code int, err error) *AppErr {
	return &AppErr{Message: msg, Code: code, Error: err}
}
