package service

import "aclue.de/codetalks-lambda-publisher/model"

type dynamodbService struct {
}

type DynamodbService interface {
	GetById(string) (model.Entity, error)
}

func NewDynamodbService() DynamodbService {
	return &dynamodbService{}
}

func (s dynamodbService) GetById(id string) (model.Entity, error) {
	return model.Entity{}, nil
}
