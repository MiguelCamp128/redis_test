package main

import (
	//context es para mover información entre una cadena de llamadas, lo requiero redis desde
	// la v8
	//NO RECOMENDABLE

	"crypto/tls"
	"log"
)

func main() {



	
	
		cert, err := tls.LoadX509KeyPair("/home/jenkins/.ssh/redis.crt", "/home/jenkins/.ssh/redis.key")
		if err != nil {
			log.Fatal(err)
		}
		cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
		listener, err := tls.Listen("10.10.50.116", ":6379", cfg)
		if err != nil {
			log.Fatal(err)
		}
		_ = listener
	}
		
// 	redisHost := os.Getenv("10.10.50.116") // e.g. "1.2.3.4", "127.0.0.1", "localhost", "redis.acmecorp.com"

// 	cert, err := tls.LoadX509KeyPair("/home/jenkins/.ssh/redis.crt", "/home/jenkins/.ssh/redis.key")
// 	if err != nil {
// 		log.Fatal("Certs", err, cert)
// 	}

// 	caCert, err := ioutil.ReadFile("/home/jenkins/.ssh/CA.crt")
// 	if err != nil {
// 		log.Fatal("CertsCA2", err, caCert)
// 	}
// 	pool := x509.NewCertPool()
// 	pool.AppendCertsFromPEM(caCert)

// 	ctx := context.Background()


// 	client := redis.NewClient(&redis.Options{
// 		Addr: net.JoinHostPort(redisHost, "6379"),
// 		Password: "redis-password", 
// 		TLSConfig: &tls.Config{
// 			MinVersion: tls.VersionTLS12,
// 			//InsecureSkipVerify: true,
// 			ServerName:   "redis-server",
// 			Certificates: []tls.Certificate{cert},
// 			RootCAs:      pool,
// 		},
		
// 	})
// //nada
// 	//duration := time.Second
	
			
// 	// err = client.Set(ctx, "key", "key-content", 0).Err()
// 	// if err != nil {
// 	// 	time.Sleep(duration)
// 	// 	log.Fatal("SET", err)
// 	// }
	
// 	val, err := client.Get(ctx, "key").Result()
// 	if err != nil {
// 		log.Fatal("GET", err)
// 	}
// 	fmt.Println("key", val)

// 	val2, err := client.Get(ctx, "key2").Result()
// 	if err == redis.Nil {
// 		fmt.Println("key2 does not exist")
// 	} else if err != nil {
// 		panic(err)
// 	} else {
// 		fmt.Println("key2", val2)
// 	}
// 	// Output: key value
// 	// key2 does not exist
//}