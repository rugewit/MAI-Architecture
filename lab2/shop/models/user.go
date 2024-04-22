package models

import "time"

type User struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Lastname     string    `json:"lastname"`
	Password     string    `json:"password"`
	CreationDate time.Time `json:"creationDate"`
}
