package models

type SignUpUser struct {
	Name     string `json:"name" example:"Alex"`
	Lastname string `json:"lastname" example:"Ivanov"`
	Password string `json:"password" example:"qwerty"`
	Login    string `json:"login" example:"AlexHere"`
}
