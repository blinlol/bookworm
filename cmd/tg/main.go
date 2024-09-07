package main

import (
	"math/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/blinlol/bookworm/model"
)

// get books
// choice book by keyboard and get its random quote in one message

var apiurl string = os.Getenv("API_URI_FOR_UI")


func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TGBOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	main_wg := sync.WaitGroup{}

	for u := range updates {
		update := u
		main_wg.Add(1)

		go func(){
			defer main_wg.Done()

			if update.Message != nil {
				if update.Message.IsCommand() {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

					switch update.Message.Command() {
					case "books":
						msg.Text = "books list"
						msg.ReplyMarkup = getBooksInlineKeybordMarkup()
					default:
						msg.Text = "unknown command"
					}
					if _, err := bot.Send(msg); err != nil {
						log.Printf("Error while sending message: %v", err)
					}
				}
			} else if update.CallbackQuery != nil {
				book_id := update.CallbackQuery.Data
				quote := getRandomQuote(book_id)
				callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
				if _, err = bot.Request(callback) ; err != nil {
					log.Printf("Error sending callback: %v", err)
					return
				}
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "no quotes for this book")
				if quote != nil {
					msg.Text = quote.Text
				}
				if _, err := bot.Send(msg); err != nil {
					log.Printf("Error sending message: %v", err)
				}
			}
		}()
	}
	main_wg.Wait()
}

func getBooksInlineKeybordMarkup() *tgbotapi.InlineKeyboardMarkup {
	var btn tgbotapi.InlineKeyboardButton
	keyboard := tgbotapi.NewInlineKeyboardMarkup()
	for _, book := range getBooks() {
		btn.Text = fmt.Sprintf("\"%s\" %s", book.Title, book.Author)
		btn.CallbackData = &book.Id
		keyboard.InlineKeyboard = append(
			keyboard.InlineKeyboard,
			[]tgbotapi.InlineKeyboardButton{btn},
		)
	}
	if len(keyboard.InlineKeyboard) == 0 {
		keyboard.InlineKeyboard = append(
			keyboard.InlineKeyboard,
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("books not found", "no books"),
			))
	}
	return &keyboard
}

func getBooks() []*model.Book {
	resp, err := http.Get(apiurl + "/api/books")
	if err != nil {
		log.Printf("Error get books from api: %v", err)
		return make([]*model.Book, 0)
	}
	defer resp.Body.Close()
	var resp_model model.GetBooksResponse
	err = json.NewDecoder(resp.Body).Decode(&resp_model)
	if err != nil {
		log.Printf("Error. Wrong response from api: %v", err)
		return make([]*model.Book, 0)
	}
	return resp_model.Books
}

func getRandomQuote(book_id string) *model.Quote {
	resp, err := http.Get(apiurl + "/api/quotes/" + book_id)
	if err != nil {
		log.Printf("Error fetching random quote: %v", err)
		return nil
	}
	defer resp.Body.Close()

	var resp_model model.GetQuotesResponse
	err = json.NewDecoder(resp.Body).Decode(&resp_model)
	if err != nil {
		log.Printf("Error. Wrong response from api: %v", err)
		return nil
	}
	if len(resp_model.Quotes) == 0 {
		return nil
	}

	return resp_model.Quotes[rand.Intn(len(resp_model.Quotes))]
}