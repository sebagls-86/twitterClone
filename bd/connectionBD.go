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

	errorLoad := godotenv.Load()
	if errorLoad != nil {
		panic(error1)
	}

	mongodb := os.Getenv("MONGODB_URI")

	var clientOptions = options.Client().ApplyURI(mongodb)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexion exitosa")

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
