package httphelper

import (
	"fmt"      // used to print  things
	"net/http" //http
	"strings"

	"lev/datamanage"

	"github.com/go-redis/redis"
)

const (
	DB_USER  = "user:"
	DB_TOKEN = "token:"
)

// TokenCheck - is the function, that retrieves Token, then
// checks its existense in DB and returns the result (found or not)
// example: curl -kv -H "Authorization: Bearer {UUID}" -X GET http://127.0.0.1:8080/userinfo
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
