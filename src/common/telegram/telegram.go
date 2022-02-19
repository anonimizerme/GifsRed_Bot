package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"main/src/common/redgifs"
	"net/http"
)

type ResultGif struct {
	Type     string `json:"type"`
	Id       string `json:"id"`
	GifUrl   string `json:"gif_url"`
	ThumbUrl string `json:"thumb_url"`
	//Title    string `json:"title"`
	//Caption  string `json:"caption"`
}

type ResponseBody struct {
	InlineQueryId string      `json:"inline_query_id"`
	CacheTime     int         `json:"cache_time"`
	Results       []ResultGif `json:"results"`
}

type TelegramBot struct {
	Token string
}

func GetTelegramBot(token string) TelegramBot {
	return TelegramBot{Token: token}
}

func (b *TelegramBot) AnswerInlineQuery(queryId string, response *redgifs.Response) (bool, error) {
	var results []ResultGif

	for _, gif := range response.Gifs {
		results = append(results, ResultGif{
			Type:     "gif",
			Id:       gif.Id,
			GifUrl:   gif.Urls.HDUrl,
			ThumbUrl: gif.Urls.Thumbnail,
		})
	}

	responseBytes, err := json.Marshal(ResponseBody{
		InlineQueryId: queryId,
		Results:       results,
		CacheTime:     1,
	})

	if err != nil {
		return false, err
	}

	fmt.Printf("Send JSON to telegram: %s", responseBytes)

	_, err = http.Post(
		"https://api.telegram.org/"+b.Token+"/answerInlineQuery",
		"application/json",
		bytes.NewReader(responseBytes),
	)

	if err != nil {
		fmt.Printf("Error after request: %v", err)
		return false, err
	}

	return true, nil
}
