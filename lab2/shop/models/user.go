package models

import "time"

type User struct {
	Id           string
	Name         string
	Lastname     string
	Password     string
	CreationDate time.Time
	BasketId     string
}
