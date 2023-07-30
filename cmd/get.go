package main

import (
	"fmt"
	"context"
	// "encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

type Request struct {
	Message string `json:"message"`
}

type Response struct {
	StatusCode int    `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string `json:"body"`
}

func get(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  fmt.Println("Hello, World!")

	// Accessing the payload/body
	body := request.Body

	// Accessing query string parameters
	paramValue := request.QueryStringParameters["paramName"]

	// Accessing headers
	headerValue := request.Headers["HeaderName"]

	// Your logic here...
	fmt.Println("Payload:", body)
	fmt.Println("Query Parameter:", paramValue)
	fmt.Println("Header:", headerValue)
		
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{
			"Content-Type": "application/json",
		},
		Body: "GET: Hello from Lambda! " + headerValue + " .",
	}
	return response, nil
}