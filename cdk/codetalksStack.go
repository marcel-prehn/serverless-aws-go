package main

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3notifications"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CodetalksStackProps struct {
	awscdk.StackProps
}

func NewCodetalksStack(scope constructs.Construct, id string, props *CodetalksStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// Lambdas
	importerLambda := newGoLambda(stack, "importer")
	enricherLambda := newGoLambda(stack, "enricher")
	writerLambda := newGoLambda(stack, "writer")
	publisherLambda := newGoLambda(stack, "publisher")

	// S3 Bucket
	bucket := awss3.NewBucket(scope, jsii.String("bucket"), &awss3.BucketProps{})
	bucket.GrantRead(importerLambda, nil)
	bucket.AddObjectCreatedNotification(awss3notifications.NewLambdaDestination(importerLambda))

	// DynamoDB
	dataTable := awsdynamodb.NewTable(stack, jsii.String("data-table"), &awsdynamodb.TableProps{
		TableName: jsii.String("data-table"),
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("id"),
			Type: awsdynamodb.AttributeType_STRING,
		},
	})
	dataTable.GrantReadWriteData(writerLambda)
	dataTable.GrantReadData(publisherLambda)
	writerLambda.AddEnvironment(jsii.String("TABLE_NAME"), dataTable.TableName(), &awslambda.EnvironmentOptions{})
	publisherLambda.AddEnvironment(jsii.String("TABLE_NAME"), dataTable.TableName(), &awslambda.EnvironmentOptions{})

	// SQS Queues
	enricherQueue := awssqs.NewQueue(scope, jsii.String("enricher-queue"), &awssqs.QueueProps{})
	enricherQueue.GrantSendMessages(importerLambda)
	enricherQueue.GrantConsumeMessages(enricherLambda)
	writerQueue := awssqs.NewQueue(scope, jsii.String("writer-queue"), &awssqs.QueueProps{})
	writerQueue.GrantSendMessages(enricherLambda)
	writerQueue.GrantConsumeMessages(writerLambda)

	// API Gateway
	apiIamConditions := make(map[string]interface{})
	apiIamConditions["StringEquals"] = map[string]string{"aws:SourceAccount": *stack.Account()}
	api := awsapigateway.NewRestApi(stack, jsii.String("rest-api"), &awsapigateway.RestApiProps{
		RestApiName:   jsii.String("rest-api"),
		EndpointTypes: &[]awsapigateway.EndpointType{awsapigateway.EndpointType_REGIONAL},
	})
	aobResource := api.Root().AddResource(jsii.String("data"), &awsapigateway.ResourceOptions{})
	aobResource.AddMethod(jsii.String("GET"),
		awsapigateway.NewLambdaIntegration(publisherLambda, &awsapigateway.LambdaIntegrationOptions{
			Proxy: jsii.Bool(true),
		}), aobResource.DefaultMethodOptions())

	return stack
}

func main() {
	app := awscdk.NewApp(nil)
	NewCodetalksStack(app, "AobServiceStack", &CodetalksStackProps{
		awscdk.StackProps{
			Env: nil,
		},
	})
	app.Synth(nil)
}

func newGoLambda(stack constructs.Construct, functionName string) awscdklambdagoalpha.GoFunction {
	lambda := awscdklambdagoalpha.NewGoFunction(stack, jsii.String(fmt.Sprintf("aob-lambda-%s", functionName)), &awscdklambdagoalpha.GoFunctionProps{
		FunctionName: jsii.String(functionName),
		Entry:        jsii.String(fmt.Sprintf("../functions/%s/main.go", functionName)),
	})
	return lambda
}
