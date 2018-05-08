package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"log"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("hello")
	Hello()
	return events.APIGatewayProxyResponse{
		Body: "Go serverless local! Hello",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
