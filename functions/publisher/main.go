package main

import (
	"net/http"

	"aclue.de/codetalks-lambda-publisher/service"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	mappingService := service.NewMappingService()
	dynamodbService := service.NewDynamodbService()

	mapping, err := mappingService.Unmarshal(request.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}
	}
	entity, err := dynamodbService.GetById(mapping.Id)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}
	}
	response, err := mappingService.Marshal(entity)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       response,
	}
}
