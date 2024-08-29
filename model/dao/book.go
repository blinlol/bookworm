package dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/blinlol/bookworm/model"
)




var collectionName string = "Books"


func getCollection() * mongo.Collection{
	return DBClient.Database(DBName).Collection(collectionName)
}


func AllBooks() []*model.Book {
	coll := getCollection()
	cursor, err := coll.Find(DBContext, bson.D{})
	if err != nil {
		DBLogger.Sugar().Errorln("error while find all books:", err)
		return nil
	}
	var books []*model.Book
	err = cursor.All(DBContext, &books)
	if err != nil {
		DBLogger.Sugar().Errorln(err)
		return nil
	}
	return books
}


func FindBook(title string, author string) *model.Book {
	coll := getCollection()
	res := coll.FindOne(
		DBContext,
		bson.D{{"title", title}, {"author", author}},
	)

	var book model.Book
	err := res.Decode(&book)
	if err == mongo.ErrNoDocuments {
		return nil
	} else if err != nil {
		DBLogger.Sugar().Fatalln(err)		
	}
	return &book
}


func FindBookById(id string) *model.Book {
	res := getCollection().FindOne(
		DBContext,
		bson.D{{"_id", id}},
	)
	var b model.Book
	err := res.Decode(&b)
	if err == mongo.ErrNoDocuments {
		return nil
	} else if err != nil {
		DBLogger.Sugar().Errorln(err)
		return nil
	}
	return &b
}


func AddBook(title string, author string) *model.Book {
	if b := FindBook(title, author); b != nil {
		return b
	}
	coll := getCollection()
	res, err := coll.InsertOne(
		DBContext,
		model.Book{
			Id: primitive.NewObjectID().Hex(),
			Title: title,
			Author: author,
			Quotes: make([]*model.Quote, 0),
		},
	)
	if err != nil {
		DBLogger.Sugar().Errorln(err)
		return nil
	}
	// id := res.InsertedID.(primitive.ObjectID)
	return FindBookById(res.InsertedID.(string))
}


func DeleteBook(title string, author string) {
	coll := getCollection()
	_, err := coll.DeleteOne(DBContext, bson.D{{"title", title}, {"author", author}})
	if err != nil {
		DBLogger.Sugar().Fatalln(err)
	}
}


func DeleteBookById(id string) {
	res := getCollection().FindOneAndDelete(
		DBContext,
		bson.D{{"_id", id}},
	)
	if res.Err() == mongo.ErrNoDocuments {
		DBLogger.Sugar().Warnln("document id=", id, " not found, so not deleted")
	} else if res.Err() != nil {
		DBLogger.Sugar().Errorln(res.Err())
	}
}


func AddQuote(book *model.Book, quote *model.Quote) {
	book.Quotes = append(book.Quotes, quote)
	_, err := getCollection().UpdateByID(
		DBContext,
		book.Id,
		bson.D{{"$set", bson.D{{"quotes", book.Quotes}}}},
	)
	if err != nil {
		DBLogger.Sugar().Error(err)
	}
}


