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
	"net"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
		
	redisHost := os.Getenv("10.10.50.116") // e.g. "1.2.3.4", "127.0.0.1", "localhost", "redis.acmecorp.com"

	cert, err := tls.LoadX509KeyPair("/home/jenkins/.ssh/redis.crt", "/home/jenkins/.ssh/redis.key")
	if err != nil {
		panic(err)
	}

	caCert, err := ioutil.ReadFile("/home/jenkins/.ssh/CA.crt")
	if err != nil {
		panic(err)
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caCert)

	ctx := context.Background()


	client := redis.NewClient(&redis.Options{
		Addr: net.JoinHostPort(redisHost, "6379"),
		Password: "redis-password", 
		TLSConfig: &tls.Config{
			//MinVersion: tls.VersionTLS12,
			//InsecureSkipVerify: true,
			//ServerName:   "10.10.50.116",
			Certificates: []tls.Certificate{cert},
			RootCAs:      pool,
		},
		
	})
//nada
	duration := time.Second
	
			
	err = client.Set(ctx, "key", "key-content", 0).Err()
	if err != nil {
		fmt.Println("Error en el set")
		time.Sleep(duration*5)
		panic(err)
	}
	
	val, err := client.Get(ctx, "prueba").Result()
	if err != nil {
		panic(err)
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
	// Output: key value
	// key2 does not exist
}