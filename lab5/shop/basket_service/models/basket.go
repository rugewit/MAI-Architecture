package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Basket struct {
	Id         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserId     primitive.ObjectID `json:"userId" bson:"userId"`
	Products   []Product          `json:"products" bson:"products"`
	TotalPrice float64            `json:"totalPrice" bson:"totalPrice"`
}
