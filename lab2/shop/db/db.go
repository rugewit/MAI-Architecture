package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"shop/config"
)

type PostgresDb struct {
	Post     string
	User     string
	Password string
	DbName   string
	DbUrl    string
	Pool     *pgxpool.Pool
	ctx      context.Context
}

func NewPostgresDb(ctx context.Context, pgConfig *config.PostgresConfig) (*PostgresDb, error) {
	dbPort := pgConfig.Port
	dbUser := pgConfig.User
	dbPassword := pgConfig.Password
	dbName := pgConfig.Name

	dbUrl := fmt.Sprintf("host=0.0.0.0 port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbPort, dbUser, dbPassword, dbName)

	pool, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		log.Fatalln("cannot connect to postgres sql", err)
		return nil, err
	}
	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalln("cannot ping postgres sql", err)
		return nil, err
	}
	return &PostgresDb{dbPort, dbUser, dbPassword, dbName,
		dbUrl, pool, ctx}, nil
}

func (db *PostgresDb) Close(ctx context.Context) {
	db.Pool.Close()
}
