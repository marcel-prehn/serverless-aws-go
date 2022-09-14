package service

import "aclue.de/codetalks-lambda-publisher/model"

type mappingService struct {
}

type MappingService interface {
	Unmarshal(string) (model.Mapping, error)
	Marshal(model.Entity) (string, error)
}

func NewMappingService() MappingService {
	return &mappingService{}
}

func (s mappingService) Unmarshal(body string) (model.Mapping, error) {
	return model.Mapping{}, nil
}

func (s mappingService) Marshal(entity model.Entity) (string, error) {
	return "", nil
}
