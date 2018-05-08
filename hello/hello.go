package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

func Hello() {
	sess := session.Must(session.NewSession())
	client := lambda.New(sess)
	functionName := "go-lambda-dev-world"
	requestResponse := "RequestResponse"
	_, err := client.Invoke(&lambda.InvokeInput{
		FunctionName:   &functionName,
		InvocationType: &requestResponse,
	})
	if err != nil {
		log.Println("Erro ao chamar lambda world", err)
	}
}
