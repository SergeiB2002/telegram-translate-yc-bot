package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var BOT_TOKEN = os.Getenv("BOT_TOKEN")
var chatId int
var respnsMap map[string]string
var user *User

func sendMsg(msg string) {

	chatIdStr := strconv.Itoa(chatId)
	text := url.QueryEscape(msg)
	fmt.Println(TELEGRAM_API_TEMPLATE + "/sendMessage?chat_id=" + chatIdStr + "&text=" + text)
	reqst, err := http.NewRequest("POST", TELEGRAM_API_TEMPLATE+BOT_TOKEN+"/sendMessage?chat_id="+chatIdStr+
		"&text="+text, nil)
	if err != nil {
		fmt.Println("from sendMsg", err.Error())
	}
	_, err = http.DefaultClient.Do(reqst)
	if err != nil {
		fmt.Println("from sendMsg2", err.Error())
	}
}

func printHello() {
	msg := "Привет! Это бот для работы с Яндекс Переводчиком!"
	sendMsg(msg)
}

func updateProc(update *TelegramMessageResponse) {

	if update.Message.Text == "/start" {
		printHello()
		sendMsg(HELP)
		return
	}

	if update.Message.Text == "/help" {
		sendMsg(HELP)
		return
	}

	splitMsg := strings.Split(update.Message.Text, " | ")
	switch splitMsg[0] {

	case "/определить":
		if len(splitMsg) != 2 {
			sendMsg("Некорректный ввод входных данных. Введите команду по следующему шаблону:\n\n" +
				"`/определить | <текст>`")
			return
		}
		language, notSupport := detectLanguage(splitMsg[1])
		if notSupport {
			return
		}
		sendMsg("Язык введённого текста: " + language)
		return

	case "/перевести":
		if len(splitMsg) != 3 {
			sendMsg("Некорректный ввод входных данных. Введите команду по следующему шаблону:\n\n" +
				"`/перевести | <код языка> | <текст>`")
			return
		}

		respns := translateText(splitMsg[1], splitMsg[2])
		language, notSupport := defineLanguageByLanguageCode(respns.Translations[0].DetectedLanguageCode)
		if notSupport {
			sendMsg("\n\nПеревод введённого текста:\n\n" + respns.Translations[0].Text)
			return
		}
		sendMsg("\n\nПеревод введённого текста:\n\n" + respns.Translations[0].Text + "\n\nЯзык введённого текста: " + language)
		return

	default:
		sendMsg("Неизвестная команда.")
		sendMsg(HELP)
	}

}

type Response struct {
	StatusCode int         `json:"statusCode"`
	Body       interface{} `json:"body"`
}

func Main(ctx context.Context, reqst []byte) (*Response, error) {

	var respns TelegramMessageResponse

	err := json.Unmarshal(reqst, &respnsMap)
	if err != nil {
		fmt.Println("ошибка", err.Error())
	}

	fmt.Println("мап", respnsMap)
	fmt.Println("тело мапа", respnsMap["body"])

	err = json.Unmarshal([]byte(respnsMap["body"]), &respns)
	if err != nil {
		fmt.Println("ошибка при распарсивании строки:", err)
	}

	fmt.Println()
	fmt.Println("структура", respns)

	chatId = respns.Message.Chat.Id

	user = &User{Id: respns.Message.From.Id, Name: respns.Message.From.Username, Token: "", ChatId: chatId}

	updateProc(&respns)

	return &Response{
		StatusCode: 200,
		Body:       "успешно",
	}, nil

}
