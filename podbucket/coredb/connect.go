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
	Db               = "podbucket"
	Port             = 27017
	SelectionTimeout = 8
)

var instance *mongo.Client

func connect() *mongo.Client {
	dbURI := strings.Join([]string{"mongodb://", Host, ":", strconv.Itoa(Port), "/?serverSelectionTimeoutMS=", strconv.Itoa(SelectionTimeout * 1000)}, "")
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

	clioutput.Success("Connected to \"" + Db + "\" database")

	return client
}

// Client CoreDB instance
func Client() *mongo.Client {
	if instance == nil {
		instance = connect()
	} else {
		clioutput.Info("Reusing database \"" + Db + "\" connection instance")
	}
	return instance
}
