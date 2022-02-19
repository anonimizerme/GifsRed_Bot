package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ResultGif struct {
	Type     string `json:"type"`
	Id       string `json:"id"`
	GifUrl   string `json:"gif_url"`
	ThumbUrl string `json:"thumb_url"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
}

type ResponseBody struct {
	InlineQueryId string      `json:"inline_query_id"`
	Results       []ResultGif `json:"results"`
}

type TelegramBot struct {
	Token string
}

func GetTelegramBot(token string) TelegramBot {
	return TelegramBot{Token: token}
}

func (b *TelegramBot) AnswerInlineQuery(queryId string, query string) (bool, error) {
	responseBytes, err := json.Marshal(ResponseBody{
		InlineQueryId: queryId,
		Results: []ResultGif{{
			Type:     "gif",
			Id:       fmt.Sprintf("id: %s", query),
			GifUrl:   "https://thumbs2.redgifs.com/BiodegradableJealousAnaconda-mobile.mp4#t=0",
			ThumbUrl: "https://thumbs2.redgifs.com/BiodegradableJealousAnaconda-mobile.mp4#t=0",
			Title:    fmt.Sprintf("title: %s", query),
			Caption:  fmt.Sprintf("caption: %s", query),
		}},
	})

	if err != nil {
		return false, err
	}

	_, err = http.Post(
		"https://api.telegram.org/"+b.Token+"/answerInlineQuery",
		"application/json",
		bytes.NewReader(responseBytes),
	)

	if err != nil {
		return false, err
	}

	return true, nil
}
