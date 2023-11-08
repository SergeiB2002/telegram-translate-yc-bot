package main

type User struct {
	Id     int
	Name   string
	Token  string
	ChatId int
}

type TelegramMessageResponse struct {
	UpdateId int `json:"update_id"`
	Message  struct {
		MessageId int `json:"message_id"`
		From      struct {
			Id           int    `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			Username     string `json:"username"`
			LanguageCode string `json:"language_code"`
		} `json:"from"`
		Chat struct {
			Id        int    `json:"id"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
			Type      string `json:"type"`
		} `json:"chat"`
		Entity []Entities `json:"entities"`
		Date   int        `json:"date"`
		Text   string     `json:"text"`
	} `json:"message"`
}

type Entities struct {
	Type     string `json:"type"`
	Offset   int    `json:"offset"`
	Length   int    `json:"length"`
	Language string `json:"language"`
}

type DetectLanguageResponse struct {
	LanguageCode string `json:"languageCode"`
}

type TranslateTextResponse struct {
	Translations []struct {
		Text                 string `json:"text"`
		DetectedLanguageCode string `json:"detectedLanguageCode"`
	} `json:"translations"`
}

func (e Entities) IsCode() bool {
	return e.Type == "code"
}
func (e Entities) IsPre() bool {
	return e.Type == "pre"
}
