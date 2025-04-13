package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	godotenv.Load()
	clientOpts := options.Client().ApplyURI(
		fmt.Sprintf("%v", os.Getenv("DB_URL")))
	client, err := mongo.Connect(clientOpts)
	if err != nil {
		log.Panic(err)
	}
	User_collection := client.Database("todocluster").Collection("User")
	User, err := User_collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Panic(err)
	}
	for User.Next(context.TODO()) {
		var post bson.M
		err := User.Decode(&post)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(post)
	}
	defer User.Close(context.TODO())
	defer client.Disconnect(context.TODO())
}
