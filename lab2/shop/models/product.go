package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type Product struct {
	Id           string          `json:"id"`
	Price        decimal.Decimal `json:"price"`
	IsSold       bool            `json:"isSold"`
	CreationDate time.Time       `json:"creationDate"`
	SoldDate     time.Time       `json:"soldDate"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	Categories   []string        `json:"categories"`
	Material     string          `json:"material"`
}
