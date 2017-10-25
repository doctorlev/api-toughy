package storagemanage

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http" //http
	"os"

	"lev/httphelper"
	// "github.com/go-redis/redis"
)

type HttpResponse struct {
	Status string `json:"status"`
}

const (
	// DB_USER  = "user:"
	DB_TOKEN = "token:"
)

func handlerUpload(rw http.ResponseWriter, req *http.Request) {
	var (
	// usr    datamanage.RecCheck
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

		// create target file on the docker dir root/tmp/uploadfile/
		out, err := os.Create("/tmp/uploadedfile")
		if err != nil {
			fmt.Fprintf(rw, "Unable to create the file for writing. Check your write access privilege")
			return
		}

		defer out.Close()

		// write the content from POST to the file
		_, err = io.Copy(out, file)
		if err != nil {
			fmt.Fprintln(rw, err)
		}

		fmt.Fprintf(rw, "File uploaded successfully : \n")
		fmt.Fprintf(rw, header.Filename)
		fmt.Fprintf(rw, "\n")

		//**************

		json.NewEncoder(rw).Encode(HttpResponse{Status: "file loaded"})
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
