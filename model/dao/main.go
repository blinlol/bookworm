package dao

import (
	// 	"go.mongodb.org/mongo-driver/bson"
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var DBClient *mongo.Client
var DBContext context.Context
var DBLogger *zap.Logger

var DBName string = "bookworm_db"

func initContext() {
	DBContext = context.Background()
}

func initLogger() {
	var err error
	DBLogger, err = zap.NewDevelopment()
	if err != nil {
		log.Fatalln(err)
	}
}

func initClient() {
	var err error
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017"
		DBLogger.Sugar().Warnln("MONGO_URI doesn't set. Use uri =", uri)
	}
	DBClient, err = mongo.Connect(
		DBContext,
		options.Client().ApplyURI(uri).SetConnectTimeout(time.Second),
	)
	if err != nil {
		DBLogger.Sugar().Fatalln(err)
	}
}

func init() {
	initContext()
	initLogger()
	initClient()
}
