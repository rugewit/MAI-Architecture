package services

import (
	"context"
	"errors"
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

	// Insert into Users
	query := `INSERT INTO Users 
    (Id, Name, Lastname, Password, CreationDate, BasketId)
	values($1, $2, $3, $4, $5, $6)`
	_, err = conn.Exec(ctx, query, user.Id, user.Name, user.Lastname, user.Password, user.CreationDate,
		user.BasketId)
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
	query := `UPDATE Users SET Name = $1, Lastname = $2, Password = $3, CreationDate = $4, BasketId = $5 WHERE Id = $6`
	_, err = conn.Exec(ctx, query, newUser.Name, newUser.Lastname, newUser.Password, newUser.CreationDate,
		newUser.BasketId, id)
	if err != nil {
		log.Println("cannot update user", err)
		return err
	}
	return nil
}

func (service UserService) GetUserById(id string, ctx context.Context) (*models.User, error) {
	conn, err := service.Db.Pool.Acquire(ctx)
	if err != nil {
		log.Println("cannot acquire a database connection", err)
		return nil, err
	}
	defer conn.Release()

	// Query to get user by ID
	query := `SELECT Id, Name, Lastname, Password, CreationDate, BasketId FROM Users WHERE Id = $1`
	row := conn.QueryRow(ctx, query, id)

	// Initialize a user object to store the result
	user := new(models.User)

	// Scan the row into user object
	err = row.Scan(user.Id, user.Name, user.Lastname, user.Password, user.CreationDate, user.BasketId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// Return nil if no user found with the given ID
			return nil, nil
		}
		log.Println("cannot get user by ID", err)
		return nil, err
	}

	return user, nil
}

func (service UserService) GetManyUsers(limit int, ctx context.Context) ([]models.User, error) {
	conn, err := service.Db.Pool.Acquire(ctx)
	if err != nil {
		log.Println("cannot acquire a database connection", err)
		return nil, err
	}
	defer conn.Release()

	// Query to get many users with limit
	query := `SELECT Id, Name, Lastname, Password, CreationDate, BasketId FROM Users LIMIT $1`
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
		err := rows.Scan(&user.Id, &user.Name, &user.Lastname, &user.Password, &user.CreationDate, &user.BasketId)
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
