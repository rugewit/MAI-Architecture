package services

import (
	"context"
	"errors"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/jackc/pgx/v5"
	"log"
	"shop/db"
	"shop/models"
)

type UserService struct {
	Db *db.PostgresDb
}

func NewUserService(db *db.PostgresDb) *UserService {
	return &UserService{
		Db: db,
	}
}

func (service UserService) InsertUser(user *models.User, ctx context.Context) error {
	conn, err := service.Db.Pool.Acquire(ctx)
	if err != nil {
		log.Println("cannot acquire a database connection", err)
		return err
	}
	defer conn.Release()

	userId := gofakeit.UUID()

	user.Id = userId
	user.CreationDate = gofakeit.Date()

	// Insert into Users
	query := `INSERT INTO Users 
    (Id, Name, Lastname, Password, CreationDate)
	values($1, $2, $3, $4, $5)`
	_, err = conn.Exec(ctx, query, user.Id, user.Name, user.Lastname, user.Password, user.CreationDate)
	if err != nil {
		log.Println("cannot insert into Users", err)
		return err
	}
	return nil
}

func (service UserService) DeleteUser(id string, ctx context.Context) error {
	conn, err := service.Db.Pool.Acquire(ctx)
	if err != nil {
		log.Println("cannot acquire a database connection", err)
		return err
	}
	defer conn.Release()

	exists, err := service.CheckUserIDExists(id, ctx)
	if err != nil {
		return err
	}
	if !exists {
		return NotFoundUserErr
	}
	// Delete from Users
	query := `DELETE FROM Users WHERE Id = $1`
	_, err = conn.Exec(ctx, query, id)
	if err != nil {
		log.Println("cannot delete user", err)
		return err
	}
	return nil
}

func (service UserService) UpdateUser(id string, newUser *models.User, ctx context.Context) error {
	conn, err := service.Db.Pool.Acquire(ctx)
	if err != nil {
		log.Println("cannot acquire a database connection", err)
		return err
	}
	defer conn.Release()

	// Update Users
	query := `UPDATE Users SET Name = $1, Lastname = $2, Password = $3, CreationDate = $4 WHERE Id = $5`
	_, err = conn.Exec(ctx, query, newUser.Name, newUser.Lastname, newUser.Password, newUser.CreationDate, id)
	if err != nil {
		log.Println("cannot update user", err)
		return err
	}
	return nil
}

var NotFoundUserErr error = errors.New("user not found")

func (service UserService) GetUserById(id string, ctx context.Context) (*models.User, error) {
	conn, err := service.Db.Pool.Acquire(ctx)
	if err != nil {
		log.Println("cannot acquire a database connection", err)
		return nil, err
	}
	defer conn.Release()

	// Query to get user by ID
	query := `SELECT Id, Name, Lastname, Password, CreationDate FROM Users WHERE Id = $1`
	row := conn.QueryRow(ctx, query, id)

	// Initialize a user object to store the result
	user := models.User{}

	// Scan the row into user object
	err = row.Scan(&user.Id, &user.Name, &user.Lastname, &user.Password, &user.CreationDate)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// Return nil if no user found with the given ID
			return nil, NotFoundUserErr
		}
		log.Println("cannot get user by ID", err)
		return nil, err
	}

	return &user, nil
}

func (service UserService) GetManyUsers(limit int, ctx context.Context) ([]models.User, error) {
	conn, err := service.Db.Pool.Acquire(ctx)
	if err != nil {
		log.Println("cannot acquire a database connection", err)
		return nil, err
	}
	defer conn.Release()

	// Query to get many users with limit
	query := `SELECT Id, Name, Lastname, Password, CreationDate FROM Users LIMIT $1`
	rows, err := conn.Query(ctx, query, limit)
	if err != nil {
		log.Println("cannot get many users", err)
		return nil, err
	}
	defer rows.Close()

	// Initialize a slice to store user objects
	var users []models.User

	// Iterate through the rows and scan each user into the slice
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Name, &user.Lastname, &user.Password, &user.CreationDate)
		if err != nil {
			log.Println("error scanning user row", err)
			return nil, err
		}
		users = append(users, user)
	}

	// Check for errors during rows iteration
	if err = rows.Err(); err != nil {
		log.Println("error iterating user rows", err)
		return nil, err
	}

	return users, nil
}

func (service UserService) CheckUserIDExists(id string, ctx context.Context) (bool, error) {
	conn, err := service.Db.Pool.Acquire(ctx)
	if err != nil {
		log.Println("cannot acquire a database connection", err)
		return false, err
	}
	defer conn.Release()

	// Query to check if id exists in the Users table
	query := "SELECT EXISTS(SELECT 1 FROM Users WHERE Id = $1)"
	var exists bool
	err = conn.QueryRow(ctx, query, id).Scan(&exists)
	if err != nil {
		log.Println("cannot check if ID exists", err)
		return false, err
	}

	return exists, nil
}

func (service UserService) PatternSearchUsers(nameMask, lastnameMask string, limit int,
	ctx context.Context) ([]models.User, error) {
	conn, err := service.Db.Pool.Acquire(ctx)
	if err != nil {
		log.Println("cannot acquire a database connection", err)
		return nil, err
	}
	defer conn.Release()

	// Query to search users with limit and mask
	query := `SELECT Id, Name, Lastname, Password, CreationDate FROM Users WHERE Name LIKE $1 AND Lastname LIKE $2 LIMIT $3`
	rows, err := conn.Query(ctx, query, nameMask, lastnameMask, limit)
	if err != nil {
		log.Println("cannot search users", err)
		return nil, err
	}
	defer rows.Close()

	// Initialize a slice to store user objects
	users := make([]models.User, 0)

	// Iterate through the rows and scan each user into the slice
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Name, &user.Lastname, &user.Password, &user.CreationDate)
		if err != nil {
			log.Println("error scanning user row", err)
			return nil, err
		}
		users = append(users, user)
	}

	// Check for errors during rows iteration
	if err = rows.Err(); err != nil {
		log.Println("error iterating user rows", err)
		return nil, err
	}

	return users, nil
}
