package utils

import (
	"strings"

	"github.com/blinlol/bookworm/model"
	"go.mongodb.org/mongo-driver/bson"
)

// used to remove go vet warinings
func E(key string, value interface{}) bson.E {
	return bson.E{Key: key, Value: value}
}

func ParseQuotes(input string, sep string) []*model.Quote {
	quotes := make([]*model.Quote, 0)
	for _, line := range strings.Split(input, sep) {
		line = strings.Trim(line, " \n\t")
		quotes = append(quotes, &model.Quote{Text: line})
	}
	return quotes
}
