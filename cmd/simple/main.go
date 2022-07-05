package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/paulwizviz/learn-mongodb/internal/mongoutil"
	"go.mongodb.org/mongo-driver/bson"
)

type Person struct {
	GivenName string `bson:"given_name"`
	Surname   string `bson:"surname"`
}

func monogoOps() ([]bson.M, error) {
	uri := "mongodb://admin:admin@db:27017"

	conn, err := mongoutil.NewConnection(context.TODO(), uri)
	if err != nil {
		return nil, err
	}
	defer mongoutil.Disconnect(context.TODO(), conn)

	p := Person{
		GivenName: "John",
		Surname:   "Doe",
	}

	err = mongoutil.InsertOneDocument(context.TODO(), conn, "test", "test", p)
	if err != nil {
		return nil, err
	}

	result, err := mongoutil.FindAllDocument(context.TODO(), conn, "test", "test")
	if err != nil {
		return nil, err
	}

	return result, nil

}

func handler(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	result, err := monogoOps()
	if err != nil {
		rw.WriteHeader(4000)
		rw.Write([]byte(err.Error()))
	}

	respData, err := json.Marshal(result)
	if err != nil {
		rw.WriteHeader(4001)
		rw.Write([]byte(err.Error()))
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(respData)

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", 4040), router))
}
