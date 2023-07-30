package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler() (string, error) {
  fmt.Println("Hello, World!")
	return "Hello, World!", nil
}

func main() {
	lambda.Start(handler)
}
