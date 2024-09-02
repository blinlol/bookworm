package dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/blinlol/bookworm/model"
	"github.com/blinlol/bookworm/utils"
)

var collectionName string = "Books"

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
		bson.D{utils.E("_id", id)},
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
		bson.D{utils.E("title", title), utils.E("author", author)},
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
			Id:     "book-" + primitive.NewObjectID().Hex(),
			Title:  b.Title,
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
		bson.D{utils.E("_id", id)},
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
		bson.D{utils.E("$set", bson.D{
			utils.E("title", b.Title),
			utils.E("author", b.Author),
			utils.E("quotes", b.Quotes)},
		)},
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
		bson.D{utils.E("$set", bson.D{utils.E("quotes", book.Quotes)})},
	)
	if err != nil {
		DBLogger.Sugar().Error(err)
	}
}

func getCollection() *mongo.Collection {
	return DBClient.Database(DBName).Collection(collectionName)
}
