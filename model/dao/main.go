package dao

import (
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

	ConnString = "postgresql://bookworm_user:123@localhost:5432/bookworm_db"
}