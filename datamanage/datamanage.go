package datamanage

import (
	"fmt"
	"strings"

	"github.com/go-redis/redis"
)

type AuthRequest struct {
	Username string
	Password string
}

type RecCheck struct {
	KeyName string
}

// initialize new redis client function as a ponter to "redis.Client" method of redis lib
func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "172.17.0.1:6379", // IP of docker
		Password: "",                // no password set
		DB:       0,                 // use default DB
	})

	pong, err := client.Ping().Result() // sanity test - to reply 'pong' on 'ping'
	fmt.Println(pong, err)              // and print  it as "PONG <nil>"

	return client
}

// WriteRedis is function to write to REDIS (the pair key-value is taken from data structure
// which values are withdrawn from HTTP Req (jsonpair)
func WriteRedis(client *redis.Client, jsonpair AuthRequest) (res string) {
	fmt.Println("writing to redis")

	err := client.Set(jsonpair.Username, jsonpair.Password, 0).Err()
	if err != nil {
		panic(err)
	}
	if strings.HasPrefix(jsonpair.Username, "token:") {
		client.Expire(jsonpair.Username, 6000000000000)
	}
	val, err := client.Get(jsonpair.Username).Result() //testing entered key-val
	if err != nil {
		fmt.Println("Failed")
		res = "couldn't find the key"
	}

	if val == jsonpair.Password {
		fmt.Println("Writing Successful")
		res = "successful"
	} else {
		res = "failed"
	}

	fmt.Println("new key and value are: ", jsonpair.Username, val)
	return res
}

// function to return the VAL by the KEY from redis
func ReadRedis(client *redis.Client, jsonval RecCheck) string {
	fmt.Println("ReadRedis: ", jsonval.KeyName) //debug

	val, err := client.Get(jsonval.KeyName).Result()
	if err != nil {
		fmt.Println("Failed")
		val = "not found"
	}

	fmt.Println("value: ", val)
	return val
}

//function to validate the existance of the key in redis
func ExistsRedis(client *redis.Client, jsonval RecCheck) string {
	fmt.Println("ExistsRedis: ", jsonval.KeyName) //debug
	var result string
	val, err := client.Exists(jsonval.KeyName).Result()
	if err != nil {
		fmt.Println("Failed")
	}
	if val == 1 {
		result = "user exists"
	}
	if val == 0 {
		result = "user not found"
	}
	return result
}
