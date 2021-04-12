package db

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"ws/config"
)

var Db *gorm.DB
var Redis *redis.Client

func Setup() {
	dns := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Mysql.Username,
		config.Mysql.Password,
		config.Mysql.Host,
		config.Mysql.Port,
		config.Mysql.Name,
	)
	db, err := gorm.Open(mysql.Open(dns),
		&gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	Db = db
	Redis = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Auth, // no password set
		DB:       0,  // use default DB
	})
}