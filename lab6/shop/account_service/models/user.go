package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id           primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	BasketId     primitive.ObjectID `json:"basket_id" bson:"basket_id"`
	Login        string             `json:"login" bson:"login"`
	Name         string             `json:"name" bson:"name"`
	Lastname     string             `json:"lastname" bson:"lastname"`
	Password     string             `json:"password" bson:"password"`
	CreationDate time.Time          `json:"creationDate" bson:"creationDate"`
}
