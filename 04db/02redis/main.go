package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

func main() {
	initClient()

	redisExample()
}

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123465", // no password set
		DB:       0,        // use default DB
	})

	_, err = rdb.Ping(context.TODO()).Result()
	if err != nil {
		return err
	}
	return nil
}

func redisExample() {
	err := rdb.Set(context.TODO(), "score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}

	val, err := rdb.Get(context.TODO(), "score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}
	fmt.Println("score", val)

	val2, err := rdb.Get(context.TODO(), "name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}
}
