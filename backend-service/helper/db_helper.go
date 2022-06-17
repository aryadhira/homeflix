package helper

import (
	"context"
	"homeflix/configuration"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBConnect() (*mongo.Database, error) {
	//get configuration
	config := configuration.Config{}
	conf := config.GetConfig()

	var ctx = context.Background()
	clientOptions := options.Client()
	connectionstring := "mongodb://" + conf.DatabaseServer + ":" + conf.DatabasePort
	clientOptions.ApplyURI(connectionstring)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database(conf.Database), nil
}
