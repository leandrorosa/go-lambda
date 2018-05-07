package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
)

type Response struct {
	Message string `json:"message"`
}


func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       "Go serverless local! Leandro bacana",
		StatusCode: 200,
	}, nil
}
func main() {
	lambda.Start(Handler)
}
