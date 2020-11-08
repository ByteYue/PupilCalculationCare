package main

import (
	"context"
	"log"
	"my-app/db"
	"my-app/web"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.TODO(), clientOptions())
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())
	mongoDB := db.NewMongo(client)
	// CORS is enabled only in prod profile
	cors := os.Getenv("profile") == "prod"
	app := web.NewApp(mongoDB, cors)
	err = app.Serve()
	log.Println("Error", err)
}

//生产环境用
//func clientOptions() *options.ClientOptions {
//	host := "db"
//	if os.Getenv("profile") != "prod" {
//		host = "localhost"
//	}
//	return options.Client().ApplyURI(
//		"mongodb://" + host + ":27017",
//	)
//}

//开发环境使用
func clientOptions() *options.ClientOptions {
	//host := "db"
	//if os.Getenv("profile") != "prod" {
	//	host = "localhost"
	//}
	return options.Client().ApplyURI(
		"mongodb://" + "127.17.0.1" + ":27017",
	)
	//mac:"mongodb://"+"host.docker.internal"+":27017"
}
