package service

import (
	"aclue.de/codetalks-lambda-importer/model"
	"github.com/aws/aws-lambda-go/events"
)

type s3Service struct {
}

type S3Service interface {
	Read(events.S3EventRecord) (model.S3File, error)
}

func NewS3Service() S3Service {
	return &s3Service{}
}

func (s s3Service) Read(record events.S3EventRecord) (model.S3File, error) {
	return model.S3File{}, nil
}
