package config

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var Cache *redis.Client

func ConnectDb() {
	const (
		user     = "dustin"
		password = "12345"
		dbname   = "cardb"
	)
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	//connecct to database using gorm
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	//db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("Error openning database")
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("unable to move to gorm")
		panic(err)
	}

	if err = sqlDB.Ping(); err != nil {
		fmt.Println("Error connecting to database")
		panic(err)
	}
	fmt.Println("Successfully connected to database")
	Db = db
}
func ConnectCache() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	cmd := rdb.Ping(ctx)
	if cmd.Err() != nil {
		fmt.Println("Error connecting to caching database")
		panic(cmd.Err())
	}

	fmt.Println("Successfully connected to caching database")

	Cache = rdb
}
