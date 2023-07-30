package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	StatusCode int    `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string `json:"body"`
}

func handler() (Response, error) {
  fmt.Println("Hello, World!")
	response := Response{
		StatusCode: 200,
		Headers:    map[string]string{
			"Content-Type": "application/json",
		},
		Body: "Hello, World!",
	}
	return response, nil
}

func main() {
	lambda.Start(handler)
}
