package utils

import (
	"go.mongodb.org/mongo-driver/bson"
)

// used to remove go vet warinings
func E(key string, value interface{}) bson.E {
	return bson.E{Key: key, Value: value}
}
