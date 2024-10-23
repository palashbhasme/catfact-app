package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	client *mongo.Client
}

func NewServer(client *mongo.Client) *Server {
	return &Server{
		client: client,
	}
}

func (server *Server) handleAllFacts(writer http.ResponseWriter, request *http.Request) {

	coll := server.client.Database("catfact").Collection("facts")
	cursor, err := coll.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Fatalln("Error in Query, handleAllFacts:", err)
	}

	defer cursor.Close(context.TODO())

	results := []bson.M{}
	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Fatalln(err)
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(results); err != nil {
		log.Fatalln("Error in Query, handleAllFacts:", err)
	}

}

type CatFactWorker struct {
	client *mongo.Client
}

func NewWorker(client *mongo.Client) *CatFactWorker {
	return &CatFactWorker{
		client: client,
	}
}

func (cfw *CatFactWorker) getFacts() error {
	coll := cfw.client.Database("catfact").Collection("facts")

	ticker := time.NewTicker(3 * time.Second)

	for {
		resp, err := http.Get("http://catfact.ninja/fact")
		if err != nil {
			return err
		}

		var fact bson.M
		if err = json.NewDecoder(resp.Body).Decode(&fact); err != nil {
			return err
		}

		if _, err := coll.InsertOne(context.TODO(), fact); err != nil {
			return err
		}

		<-ticker.C
	}
}
func main() {
	options := options.Client().ApplyURI("mongodb://localhost:27017/?ssl=false")
	client, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		panic(err)
	}

	worker := NewWorker(client)
	go worker.getFacts()

	port := 3000
	addr := fmt.Sprintf(":%d", port)

	server := NewServer(client)
	http.HandleFunc("/givemefacts", server.handleAllFacts)
	fmt.Printf("Sever listning on port %d", port)
	http.ListenAndServe(addr, nil)

}
