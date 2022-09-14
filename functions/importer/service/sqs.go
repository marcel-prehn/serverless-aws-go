package service

import "aclue.de/codetalks-lambda-importer/model"

type sqsService struct {
}

type SqsService interface {
	Publish(model.Object) error
}

func NewSqsService() SqsService {
	return &sqsService{}
}

func (s sqsService) Publish(object model.Object) error {
	return nil
}
