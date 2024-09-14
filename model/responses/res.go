package responses

type Register struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Login struct {
	Token string `json:"token"`
}
