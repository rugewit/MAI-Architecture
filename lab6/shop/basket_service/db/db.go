package db

import (
	"basket_service/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"time"
)

type MongoDb struct {
	Port           string
	User           string
	Password       string
	DbName         string
	DbUrl          string
	CollectionName string
	//Ctx            context.Context
	Client             *mongo.Client
	UserCollection     *mongo.Collection
	BasketCollection   *mongo.Collection
	ProductsCollection *mongo.Collection
}

func NewMongoDb(pgConfig *config.MongoConfig) (*MongoDb, error) {
	dbPort := pgConfig.Port
	dbUser := pgConfig.User
	dbPassword := pgConfig.Password
	dbName := pgConfig.Name
	dbHost := pgConfig.Host
	dbUrl := fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	dbUrl = pgConfig.Uri

	clientOptions := options.Client().ApplyURI(dbUrl)

	// Настройка WriteConcern для надежной записи
	wc := writeconcern.New(writeconcern.WMajority(), writeconcern.WTimeout(5*time.Second))
	clientOptions.SetWriteConcern(wc)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	// ping mongo db
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}

	userCollection := client.Database(dbName).Collection(pgConfig.UsersCollection)
	basketCollection := client.Database(dbName).Collection(pgConfig.BasketsCollection)
	productsCollection := client.Database(dbName).Collection(pgConfig.ProductsCollection)

	mongoDb := &MongoDb{
		Port:     dbPort,
		User:     dbUser,
		Password: dbPassword,
		DbName:   dbName,
		DbUrl:    dbUrl,
		//Ctx:            ctx,
		Client:             client,
		UserCollection:     userCollection,
		BasketCollection:   basketCollection,
		ProductsCollection: productsCollection,
	}
	return mongoDb, nil
}

func (db *MongoDb) Close(ctx context.Context) {

}
