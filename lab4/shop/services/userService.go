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

type UserService struct {
	Collection *mongo.Collection
}

func NewUserService(db *db.MongoDb) *UserService {
	return &UserService{
		Collection: db.UserCollection,
	}
}

func (service UserService) InsertUser(user *models.User, ctx context.Context) (primitive.ObjectID, error) {
	res, err := service.Collection.InsertOne(ctx, user)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	insertedId := res.InsertedID.(primitive.ObjectID)
	return insertedId, nil
}

func (service UserService) DeleteUser(id string, ctx context.Context) error {
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

func (service UserService) UpdateUser(id string, newUser *models.User, ctx context.Context) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = service.Collection.ReplaceOne(ctx, bson.M{"_id": objectId}, newUser)

	if err != nil {
		return err
	}

	return nil
}

var NotFoundUserErr error = errors.New("user not found")

func (service UserService) GetUserById(id string, ctx context.Context) (*models.User, error) {
	// load from mongo db
	message := models.User{}
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

func (service UserService) GetManyUsers(limit int, ctx context.Context) ([]models.User, error) {
	var messages []models.User = make([]models.User, 0)

	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))

	cur, err := service.Collection.Find(ctx, bson.D{{}}, findOptions)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		message := models.User{}
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

func (service UserService) CheckUserIDExists(id string, ctx context.Context) (bool, error) {
	return true, nil
}

// .* matches any characters (except for line terminators) zero or more times
// . matches any one character
func (service UserService) PatternSearchUsers(nameMask, lastnameMask string, limit int,
	ctx context.Context) ([]models.User, error) {

	// db.users.find({'name': {'$regex': '.*sometext.*'}})

	// Construct regular expressions for name and lastname masks
	//nameRegex := ".*" + nameMask + ".*"
	//lastnameRegex := ".*" + lastnameMask + ".*"

	// Construct the MongoDB query
	query := bson.M{
		"$and": []bson.M{
			{"name": bson.M{"$regex": nameMask, "$options": "i"}},         // Case-insensitive regex for name
			{"lastname": bson.M{"$regex": lastnameMask, "$options": "i"}}, // Case-insensitive regex for lastname
		},
	}

	// Set options for limiting results
	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))

	// Execute the query
	cur, err := service.Collection.Find(ctx, query, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	// Decode results into models.User slice
	var users []models.User
	for cur.Next(ctx) {
		var user models.User
		if err := cur.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (service UserService) GetUserByLogin(login string, ctx context.Context) (*models.User, error) {
	user := models.User{}
	res := service.Collection.FindOne(ctx, bson.M{"login": login})

	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return nil, NotFoundUserErr
		}
		return nil, res.Err()
	}

	err := res.Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
