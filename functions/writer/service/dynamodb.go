package service

type dynamodbService struct {
}

type DynamodbService interface {
	Save(string) error
}

func NewDynamodbService() DynamodbService {
	return &dynamodbService{}
}

func (s dynamodbService) Save(body string) error {
	return nil
}
