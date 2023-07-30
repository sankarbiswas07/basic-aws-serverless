package main

import (
	"fmt"
	"context"
	// "encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/sankarbiswas07/basic-aws-serverless/cmd/get"
	"github.com/sankarbiswas07/basic-aws-serverless/cmd/post"
)

func router(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
 	// Define your routes and handlers here
	 switch request.Path {
	case "/get":
		return functionA.get(ctx, request)

	case "/post":
		return functionB.post(ctx, request)

	default:
		return events.APIGatewayProxyResponse{
			StatusCode: 404,
			Body:       "Not Found",
		}, nil
	}
}

func main() {
	lambda.Start(router)
}
