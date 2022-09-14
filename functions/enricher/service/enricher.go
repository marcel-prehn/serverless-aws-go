package service

import (
	"aclue.de/codetalks-lambda-enricher/model"
	"github.com/aws/aws-lambda-go/events"
)

type enricherService struct {
}

type EnricherService interface {
	Enrich(events.SQSMessage) (model.EnricherModel, error)
}

func NewEnricherService() EnricherService {
	return &enricherService{}
}

func (s enricherService) Enrich(record events.SQSMessage) (model.EnricherModel, error) {
	return model.EnricherModel{}, nil
}
