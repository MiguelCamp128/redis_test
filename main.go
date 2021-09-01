package main

import (
	//context es para mover información entre una cadena de llamadas, lo requiero redis desde
	// la v8
	//NO RECOMENDABLE
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()
 
	rdb := redis.NewClient(&redis.Options{
		Addr:	  "10.10.50.116:6379",
		Password: "redis-password", // no password set
		DB:		  0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "cambio", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "prueba").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}