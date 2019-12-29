package coredb

import (
	"context"
	"strconv"
	"strings"

	"github.com/bounteous/podbucket-backend/podbucket/clioutput"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Host             = "localhost"
	Port             = 27017
	SelectionTimeout = 8
)

var instance *mongo.Client
var dbDomain string

func connect() *mongo.Client {
	dbDomain = strings.Join([]string{"mongodb://", Host, ":", strconv.Itoa(Port)}, "")
	// TODO use GO milisecons type
	dbURI := strings.Join([]string{dbDomain, "/?serverSelectionTimeoutMS=", strconv.Itoa(SelectionTimeout * 1000)}, "")
	clioutput.Info("Trying to connect to " + dbURI)

	// Set client options
	clientOptions := options.Client().ApplyURI(dbURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		clioutput.Error(err.Error())
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		clioutput.Fatal(err)
	}

	clioutput.Success("Connected to \"" + dbDomain + "\"")

	return client
}

// Client CoreDB instance
func Client() *mongo.Client {
	if instance == nil {
		instance = connect()
	} else {
		clioutput.Info("Reusing connection instance with " + dbDomain)
	}
	return instance
}
