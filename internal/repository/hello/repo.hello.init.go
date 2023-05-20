package hello

import (
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type HelloRepository struct {
	Database *gorm.DB
	Redis    *redis.Client
	Mongo    *mongo.Client
}

func Init(db *gorm.DB, rd *redis.Client, mg *mongo.Client) (*HelloRepository, error) {
	r := &HelloRepository{Database: db, Redis: rd, Mongo: mg}
	return r, nil
}
