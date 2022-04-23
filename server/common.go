package main

import (
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	
	"github.com/subosito/gotenv"
)

var MongoCollection *mongo.Collection

func init() {
	gotenv.Load()
}

// var user string = os.Getenv("MONGO_USER")
// var password string = os.Getenv("MONGO_PASSWORD")
// var domain string = os.Getenv("MONGO_DOMAIN")
var ConnectionString string = os.Getenv("MONGO_CONNECTION_STRING")
