package services

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"shop/db"
	"shop/models"
)

type ProductService struct {
	Collection *mongo.Collection
}

func NewProductService(db *db.MongoDb) *ProductService {
	return &ProductService{
		Collection: db.ProductsCollection,
	}
}

func (service ProductService) InsertProduct(product *models.Product, ctx context.Context) (primitive.ObjectID, error) {
	res, err := service.Collection.InsertOne(ctx, product)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	insertedID := res.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func (service ProductService) DeleteProduct(id string, ctx context.Context) error {
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

func (service ProductService) UpdateProduct(id string, newProduct *models.Product, ctx context.Context) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = service.Collection.ReplaceOne(ctx, bson.M{"_id": objectId}, newProduct)

	if err != nil {
		return err
	}

	return nil
}

var NotFoundProductErr error = errors.New("product not found")

func (service ProductService) GetProductById(id string, ctx context.Context) (*models.Product, error) {
	// load from mongo db
	message := models.Product{}
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

func (service ProductService) GetManyProducts(limit int, ctx context.Context) ([]models.Product, error) {
	var messages []models.Product = make([]models.Product, 0)

	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))

	cur, err := service.Collection.Find(ctx, bson.D{{}}, findOptions)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		message := models.Product{}
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
