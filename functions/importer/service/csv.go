package service

import "aclue.de/codetalks-lambda-importer/model"

type csvService struct {
}

type CsvService interface {
	ParseFile(model.S3File) (model.Object, error)
}

func NewCsvService() CsvService {
	return &csvService{}
}

func (s csvService) ParseFile(file model.S3File) (model.Object, error) {
	return model.Object{}, nil
}
