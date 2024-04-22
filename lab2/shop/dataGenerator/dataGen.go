package data_generator

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/shopspring/decimal"
	"math/rand"
	"os"
	"reflect"
	"shop/hash"
	"shop/models"
	"strconv"
	"strings"
	"time"
)

const FileSqlPath = "migrations/init.sql"

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

func writeSchemaToFile(filename string, sqlStatements []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, sql := range sqlStatements {
		_, err := file.WriteString(sql + "\n")
		if err != nil {
			return err
		}
	}

	_, err = file.WriteString("\n")
	if err != nil {
		return err
	}
	return nil
}

func createTables() {
	sqlStatements := []string{
		`CREATE TABLE Users (
                       Id VARCHAR(255) PRIMARY KEY,
                       Name VARCHAR(255),
                       Lastname VARCHAR(255),
                       Password VARCHAR(255),
                       CreationDate TIMESTAMP
);`,
	}

	err := writeSchemaToFile(FileSqlPath, sqlStatements)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func writeSQLCommands[T any](file *os.File, tableName string, data []T, keys []string) {
	_, err := file.WriteString(fmt.Sprintf("INSERT INTO %s (", tableName))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	_, err = file.WriteString(strings.Join(keys, ", ") + ")\nVALUES\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	for i, item := range data {
		if i > 0 {
			_, err := file.WriteString(",\n")
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		}
		itemValue := reflect.ValueOf(item)
		var values []string
		for j := 0; j < itemValue.NumField(); j++ {
			field := itemValue.Field(j)
			var value string
			switch field.Interface().(type) {
			case string:
				value = fmt.Sprintf("'%s'", field)
			case decimal.Decimal:
				value = field.Interface().(decimal.Decimal).String()
			case time.Time:
				value = fmt.Sprintf("'%s'", field.Interface().(time.Time).Format("2006-01-02 15:04:05"))
			case bool:
				value = fmt.Sprintf("%t", field.Interface().(bool))
			case []string:
				arr := field.Interface().([]string)
				var quoted []string
				for _, el := range arr {
					quoted = append(quoted, strconv.Quote(el))
				}
				value = fmt.Sprintf("'{%s}'", strings.Join(quoted, ", "))
			case map[string]int:
				dict := field.Interface().(map[string]int)
				dictJson, err := json.Marshal(dict)
				if err != nil {
					fmt.Println("error marshelling to json")
					return
				}
				value = fmt.Sprintf("'%s'", string(dictJson))
			default:
				value = fmt.Sprintf("'%v'", field)
			}
			values = append(values, value)
		}
		_, err := file.WriteString(fmt.Sprintf("(%s)", strings.Join(values, ", ")))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
	_, err = file.WriteString(";\n\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func GenFakeData(userCount, productCount int) {
	rand.New(rand.NewSource(time.Now().Unix()))

	// populate users
	users := make([]models.User, userCount)
	for i := 0; i < userCount; i++ {
		users[i].Id = gofakeit.UUID()
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

	// export to sql
	createTables()

	file, err := os.OpenFile(FileSqlPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	keys, err := getFieldNamesAsStringSlice(users[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	writeSQLCommands(file, "Users", users, keys)
}
