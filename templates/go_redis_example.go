package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
)

func stringRedis(client *redis.Client) {
	key := "name"
	value := "大太阳aaa"
	err := client.Set(key, value, 0).Err()
	checkError(err)

	v2, err := client.Get(key).Result()
	checkError(err)
	fmt.Println(v2)

	client.Del(key)
}

func checkError(err error) {
	if err != nil {
		if err == redis.Nil {
			fmt.Println("key不存在...")
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	stringRedis(client)
}
