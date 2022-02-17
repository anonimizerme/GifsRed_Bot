package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request events.APIGatewayProxyRequest

type Response events.APIGatewayProxyResponse

type RequestBody struct {
	UpdateId    int `json:"update_id"`
	InlineQuery struct {
		Id    string `json:"id"`
		Query string `json:"query"`
	} `json:"inline_query"`
}

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

func Handler(_ context.Context, req Request) (Response, error) {
	var body RequestBody

	err := json.Unmarshal([]byte(req.Body), &body)

	if err != nil {
		fmt.Println(err)
		return Response{
			StatusCode: 500,
		}, err
	}

	responseBytes, err := json.Marshal(ResponseBody{
		InlineQueryId: "28897460322723016",
		Results: []ResultGif{{
			Type:     "gif",
			Id:       fmt.Sprintf("id: %s", body.InlineQuery.Query),
			GifUrl:   "https://thumbs2.redgifs.com/BiodegradableJealousAnaconda-mobile.mp4#t=0",
			ThumbUrl: "https://thumbs2.redgifs.com/BiodegradableJealousAnaconda-mobile.mp4#t=0",
			Title:    fmt.Sprintf("title: %s", body.InlineQuery.Query),
			Caption:  fmt.Sprintf("caption: %s", body.InlineQuery.Query),
		}},
	})

	return Response{
		StatusCode: 200,
		Body:       string(responseBytes[:]),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
