package httphelper

import (
	"fmt"      // used to print  things
	"net/http" //http
	"strings"

	"lev/datamanage"

	"github.com/go-redis/redis"
)

// type AuthResponse struct {
// 	Token string `json:"token"`
// }

// type HttpResponse struct {
// 	Status string `json:"status"`
// }

const (
	DB_USER  = "user:"
	DB_TOKEN = "token:"
)

// The function retrieves Token, checks it in DB and returns result (found or not)
// example: curl -kv -H "Authorization: Bearer 3ee685b7-9466-4b19-83be-8727f1f44af1" -X GET http://127.0.0.1:8080/userinfo
func TokenCheck(req *http.Request) string {
	var (
		usr    datamanage.RecCheck
		client *redis.Client
	)
	// parse Token
	reqToken := req.Header.Get("Authorization")
	fmt.Printf("helper debug: [%#v]\n", reqToken) //debug
	splitToken := strings.Split(reqToken, "Bearer")
	fmt.Println(splitToken) //debug
	reqToken = splitToken[1]
	reqToken = strings.Trim(reqToken, " ")
	fmt.Println(reqToken) //debug
	//  and check if it exists in Redis
	client = datamanage.InitRedis()
	usr.KeyName = DB_TOKEN + reqToken
	fmt.Printf("RecCheck: [%#v]\n", usr)
	res := datamanage.ExistsRedis(client, usr)
	fmt.Printf("res: [%s]\n", res)

	return res
}
