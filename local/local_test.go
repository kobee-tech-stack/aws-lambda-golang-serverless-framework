package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
	"github.com/kobee-tech-stack/aws-lambda-golang-serverless-framework/services"
	"testing"
)

func TestLambdaLocally(t *testing.T) {
	// fake a request and call a handler
	request := events.APIGatewayV2HTTPRequest{
		Body: "{\"name\":\"Iphone X\", \"price\": 900}",
	}
	tableName := "global-product-development"
	dynamodb := services.NewDynamoDBStore(context.TODO(), tableName)
	domain := services.NewProductsDomain(dynamodb)
	handler := services.NewAPIGatewayV2Handler(domain)

	handler.PutHandler(context.TODO(), request)
}

func TestLambdaLocally2(t *testing.T) {
	// uuid in uuid.UUID form
	uid := uuid.New()
	fmt.Println(uid)
	// uuid v4 in a string form
	uidStr := uuid.NewString()
	fmt.Println(uidStr)

}
