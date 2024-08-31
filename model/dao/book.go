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

// return nil if error
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

// return nil if document not found or error
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

// return nil if document not found or error
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
		DBLogger.Sugar().Errorln(err)
		return nil
	}
	return &book
}

// return added book or nil if error
func AddBook(b *model.Book) *model.Book {
	if b := FindBook(b.Title, b.Author); b != nil {
		return b
	}
	coll := getCollection()
	res, err := coll.InsertOne(
		DBContext,
		model.Book{
			Id: "book-" + primitive.NewObjectID().Hex(),
			Title: b.Title,
			Author: b.Author,
			Quotes: make([]*model.Quote, 0),
		},
	)
	if err != nil {
		DBLogger.Sugar().Errorln(err)
		return nil
	}
	return FindBookById(res.InsertedID.(string))
}


func DeleteBook(b *model.Book) {
	DeleteBookById(b.Id)
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


func UpdateBook(b *model.Book) *model.Book {
	_, err := getCollection().UpdateByID(
		DBContext,
		b.Id,
		bson.D{{"$set", bson.D{
				{"title", b.Title},
				{"author", b.Author},
				{"quotes", b.Quotes},
			},
		}},
	)
	if err != nil {
		DBLogger.Sugar().Error(err)
		return nil
	}
	return FindBookById(b.Id)
}

// should use UpdateBook
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
