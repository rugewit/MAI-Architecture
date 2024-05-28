package data_generator

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
	"os"
	"reflect"
	"shop/hash"
	"shop/models"
	"time"
)

const FileJsonPath = "migrations/records.json"

func pickNItems[T any](array []T, n int) []T {
	output := make([]T, n)
	for i := 0; i < n; i++ {
		output[i] = array[rand.Intn(len(array))]
	}
	return output
}

func getFieldNamesAsStringSlice(obj any) ([]string, error) {
	val := reflect.ValueOf(obj)
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input is not a struct")
	}

	var names []string
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		names = append(names, field.Name)
	}

	return names, nil
}

func GenFakeData(userCount, productCount int) {
	rand.New(rand.NewSource(time.Now().Unix()))

	// populate users
	users := make([]models.User, userCount)
	for i := 0; i < userCount; i++ {
		objectId := primitive.NewObjectID()
		users[i].Id = objectId
		users[i].Login = gofakeit.Username()
		users[i].Name = gofakeit.FirstName()
		users[i].Lastname = gofakeit.LastName()
		openPass := gofakeit.Password(true, true, false, false,
			false, 7)
		hashedPass, err := hash.HashPassword(openPass)
		if err != nil {
			fmt.Println(err)
			return
		}
		users[i].Password = hashedPass
		users[i].CreationDate = gofakeit.PastDate()
	}

	// populate products
	products := make([]models.Product, productCount)
	for i := 0; i < productCount; i++ {
		objectId := primitive.NewObjectID()
		products[i].Id = objectId
		products[i].Name = gofakeit.Product().Name
		var err error
		products[i].Price = gofakeit.Product().Price
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// populate baskets
	baskets := make([]models.Basket, userCount)
	for i := 0; i < userCount; i++ {
		objectId := primitive.NewObjectID()
		baskets[i].Id = objectId
		itemsCount := rand.Intn(4)
		baskets[i].Products = pickNItems(products, itemsCount)
		totalPrice := 0.0
		for j := 0; j < itemsCount; j++ {
			totalPrice += baskets[i].Products[j].Price
		}
		baskets[i].TotalPrice = totalPrice
	}

	// assign basket to users and vice versa
	for i := 0; i < userCount; i++ {
		users[i].BasketId = baskets[i].Id
		baskets[i].UserId = users[i].Id
	}

	data := map[string]any{
		"users":    users,
		"products": products,
		"baskets":  baskets,
	}

	file, err := os.OpenFile(FileJsonPath, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		fmt.Printf("Failed to encode data: %v", err)
		return
	}

}
