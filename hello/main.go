package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("hello")
	Hello()
	return events.APIGatewayProxyResponse{
		Body:       "Go serverless local! Leandro bacana",
		StatusCode: 200,
	}, nil
}
func main() {
	lambda.Start(Handler)
}
