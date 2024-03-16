package models

import "github.com/shopspring/decimal"

type Basket struct {
	Id         string
	UserId     string
	TotalPrice decimal.Decimal
	CountMap   map[string]int
	ItemsMap   map[string]Product
}
