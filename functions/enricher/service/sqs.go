package service

import "aclue.de/codetalks-lambda-enricher/model"

type sqsService struct {
}

type SqsService interface {
	Publish(model.EnricherModel) error
}

func NewSqsService() SqsService {
	return &sqsService{}
}

func (s sqsService) Publish(model model.EnricherModel) error {
	return nil
}
