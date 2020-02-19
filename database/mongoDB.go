package database

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/x/network/connstring"
	"log"
)

type model interface {
	InitModel(d *mongo.Database)
}

var (
	dbName = "go-auth"
	models []model
	client *mongo.Client
	db *mongo.Database
)

func Connect(uri string)  {
	// parse uri to get name of database
	cs, err := connstring.Parse(uri)
	if err != nil {
		log.Fatal(err)
	}

	if cs.Database != "" {
		dbName = cs.Database
	}

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), uri)
	if err != nil { log.Fatal(err) }

	// test connection
	err = client.Ping(context.TODO(), nil)
	if err != nil { log.Fatal(err) }

	// get database
	db = client.Database(dbName)

	fmt.Println("Connected to MongoDB:", uri)
}

func Disconnect()  {
	err := client.Disconnect(context.TODO())
	if err != nil { log.Fatal(err) }

	fmt.Println("Connection to MongoDB closed")
}

func InitModels()  {
	for i := 0; i < len(models); i++ {
		models[i].InitModel(db)
	}

	// clean array of models
	models = []model{}
}

func Model(m model)  {
	models = append(models, m)
}
