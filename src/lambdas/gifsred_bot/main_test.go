package main

import (
	"context"
	"fmt"
	"testing"
)

func TestHandler(t *testing.T) {
	response, err := Handler(context.Background(), Request{
		Body: `
			{
				"update_id": 848715309,
				"inline_query": {
					"id": "28897462726138376",
					"from": {
						"id": 6728214,
						"is_bot": false,
						"first_name": "Alexey",
						"last_name": "Korzhov",
						"username": "anonimizer_me",
						"language_code": "en"
					},
					"chat_type": "supergroup",
					"query": "test query",
					"offset": ""
				}
			}
		`,
	})

	if err != nil {
		t.Errorf("Error %v", err)
	}

	fmt.Println(response)
}
