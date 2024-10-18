package services

import (
	"account_service/db"
	"account_service/models"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BasketService struct {
	Collection *mongo.Collection
}

func NewBasketService(db *db.MongoDb) *BasketService {
	return &BasketService{
		Collection: db.BasketCollection,
	}
}

func (service BasketService) InsertBasket(basket *models.Basket, ctx context.Context) (primitive.ObjectID, error) {
	res, err := service.Collection.InsertOne(ctx, basket)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	basketId := res.InsertedID.(primitive.ObjectID)
	return basketId, nil
}

func (service BasketService) DeleteBasket(id string, ctx context.Context) error {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = service.Collection.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return err
	}
	return nil
}

func (service BasketService) UpdateBasket(id string, newBasket *models.Basket, ctx context.Context) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = service.Collection.ReplaceOne(ctx, bson.M{"_id": objectId}, newBasket)

	if err != nil {
		return err
	}

	return nil
}

var NotFoundBasketErr error = errors.New("basket not found")

func (service BasketService) GetBasketById(id string, ctx context.Context) (*models.Basket, error) {
	// load from mongo db
	message := models.Basket{}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	res := service.Collection.FindOne(ctx, bson.M{"_id": objectId})

	if res.Err() != nil {
		return nil, res.Err()
	}

	err = res.Decode(&message)
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (service BasketService) GetManyBaskets(limit int, ctx context.Context) ([]models.Basket, error) {
	var messages []models.Basket = make([]models.Basket, 0)

	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))

	cur, err := service.Collection.Find(ctx, bson.D{{}}, findOptions)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		message := models.Basket{}
		if err := cur.Decode(&message); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	err = cur.Close(ctx)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
