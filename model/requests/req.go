package requests

type Register struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Comment struct {
	Content string `json:"content"`
}
