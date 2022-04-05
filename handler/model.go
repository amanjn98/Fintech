package handler

type User struct {
	Username string
	Password string
}

type ErrorResponse struct {
	Err error `json:"-"`
	Message string `json:"message"`
}