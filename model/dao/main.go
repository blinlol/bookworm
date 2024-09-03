package dao

import (
	"os"

	"go.uber.org/zap"
)


var Logger *zap.Logger
var ConnString string


func init(){
	var err error
	Logger, err = zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	ConnString = os.Getenv("POSTGRES_URI")
}