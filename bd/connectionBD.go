package bd

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConnectBD()

var error1 error

func ConnectBD() *mongo.Client {

	ctx := context.TODO()

	errorLoad := godotenv.Load()
	if errorLoad != nil {
		panic(error1)
	}

	mongodb := os.Getenv("MONGODB_URI")

	var clientOptions = options.Client().ApplyURI(mongodb)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("connection established")

	return client
}

func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return 0
	}

	return 1
}
