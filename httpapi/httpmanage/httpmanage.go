package httpmanage

import (
	"encoding/json" //manipulations with JSON
	"fmt"           // used to print  things
	"io/ioutil"
	"net/http" //http

	"lev/datamanage"
	"lev/httphelper"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

type AuthResponse struct {
	Token string `json:"token"`
}

type HttpResponse struct {
	Status string `json:"status"`
}

const (
	DB_USER  = "user:"
	DB_TOKEN = "token:"
)

func handlerUsersPost(rw http.ResponseWriter, req *http.Request) {
	var (
		pair   datamanage.AuthRequest
		client *redis.Client
		res    HttpResponse
	)
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &pair)
	if err != nil {
		panic(err)
	}

	client = datamanage.InitRedis()
	pair.Username = DB_USER + pair.Username
	// res = writeRedis(client, pair)
	res = HttpResponse{Status: datamanage.WriteRedis(client, pair)}
	json.NewEncoder(rw).Encode(res) // this is THE Marshalling
}

func handlerUsersGet(rw http.ResponseWriter, req *http.Request) {
	var (
		usr    datamanage.RecCheck
		client *redis.Client
		resp   HttpResponse
	)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &usr)
	if err != nil {
		panic(err)
	}

	client = datamanage.InitRedis()
	usr.KeyName = DB_USER + usr.KeyName
	resp = HttpResponse{Status: datamanage.ExistsRedis(client, usr)} // resp.Status = readRedis(...)
	json.NewEncoder(rw).Encode(resp)                                 // this is THE Marshalling

}

func handlerUsers(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		handlerUsersPost(rw, req)
	case "GET":
		handlerUsersGet(rw, req)
	}
}

func handlerAuthPost(rw http.ResponseWriter, req *http.Request) {
	var (
		pair, pairUuid datamanage.AuthRequest
		client         *redis.Client
		res            datamanage.RecCheck
		res1, res2     string
	)
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &pair)
	if err != nil {
		panic(err)
	}

	client = datamanage.InitRedis()
	res.KeyName = DB_USER + pair.Username
	fmt.Println(res) //debug
	res1 = datamanage.ReadRedis(client, res)

	if res1 == pair.Password {
		fmt.Println("password matches") //debug

		//generate uuid
		uuidStr := uuid.New().String()

		pairUuid.Username = DB_TOKEN + uuidStr // set new pair for redis
		pairUuid.Password = pair.Username
		// fmt.Println(pairUuid.Username, pairUuid.Password) //debug

		res2 = datamanage.WriteRedis(client, pairUuid)

		if res2 == "successful" {
			json.NewEncoder(rw).Encode(AuthResponse{Token: uuidStr})
		} else {
			json.NewEncoder(rw).Encode(HttpResponse{Status: "user not authorized"})
			// p := HttpResponse{Status: "wrong user"}
			json.NewEncoder(rw).Encode(HttpResponse{Status: "wrong user"})
		}
	}
}

func handlerAuth(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		handlerAuthPost(rw, req)
		// case "GET":
		// 	handlerUsersGet(rw, req) //not in use yet
	}
}

func handlerInfoGet(rw http.ResponseWriter, req *http.Request) {

	reqToken := httphelper.TokenCheck(req)
	if reqToken == "user exists" {
		secResp := "secret"
		json.NewEncoder(rw).Encode(HttpResponse{Status: secResp})
	} else {
		json.NewEncoder(rw).Encode(HttpResponse{Status: "authorization failed"})
	}
}

func handlerInfo(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		handlerInfoGet(rw, req)
	}
}

func Start() {
	http.HandleFunc("/users", handlerUsers)
	http.HandleFunc("/auth", handlerAuth)
	http.HandleFunc("/userinfo", handlerInfo)
	http.ListenAndServe(":8082", nil)
	//	log.Fatal(http.ListenAndServe(":8082", nil))
}
