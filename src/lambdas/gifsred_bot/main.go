package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"main/src/common/redgifs"
	"main/src/common/telegram"
	"os"
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

func Handler(_ context.Context, req Request) (Response, error) {
	var body RequestBody

	err := json.Unmarshal([]byte(req.Body), &body)

	if err != nil {
		fmt.Println(err)
		return Response{
			StatusCode: 500,
		}, err
	}

	client := telegram.GetTelegramBot(os.Getenv("TELEGRAM_BOT_TOKEN"))

	redgifsResult, err := redgifs.Search(body.InlineQuery.Query)
	if err != nil {
		fmt.Println(err)
		return Response{
			StatusCode: 500,
		}, err
	}

	_, err = client.AnswerInlineQuery(body.InlineQuery.Id, redgifsResult)

	if err != nil {
		fmt.Println(err)
		return Response{
			StatusCode: 500,
		}, err
	}

	return Response{
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
