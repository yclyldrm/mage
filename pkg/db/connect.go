package db

import (
	"context"
	"fmt"
	"log"
	"mage/pkg/configs"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBHandler struct {
	MongoDB     *mongo.Database
	RedisClient *redis.Client
}

var DbHand DBHandler

func ConnectDatabases() {
	var err error
	DbHand.MongoDB, err = connectDB()

	if err != nil {
		log.Fatalf("Error occured while connecting MongoDB, message: %v", err)
	}

	DbHand.RedisClient, err = connectRedis()
	if err != nil {
		log.Fatalf("Error occured while connecting Redis, message: %v", err)
	}

	log.Println("Databases connected")
}

func connectDB() (*mongo.Database, error) {
	dbURI := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.ovabaku.mongodb.net/?retryWrites=true&w=majority", configs.Cfg.DBUser, configs.Cfg.DBPassword)

	clientOptions := options.Client().ApplyURI(dbURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connected successfully!")
	return client.Database("mage"), err
}

func connectRedis() (*redis.Client, error) {
	opt, _ := redis.ParseURL(strings.Replace(configs.Cfg.RedisURL, "password", configs.Cfg.RedisPassword, 1))
	client := redis.NewClient(opt)
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Redis connected successfully!")
	return client, err
}
