package storagemanage

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http" //http
	"os"

	"lev/datamanage"
	"lev/httphelper"

	"github.com/google/uuid"
	// "github.com/go-redis/redis"
)

type HttpResponse struct {
	Status string `json:"status"`
}

const (
	DB_USER  = "user:"
	DB_TOKEN = "token:"
)

func handlerUpload(rw http.ResponseWriter, req *http.Request) {
	var (
		// usr datamanage.RecCheck
		obj     datamanage.HRecCheck
		fileRec datamanage.HRecord
		// res1    string
	// client *redis.Client
	//resp   HttpResponse
	)
	// parse Token and check in Redis using function
	if httphelper.TokenCheck(req) == "user exists" {
		fmt.Printf("token found  -1\n") //debug

		// parse file name and upload it
		// the FormFile function takes in the POST input id file
		file, header, err := req.FormFile("file")
		fmt.Printf("doexali: file---2\n") //debug
		if err != nil {
			panic(err)
			// fmt.Fprintf(rw, err)
			// return
		}

		defer file.Close()

		// check if target folder /tmp/{User-UUID} exists
		// and create it if doesn't exists
		path := "/tmp/" + httphelper.ParseToken(req)
		fmt.Println("path =  ", path) // debug: UUID folder exists
		// if _, err1 := os.Stat(path); os.IsNotExist(err1) {
		// 	os.Mkdir(path, 755)
		_, err1 := os.Stat(path)
		if os.IsNotExist(err1) {
			os.Mkdir(path, 755)
			fmt.Println("created new dir") // debug - folder not found ad created
		} else {
			fmt.Println("forder already exists")

		}
		// debug - list dir
		files, err := ioutil.ReadDir("/tmp/")
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
			fmt.Println("kuku", f.Name())
		}

		//check if the file is already loaded - its fileName will appear in HGET
		// for User
		client := datamanage.InitRedis()
		fmt.Println(header.Filename)             //debug
		obj.KeyName = httphelper.ParseToken(req) // to ask DB
		obj.FieldName = header.Filename          // to ask DB
		fmt.Printf("HRecCheck: [%#v]\n", obj)
		res := datamanage.HExistsRedis(client, obj)
		fmt.Printf("res: [%s]\n", res)

		if res == "key-field not found" {

			//generate file:UUID, load file and
			uuidStr := uuid.New().String()
			fileRec.KeyName = httphelper.ParseToken(req)
			fileRec.FieldName = header.Filename
			fileRec.ValueName = uuidStr

			// create target file on the docker dir root/tmp/file-UUID
			out, err := os.Create(path + "/" + uuidStr)
			if err != nil {
				fmt.Fprintf(rw, "Unable to create the file for writing. Check your write access privilege")
				return
			}

			defer out.Close()

			// create a record HSET in redis
			res1 := datamanage.HSetRedis(client, fileRec)

			if res1 == "HSET done" {
				// write the content from POST to the file
				_, err = io.Copy(out, file)
				if err != nil {
					fmt.Fprintln(rw, err)
				}

				fmt.Fprintf(rw, "File uploaded successfully : \n")
				fmt.Fprintf(rw, header.Filename)
				fmt.Fprintf(rw, "\n")

				json.NewEncoder(rw).Encode(HttpResponse{Status: "file loaded"})
			}
		}

		if res == "key-field exists" {
			json.NewEncoder(rw).Encode(HttpResponse{Status: "file already loaded"})
		}
	} else {
		json.NewEncoder(rw).Encode(HttpResponse{Status: "authorization failed"})
	}
}

func handlerStorage(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		handlerUpload(rw, req)
	}
	// case "GET":
	// 	handlerInfoGet(rw, req)
}

func Start() {
	// http.HandleFunc("/upload", handlerUpload)
	// http.HandleFunc("/auth", handlerAuth)
	http.HandleFunc("/storage", handlerStorage)
	//	log.Fatal(http.ListenAndServe(":8082", nil))
	http.ListenAndServe(":8082", nil)
}
