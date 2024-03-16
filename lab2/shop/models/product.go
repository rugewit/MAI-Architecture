package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type Product struct {
	Id           string
	Price        decimal.Decimal
	IsSold       bool
	CreationDate time.Time
	SoldDate     time.Time
	Name         string
	Description  string
	Categories   []string
	Material     string
}
