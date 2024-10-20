package model

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Phone       string `json:"phone"`
	Photo       string `json:"photo"`
	LinkTreeURL string `json:"link_tree_url"`
	Bio         string `json:"bio"`
}
