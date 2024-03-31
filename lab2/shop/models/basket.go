package models

import "github.com/shopspring/decimal"

type Basket struct {
	Id         string             `json:"id"`
	UserId     string             `json:"userId"`
	TotalPrice decimal.Decimal    `json:"totalPrice"`
	CountMap   map[string]int     `json:"countMap"`
	ItemsMap   map[string]Product `json:"itemsMap"`
}
