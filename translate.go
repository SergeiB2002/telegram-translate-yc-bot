package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var API_KEY = os.Getenv("API_KEY")

func sendRequestTranslate(method, urlString string, jsonData []byte) ([]byte, error) {
	reqst, err := http.NewRequest(method, urlString, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("ошибка (из sendRequestTranslate, во время создания запроса):", err.Error())
	}

	reqst.Header.Set("Authorization", "Api-Key "+API_KEY)
	reqst.Header.Set("Content-Type", "application/json")

	respns, err := http.DefaultClient.Do(reqst)
	if err != nil {
		fmt.Println("ошибка (из sendRequestTranslate, после выполнения запроса):", err.Error())
	}
	fmt.Println("код ответа", respns.StatusCode)

	defer respns.Body.Close()
	respnsBody, err := io.ReadAll(respns.Body)

	return respnsBody, nil
}

func detectLanguage(userText string) (string, bool) {
	var jsonData = []byte(fmt.Sprintf(`{
        "text": "%s"
        }`, userText))

	respnsBody, err := sendRequestTranslate("POST", DETECT_LANGUAGE_URL, jsonData)
	if err != nil {
		fmt.Println(err.Error())
	}

	var respns DetectLanguageResponse

	err = json.Unmarshal(respnsBody, &respns)
	if err != nil {
		fmt.Println(err.Error())
	}

	language, notSupport := defineLanguageByLanguageCode(respns.LanguageCode)

	return language, notSupport
}

func defineLanguageByLanguageCode(languageCode string) (string, bool) {

	language := ""
	flag := false

	switch languageCode {
	case "ru":
		language = "Русский"
	case "en":
		language = "Английский"
	case "es":
		language = "Испанский"
	case "ar":
		language = "Арабский"
	case "zh":
		language = "Китайский"
	case "fr":
		language = "Французский"
	case "de":
		language = "Немецкий"
	case "pt":
		language = "Португальский"
	case "hi":
		language = "Хинди"
	case "ja":
		language = "Японский"
	default:
		flag = true
		sendMsg("Простите, на данный момент бот не может распознать данный язык.\n" +
			">>Языки, которые на данный момент может распознать бот при переводе: русский, английский, немецкий, французский, испанский, португальский, арабский, китайский, японский и хинди.")
	}

	return language, flag
}

func translateText(targetLanguageCode, userText string) TranslateTextResponse {
	var jsonData = []byte(fmt.Sprintf(`{
        "targetLanguageCode": "%s",
        "texts": [
            "%s"
            ]
        }`, targetLanguageCode, userText))

	respnsBody, err := sendRequestTranslate("POST", TRANSLATE_TEXT_URL, jsonData)
	if err != nil {
		fmt.Println(err.Error())
	}

	var respns TranslateTextResponse

	err = json.Unmarshal(respnsBody, &respns)
	if err != nil {
		fmt.Println(err.Error())
	}

	return respns
}
