package model

import (
	"fmt"
	"os"
	"time"

	"GOGOGO/config"
	"github.com/garyburd/redigo/redis"
	"github.com/globalsign/mgo"
	"github.com/jinzhu/gorm"
)

// DB 数据库连接
var DB *gorm.DB

// RedisPool Redis连接池
var RedisPool *redis.Pool

// MongoDB 数据库连接
var MongoDB *mgo.Database

func initDB() {
	db, err := gorm.Open(config.DBConfig.Dialect, config.DBConfig.URL)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	//if config.ServerConfig.Env == DevelopmentMode { todo 开发阶段
	db.LogMode(true)
	//}
	db.SingularTable(true) // 表名使用单数
	db.DB().SetMaxIdleConns(config.DBConfig.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.DBConfig.MaxOpenConns)
	db.DB().SetConnMaxLifetime(60 * time.Second) // fixed for 数据库连接断开. see: https://github.com/jinzhu/gorm/issues/1822  && https://github.com/go-sql-driver/mysql/issues/657
	DB = db
}

func initRedis() {
	RedisPool = &redis.Pool{
		MaxIdle:     config.RedisConfig.MaxIdle,
		MaxActive:   config.RedisConfig.MaxActive,
		IdleTimeout: 240 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.RedisConfig.URL, redis.DialPassword(config.RedisConfig.Password))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
}

/*
 * mgo文档 http://labix.org/mgo
 * https://godoc.org/gopkg.in/mgo.v2
 * https://godoc.org/gopkg.in/mgo.v2/bson
 * https://godoc.org/gopkg.in/mgo.v2/txn
 */
func initMongo() {
	if config.MongoConfig.URL == "" {
		return
	}
	session, err := mgo.Dial(config.MongoConfig.URL)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	MongoDB = session.DB(config.MongoConfig.Database)
}

func init() {
	initDB()
	initDB2()
	initRedis()
	initMongo()
}
