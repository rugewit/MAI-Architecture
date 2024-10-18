package redis_setup

import (
	"account_service/additional"
	"account_service/config"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"log"
	"strconv"
)

type RedisAccount struct {
	Client *redis.Client
	Config *config.RedisConfig
}

func NewRedisAccount() *RedisAccount {
	err := additional.LoadViper("env/.env")
	if err != nil {
		log.Fatalln("cannot load viper")
		return nil
	}

	useRedis, err := strconv.ParseBool(viper.GetString("USE_REDIS"))
	if err != nil {
		log.Fatal(err)
	}
	redisConfig := &config.RedisConfig{
		RedisAliveTimeSec: viper.GetString("REDIS_ALIVE_TIME_SEC"),
		ConnectionUri:     viper.GetString("CONNECTION_URI"),
		UseRedis:          useRedis,
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.ConnectionUri,
		Password: "",
		DB:       0,
	})

	return &RedisAccount{
		Client: client,
		Config: redisConfig,
	}
}
