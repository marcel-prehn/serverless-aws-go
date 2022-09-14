package main

import (
	"aclue.de/codetalks-lambda-importer/service"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(event events.S3Event) error {
	s3Service := service.NewS3Service()
	csvService := service.NewCsvService()
	sqsService := service.NewSqsService()

	for _, record := range event.Records {
		file, err := s3Service.Read(record)
		if err != nil {
			return err
		}

		objects, err := csvService.ParseFile(file)
		if err != nil {
			return err
		}

		err = sqsService.Publish(objects)
		if err != nil {
			return err
		}
	}
	return nil
}
