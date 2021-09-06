package main

import (
	//context es para mover información entre una cadena de llamadas, lo requiero redis desde
	// la v8
	//NO RECOMENDABLE

	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)
func main() {

cert, err := tls.LoadX509KeyPair("/home/jenkins/.ssh/redis.crt", "/home/jenkins/.ssh/redis.key")
if err != nil {
	log.Fatal(err)
}
caCert, err := ioutil.ReadFile("/home/jenkins/.ssh/CA.crt")
if err != nil {
	log.Fatal(err)
}
caCertPool := x509.NewCertPool()
caCertPool.AppendCertsFromPEM(caCert)

	client:=redis.NewClient(&redis.Options{
		Addr:     "10.10.50.116:6379",
		Password: "redis-server",
		DB:       0,
		TLSConfig: &tls.Config{
			RootCAs: caCertPool,
			Certificates: []tls.Certificate{
				cert,
			},
		},
	})
	fmt.Println(client)


	ctx := context.Background()



// //nada
duration := time.Second
	
			
	err = client.Set(ctx, "key", "key-content", 0).Err()
	if err != nil {
		time.Sleep(duration)
		log.Fatal("SET", err)
	}
	
	val, err := client.Get(ctx, "key").Result()
	if err != nil {
		log.Fatal("GET", err)
	}
	fmt.Println("key", val)

	val2, err := client.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

	// keys,err :=redis.Strings(cn.Do("KEYS", "*"))
	// if err != nil {
	// 	log.Fatal( err)
	// }
	// for _, key := range keys{
	// 	fmt.Println(key)
	// }
	// Output: key value
	// key2 does not exist
}