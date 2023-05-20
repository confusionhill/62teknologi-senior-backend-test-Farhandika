package application

import (
	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
	_ "github.com/simukti/sqldb-logger"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"project-template/internal/config"
	"project-template/model/entity"
)

type CommonResource struct {
	Database *gorm.DB
	Redis    *redis.Client
	Mongo    *mongo.Client
}

func InitResource(cfg *config.MainConfig) (*CommonResource, error) {
	r := new(CommonResource)

	godb, err := gorm.Open(postgres.Open(cfg.Database.Host), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	godb.AutoMigrate(&entity.Business{})
	r.Database = godb
	log.Println("DATABASE INITIATED!")

	rdc := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	r.Redis = rdc
	log.Println("REDIS INITIATED!")

	log.Println("RESOURCE INITIATED")
	return r, nil
}
